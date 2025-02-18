import { DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { 
    SigningStargateClient, 
    GasPrice,
    QueryClient,
    StargateClient
} from "@cosmjs/stargate";
import { Tendermint34Client } from "@cosmjs/tendermint-rpc";
import { BeepQueryClient } from "./query";
import { BeepTxClient } from "./tx";
import { todo } from "node:test";
import { IntentEventListener } from "./event";

export class TokenFactoryClient {
    private rpcEndpoint: string;
    private queryClient: BeepQueryClient | null = null;
    private txClient: BeepTxClient | null = null;
    private tmClient: Tendermint34Client | null = null;

    constructor(rpcEndpoint: string) {
        this.rpcEndpoint = rpcEndpoint;
    }

    async connect() {
        this.tmClient = await Tendermint34Client.connect(this.rpcEndpoint);
        const baseQueryClient = new QueryClient(this.tmClient);
        this.queryClient = new BeepQueryClient(baseQueryClient);
        return {
            tmClient: this.tmClient,
            queryClient: this.queryClient 
        }
    }

    async connectWithWallet() {
        const prefix = "beep";
        const wallet = await DirectSecp256k1HdWallet.generate(12, {prefix});
        const signingClient = await SigningStargateClient.connectWithSigner(
            this.rpcEndpoint,
            wallet,
            { gasPrice: GasPrice.fromString("0.025token") }
        );
        const address = (await wallet.getAccounts()).at(0)
        if (!address) {
            return;
        }

        const balance = await signingClient.getBalance(address?.address, "token")
        this.txClient = new BeepTxClient(signingClient);

        console.log(`Address: ${address.address}, Balance: ${balance.amount}`)

        return {
            wallet: wallet,
            signingClient: signingClient
        }
    }

    async connecWallet(mnemonic: string ) {
        const wallet = await DirectSecp256k1HdWallet.fromMnemonic(mnemonic, {
            prefix: "beep", // Replace with your chain's address prefix
        });
    
        const [account] = await wallet.getAccounts();
        console.log("Sender Address:", account.address);
    
        const client = await SigningStargateClient.connectWithSigner(this.rpcEndpoint, wallet, {
            gasPrice: GasPrice.fromString("0.025stake"), // Adjust based on your chain
            // registry
        });
    
        return { client, sender: account.address };
    }

   

    async createAccount() {
        const prefix = "beep";
        const wallet = await DirectSecp256k1HdWallet.generate(12, {prefix});
        const address = await wallet.getAccounts()

        return {
            publicKey: address[0].address,
            mnemonic: wallet.mnemonic
        }
    }

    async getNativeTokenBal(address: string) {
        try {
            // Connect to the blockchain
            const client = await StargateClient.connect(this.rpcEndpoint,);
    
            // Query the balance of the address
            const balance = await client.getBalance(address, "bATOM");
    
            console.log(`Balance: ${balance.amount} ${balance.denom}`);
            return {
                balance: balance.amount,
                denom: balance.denom
            }
        } catch (error) {
            console.error("Error fetching balance:", error);
        }
    }

    get query() {
        if (!this.queryClient) throw new Error("Query client not connected");
        return this.queryClient;
    }

    get tx() {
        if (!this.txClient) throw new Error("Transaction client not connected");
        return this.txClient;
    }

    async events() {
        const listener = new IntentEventListener('ws://localhost:26657/websocket')

        // Set up handlers for all event types
        listener.setCreateIntentHandler((event: any) => {
            console.log('New intent created:', event)
        })

        listener.setAcceptIntentHandler((event: any) => {
            console.log('Intent accepted:', event)
        })

        listener.setAcceptIntentFailureHandler((event: any) => {
            console.error('Intent acceptance failed:', event)
        })

        listener.setExecutorEscrowHandler((event: any) => {
            console.log('Executor escrowed tokens:', event)
        })

        listener.setCreatorPayoutHandler((event: any) => {
            console.log('Creator received payout:', event)
        })

        listener.setExecutorPayoutHandler((event: any) => {
            console.log('Executor received payout:', event)
        })

        listener.setIntentExecutedHandler((event: any) => {
            console.log('Intent executed successfully:', event)
        })

        listener.setEscrowIBCTransferHandler((event: any) => {
            console.log('IBC transfer escrowed:', event)
        })

        listener.setEscrowNativeTokensHandler((event: any) => {
            console.log('Native tokens escrowed:', event)
        })

        listener.setIBCTransferFailureHandler((event: any) => {
            console.error('IBC transfer failed:', event)
        })

        // Connect to the blockchain and start listening
        await listener.connect()
    }
    
    async disconnect() {
        todo()
    }
}

const startEventListener = async () => {
    const listener = new IntentEventListener('ws://localhost:26657/websocket')

    // Set up handlers for all event types
    listener.setCreateIntentHandler((event: any) => {
        console.log('New intent created:', event)
    })

    listener.setAcceptIntentHandler((event: any) => {
        console.log('Intent accepted:', event)
    })

    listener.setAcceptIntentFailureHandler((event: any) => {
        console.error('Intent acceptance failed:', event)
    })

    listener.setExecutorEscrowHandler((event: any) => {
        console.log('Executor escrowed tokens:', event)
    })

    listener.setCreatorPayoutHandler((event: any) => {
        console.log('Creator received payout:', event)
    })

    listener.setExecutorPayoutHandler((event: any) => {
        console.log('Executor received payout:', event)
    })

    listener.setIntentExecutedHandler((event: any) => {
        console.log('Intent executed successfully:', event)
    })

    listener.setEscrowIBCTransferHandler((event: any) => {
        console.log('IBC transfer escrowed:', event)
    })

    listener.setEscrowNativeTokensHandler((event: any) => {
        console.log('Native tokens escrowed:', event)
    })

    listener.setIBCTransferFailureHandler((event: any) => {
        console.error('IBC transfer failed:', event)
    })

    // Connect to the blockchain and start listening
    await listener.connect()
};

// Start the listener
startEventListener().catch(console.error)

export * from './types';
export * from './query';
export * from './tx';
export * from './event';