import { SigningCosmWasmClient } from "@cosmjs/cosmwasm-stargate";
import { DirectSecp256k1HdWallet, AccountData } from "@cosmjs/proto-signing";
import { StargateClient } from "@cosmjs/stargate"
import { GasPrice } from "@cosmjs/stargate";

export class TokenFactoryClient {
    private rpcEndpoint: string;
    private contractAddress: string;

    constructor(rpcEndpoint: string, contractAddress: string) {
        this.rpcEndpoint = rpcEndpoint;
        this.contractAddress = contractAddress;
    }

    async createAccount() {
        const prefix = "neutron";
        const wallet = await DirectSecp256k1HdWallet.generate(12, {prefix});
        const address = await wallet.getAccounts()

        return {
            publicKey: address[0].address,
            mnemonic: wallet.mnemonic
        }
    }

    async connectWallet(mnemonic: string ) {
        const wallet = await DirectSecp256k1HdWallet.fromMnemonic(mnemonic, {
            prefix: "neutron",
        });

        const client = await SigningCosmWasmClient.connectWithSigner(
            this.rpcEndpoint,
            wallet,
            { gasPrice: GasPrice.fromString("0.025untrn") }
        );
    
        const [account] = await wallet.getAccounts();
        console.log("Sender Address:", account.address);
    
        return { client, sender: account };
    }

    async getNativeTokenBal(address: string) {
        try {
            // Connect to the blockchain
            const client = await StargateClient.connect(this.rpcEndpoint);

            // const RPC_ENDPOINT = "https://rpc-palvus.pion-1.ntrn.tech";
            // const client = await StargateClient.connect(RPC_ENDPOINT)
    
            // Query the balance of the address
            const balance = await client.getBalance(address, "untrn");
    
            console.log(`Balance: ${balance.amount} ${balance.denom}`);
            return {
                balance: balance.amount,
                denom: balance.denom
            }
        } catch (error) {
            console.error("Error fetching balance:", error);
        }
    }

    async query(client: SigningCosmWasmClient, queryMsg: any) {
        try {
            const result = await client.queryContractSmart(this.contractAddress, queryMsg);
            return {status: true, message: "transaction in process", result}
        } catch (error) {
            console.log('error', error)
            return {status: false, message: "Unable to perform transaction"}
        }
    }

    async tx(client: SigningCosmWasmClient, sender: AccountData, transferMsg: any) {
        try {
            const result = await client.execute(sender.address, this.contractAddress, transferMsg, "auto");
            return {status: true, message: "transaction in process", result}
        } catch (error) {
            console.log('error', error)
            return {status: false, message: "Unable to perform transaction"}
        }
    }
}

