import { StargateClient } from "@cosmjs/stargate";
import { HttpClient, Tendermint37Client } from "@cosmjs/tendermint-rpc";
import { logs } from "@cosmjs/stargate";

// Configuration interface
interface Config {
    rpcHttpEndpoint: string;
    contractAddress: string;
}

// Interface for the intent creation event
interface IntentCreatedEvent {
    intentId: string;
    status: string;
    sender: string;
    blockHeight: number;
}

class IntentListener {
    private tmClient?: Tendermint37Client;
    private stargateClient?: StargateClient;
    private config: Config;
    private lastHeight: number = 0;
    private isRunning: boolean = false;

    constructor(config: Config) {
        this.config = config;
    }

    // Initialize HTTP connection
    private async initialize(): Promise<void> {
        try {
            const httpClient = new HttpClient(this.config.rpcHttpEndpoint);
            this.tmClient = await Tendermint37Client.create(httpClient);
            if (!this.tmClient) {
                throw new Error("Failed to create Tendermint client");
            }
            this.stargateClient = await StargateClient.create(this.tmClient);
            const status = await this.tmClient.status();
            console.log(`Connected to ${this.config.rpcHttpEndpoint} (chain: ${status.nodeInfo.network}, height: ${status.syncInfo.latestBlockHeight})`);
            this.lastHeight = status.syncInfo.latestBlockHeight;
            this.isRunning = true;
        } catch (error) {
            console.error(`Failed to connect to ${this.config.rpcHttpEndpoint}:`, error instanceof Error ? error.message : String(error));
            throw error;
        }
    }

    // Parse event attributes from transaction logs
    private parseIntentEvent(txResult: any, height: number): IntentCreatedEvent | null {
        try {
            const events = txResult.events || [];
            for (const event of events) {
                if (event.type === "wasm") {
                    const attributes: { [key: string]: string } = {};
                    for (const attr of event.attributes) {
                        attributes[attr.key] = attr.value;
                    }
                    if (attributes.action === "create_intent") {
                        return {
                            intentId: attributes.intent_id || "",
                            status: attributes.status || "",
                            sender: attributes.sender || "",
                            blockHeight: height
                        };
                    }
                }
            }
            return null;
        } catch (error) {
            console.error(`Error parsing event at height ${height}:`, error instanceof Error ? error.message : String(error));
            return null;
        }
    }

    // Poll for new transactions
    async startListening(callback: (event: IntentCreatedEvent) => void): Promise<void> {
        try {
            await this.initialize();

            if (!this.stargateClient || !this.tmClient) {
                throw new Error("Client not initialized");
            }

            console.log("Polling started...");
            while (this.isRunning) {
                try {
                    const currentHeight = await this.stargateClient.getHeight();
                    const startHeight = Math.max(this.lastHeight + 1, currentHeight - 5);
                    if (currentHeight >= startHeight) {
                        for (let height = startHeight; height <= currentHeight; height++) {
                            try {
                                const query = `tx.height=${height}`;
                                const txs = await this.stargateClient.searchTx(query);
                                for (const tx of txs) {
                                    const intentEvent = this.parseIntentEvent(tx, height);
                                    if (intentEvent) {
                                        console.log(`Intent detected at height ${height}`);
                                        callback(intentEvent);
                                    }
                                }
                            } catch (txError) {
                                console.error(`Error searching transactions at height ${height}:`, txError instanceof Error ? txError.message : String(txError));
                            }
                        }
                        this.lastHeight = currentHeight;
                    }
                    await new Promise(resolve => setTimeout(resolve, 3000)); // Poll every 3 seconds
                } catch (error) {
                    console.error("Polling error:", error instanceof Error ? error.message : String(error));
                    this.isRunning = false;
                    this.reconnect(callback);
                }
            }
        } catch (error) {
            console.error("Listener error:", error instanceof Error ? error.message : String(error));
            this.reconnect(callback);
        }
    }

    // Reconnection logic
    private async reconnect(callback: (event: IntentCreatedEvent) => void): Promise<void> {
        await this.cleanup();
        console.log("Reconnecting...");
        await new Promise(resolve => setTimeout(resolve, 2000));
        await this.startListening(callback);
    }

    // Cleanup resources
    private async cleanup(): Promise<void> {
        if (this.tmClient) {
            console.log("Disconnecting...");
            this.tmClient.disconnect();
            this.tmClient = undefined;
            this.stargateClient = undefined;
            this.isRunning = false;
        }
    }
}

// Usage example
const config: Config = {
    rpcHttpEndpoint: "https://rpc-palvus.pion-1.ntrn.tech",
    contractAddress: "neutron13r9m3cn8zu6rnmkepajnm04zrry4g24exy9tunslseet0s9wrkkstcmkhr"
};

const listener = new IntentListener(config);

// Callback function to handle new intent events
const handleNewIntent = (event: IntentCreatedEvent) => {
    console.log("New Intent Created:");
    console.log(`Intent ID: ${event.intentId}`);
    console.log(`Status: ${event.status}`);
    console.log(`Creator: ${event.sender}`);
    console.log(`Block Height: ${event.blockHeight}`);
    console.log("------------------------");
};

// Start the listener
async function main() {
    try {
        await listener.startListening(handleNewIntent);
    } catch (error) {
        console.error("Startup failed:", error instanceof Error ? error.message : String(error));
    }
}

(async () => {
    main();
})();

// Handle process termination
process.on("SIGINT", async () => {
    await listener["cleanup"]();
    console.log("Shutdown complete.");
    process.exit(0);
});
