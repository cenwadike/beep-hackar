// import { BeepContractClient } from "../smart-contract-client/mono-chain-beep";
import IntentListener, { Config, IntentCreatedEvent } from "./event-listener";
import { BeepContractClient } from "../smart-contract-client/mono-chain-beep";

const config: Config = {
    rpcHttpEndpoint: "https://rpc-palvus.pion-1.ntrn.tech",
    contractAddress: "neutron13r9m3cn8zu6rnmkepajnm04zrry4g24exy9tunslseet0s9wrkkstcmkhr"
};

const listener = new IntentListener(config);
const client = new BeepContractClient(
    "neutron13r9m3cn8zu6rnmkepajnm04zrry4g24exy9tunslseet0s9wrkkstcmkhr",
    "https://rpc-palvus.pion-1.ntrn.tech"
)
console.log("MNEM: ", process.env.MNEMONIC);
client.connect(process.env.MNEMONIC!);

// Callback function to handle new intent events
const handleNewIntent = async(event: IntentCreatedEvent) => {
    console.log("New Intent Created:");
    console.log(`Intent ID: ${event.intentId}`);
    console.log(`Status: ${event.status}`);
    console.log(`Creator: ${event.sender}`);
    console.log(`Block Height: ${event.blockHeight}`);

    console.log("Validating Intent...");
    const intent_id = event.intentId;
    const intent_res = await client.getIntent(intent_id);
    const intent = intent_res.intent;

    if (intent.status != "Active") {
        console.log("Skipping Intent. Reason: Intent not active");
    }

    // increasing allowances
    console.log("Increasing expected token allowance");
    for (let i = 0; i < intent.intent_type.Swap.output_tokens.length; i++ ) {
        let token = intent.intent_type.Swap.output_tokens[i];
        await client.increaseAllowance(token.token, config.contractAddress, token.amount);
    }

    // fill intent
    console.log("Filling intent");
    const res = await client.fillIntent(intent_id, intent.intent_type.Swap.output_tokens);

    console.log("Filled intent successfully. Tx Hash: ", res.transactionHash);
    console.log("------------------------");
};

// Start the listener
export async function fillIntent() {
    try {
        listener.startListening(handleNewIntent);
    } catch (error) {
        console.error("Startup failed:", error instanceof Error ? error.message : String(error));
    }
}

// (async () => {
//     fillIntent();
// })();

// Handle process termination
process.on("SIGINT", async () => {
    await listener.cleanup();
    console.log("Shutdown complete.");
    process.exit(0);
});
