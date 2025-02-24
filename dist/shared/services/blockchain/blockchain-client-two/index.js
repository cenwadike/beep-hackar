"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.TokenFactoryClient = void 0;
const cosmwasm_stargate_1 = require("@cosmjs/cosmwasm-stargate");
const proto_signing_1 = require("@cosmjs/proto-signing");
const stargate_1 = require("@cosmjs/stargate");
const stargate_2 = require("@cosmjs/stargate");
class TokenFactoryClient {
    constructor(rpcEndpoint, contractAddress) {
        this.rpcEndpoint = rpcEndpoint;
        this.contractAddress = contractAddress;
    }
    createAccount() {
        return __awaiter(this, void 0, void 0, function* () {
            const prefix = "neutron";
            const wallet = yield proto_signing_1.DirectSecp256k1HdWallet.generate(12, { prefix });
            const address = yield wallet.getAccounts();
            return {
                publicKey: address[0].address,
                mnemonic: wallet.mnemonic
            };
        });
    }
    connectWallet(mnemonic) {
        return __awaiter(this, void 0, void 0, function* () {
            const wallet = yield proto_signing_1.DirectSecp256k1HdWallet.fromMnemonic(mnemonic, {
                prefix: "neutron",
            });
            const client = yield cosmwasm_stargate_1.SigningCosmWasmClient.connectWithSigner(this.rpcEndpoint, wallet, { gasPrice: stargate_2.GasPrice.fromString("0.025untrn") });
            const [account] = yield wallet.getAccounts();
            console.log("Sender Address:", account.address);
            return { client, sender: account };
        });
    }
    getNativeTokenBal(address) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                // Connect to the blockchain
                const client = yield stargate_1.StargateClient.connect(this.rpcEndpoint);
                // const RPC_ENDPOINT = "https://rpc-palvus.pion-1.ntrn.tech";
                // const client = await StargateClient.connect(RPC_ENDPOINT)
                // Query the balance of the address
                const balance = yield client.getBalance(address, "untrn");
                console.log(`Balance: ${balance.amount} ${balance.denom}`);
                return {
                    balance: balance.amount,
                    denom: balance.denom
                };
            }
            catch (error) {
                console.error("Error fetching balance:", error);
            }
        });
    }
    query(client, queryMsg) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const result = yield client.queryContractSmart(this.contractAddress, queryMsg);
                return { status: true, message: "transaction in process", result };
            }
            catch (error) {
                console.log('error', error);
                return { status: false, message: "Unable to perform transaction" };
            }
        });
    }
    tx(client, sender, transferMsg) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const result = yield client.execute(sender.address, this.contractAddress, transferMsg, "auto");
                return { status: true, message: "transaction in process", result };
            }
            catch (error) {
                console.log('error', error);
                return { status: false, message: "Unable to perform transaction" };
            }
        });
    }
}
exports.TokenFactoryClient = TokenFactoryClient;
