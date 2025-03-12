import { StargateClient } from "@cosmjs/stargate";
import { HttpClient, Tendermint37Client } from "@cosmjs/tendermint-rpc";
// import { BeepContractClient } from "../smart-contract-client";

// Configuration interface
export interface Config {
    rpcHttpEndpoint: string;
    contractAddress: string;
}

// Interface for the intent creation event
export interface IntentCreatedEvent {
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

    private parseIntentEvent(txResult: any, height: number): IntentCreatedEvent | null {
        try {
            const events = txResult.events || [];
            // console.log(`Raw events at height ${height}:`, JSON.stringify(events, null, 2));
            for (const event of events) {
                // console.log(`Event type: ${event.type}, Attributes:`, event.attributes);
                if (event.type === "wasm") {
                    const attributes: { [key: string]: string } = {};
                    for (const attr of event.attributes) {
                        // Use key and value directly, no base64 decoding
                        attributes[attr.key] = attr.value;
                    }
                    // console.log(`WASM attributes at height ${height}:`, attributes);
                    // Check for create_intent
                    if (attributes.action === "create_intent") {
                        console.log(`Found intent event at height ${height}:`, attributes.action);
                        return {
                            intentId: attributes.intent_id || "",
                            status: attributes.status || "",
                            sender: attributes.sender || "",
                            blockHeight: height
                        };
                    }
                }
            }
            console.log(`No create_intent event found at height ${height}`);
            return null;
        } catch (error) {
            console.error(`Error parsing event at height ${height}:`, error instanceof Error ? error.message : String(error));
            return null;
        }
    }

    async startListening(callback: (event: IntentCreatedEvent) => void): Promise<void> {
        try {
            await this.initialize();
    
            if (!this.stargateClient || !this.tmClient) {
                throw new Error("Client not initialized");
            }
    
            console.log("Polling started...");
            while (this.isRunning) {
                try {
                    console.log("Fetching current height...");
                    const currentHeight = await this.stargateClient.getHeight();
                    console.log(`Current height: ${currentHeight}, Last height: ${this.lastHeight}`);
                    const startHeight = Math.max(this.lastHeight + 1, currentHeight - 10);
                    // console.log(`Start height calculated: ${startHeight}`);
                    if (currentHeight >= startHeight) {
                        // console.log(`Processing blocks from ${startHeight} to ${currentHeight}`);
                        for (let height = startHeight; height <= currentHeight; height++) {
                            // console.log(`Querying transactions at height ${height}`);
                            const query = `tx.height=${height}`;
                            const txs = await this.stargateClient.searchTx(query);
                            // console.log(`Found ${txs.length} transactions at height ${height}`);
                            for (const tx of txs) {
                                const intentEvent = this.parseIntentEvent(tx, height);
                                if (intentEvent) {
                                    console.log(`Intent detected at height ${height}`);
                                    callback(intentEvent);
                                }
                            }
                        }
                        this.lastHeight = currentHeight;
                        console.log(`Updated lastHeight to ${this.lastHeight}`);
                    } else {
                        console.log("No new blocks to process");
                    }
                    console.log("Waiting 3 seconds...");
                    await new Promise(resolve => setTimeout(resolve, 6000));
                } catch (error) {
                    console.error("Polling error:", error instanceof Error ? error.message : String(error));
                    this.isRunning = false;
                    this.reconnect(callback);
                }
            }
            console.log("Polling loop exited");
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
    async cleanup(): Promise<void> {
        if (this.tmClient) {
            console.log("Disconnecting...");
            this.tmClient.disconnect();
            this.tmClient = undefined;
            this.stargateClient = undefined;
            this.isRunning = false;
        }
    }
}

export default IntentListener;

// Usage example
// const client = new BeepContractClient(
//     "neutron13r9m3cn8zu6rnmkepajnm04zrry4g24exy9tunslseet0s9wrkkstcmkhr",
//     "https://rpc-palvus.pion-1.ntrn.tech"
// )
// console.log("MNEM: ", process.env.MNEMONIC);
// client.connect(process.env.MNEMONIC!);

// const config: Config = {
//     rpcHttpEndpoint: "https://rpc-palvus.pion-1.ntrn.tech",
//     contractAddress: "neutron13r9m3cn8zu6rnmkepajnm04zrry4g24exy9tunslseet0s9wrkkstcmkhr"
// };

// const listener = new IntentListener(config);

// // Callback function to handle new intent events
// const handleNewIntent = async(event: IntentCreatedEvent) => {
//     console.log("New Intent Created:");
//     console.log(`Intent ID: ${event.intentId}`);
//     console.log(`Status: ${event.status}`);
//     console.log(`Creator: ${event.sender}`);
//     console.log(`Block Height: ${event.blockHeight}`);

    // console.log("Validating Intent...");
    // const intent_id = event.intentId;
    // const intent_res = await client.getIntent(intent_id);
    // const intent = intent_res.intent;

    // if (intent.status != "Active") {
    //     console.log("Skipping Intent. Reason: Intent not active");
    // }

    // // increasing allowances
    // console.log("Increasing expected token allowance");
    // for (let i = 0; i < intent.intent_type.Swap.output_tokens.length; i++ ) {
    //     let token = intent.intent_type.Swap.output_tokens[i];
    //     await client.increaseAllowance(token.token, config.contractAddress, token.amount);
    // }

    // // fill intent
    // console.log("Filling intent");
    // const res = await client.fillIntent(intent_id, intent.intent_type.Swap.output_tokens);

    // console.log("Filled intent successfully. Tx Hash: ", res.transactionHash);
//     console.log("------------------------");
// };


// // Start the listener
// async function main() {
//     try {
//         await listener.startListening(handleNewIntent);
//     } catch (error) {
//         console.error("Startup failed:", error instanceof Error ? error.message : String(error));
//     }
// }

// (async () => {
//     await main();
// })();

// // Handle process termination
// process.on("SIGINT", async () => {
//     await listener.cleanup();
//     console.log("Shutdown complete.");
//     process.exit(0);
// });
