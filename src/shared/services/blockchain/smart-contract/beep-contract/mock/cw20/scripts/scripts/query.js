import { SigningCosmWasmClient, } from "@cosmjs/cosmwasm-stargate";
import { DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { StargateClient } from "@cosmjs/stargate"
import { GasPrice } from "@cosmjs/stargate";

import dotenv from "dotenv"

dotenv.config()


const CONTRACT_ADDRESS = "neutron1jwq68g39nh88acv8258ngmfccswtmx4agsjklhm94d2vszjqjpdqfnyd92";
const CHAIN_ID = "pion-1";
const RPC_ENDPOINT = "https://rpc-palvus.pion-1.ntrn.tech";
const RECIPIENT_ADDRESS = "neutron1tn5uf2q6n5ks8a40vkf2j2tkz0c9asd0udq6t4"
console.log(1)

const mnemonic = 'vapor wheel chair nature energy people chat remove impose impact figure polar'

if (!mnemonic) {
    throw new Error("MNEMONIC environment variable is not set");
  }

const gasPrice = GasPrice.fromString("0.025untrn");
console.log(2)

const wallet = await DirectSecp256k1HdWallet.fromMnemonic(mnemonic, {
    prefix: "neutron",
});

console.log(3)
// const client = await SigningCosmWasmClient.connectWithSigner(RPC_ENDPOINT, wallet, {
//     gasPrice,
// });

const client = await StargateClient.connect(RPC_ENDPOINT)

console.log(4)
const bals = await client.getAllBalances('neutron1dr9ahejux0qtgkp2576y4znta7yvu64t0k3hm3')

const bal = await client.getBalance("neutron1dr9ahejux0qtgkp2576y4znta7yvu64t0k3hm3", "untrn")
console.log(5)


console.log("wale", bal)
console.log("all", bals)


// Connect to the blockchain with a signer client to enable transaction signing
const clienttwo = await SigningCosmWasmClient.connectWithSigner(
  RPC_ENDPOINT,
  wallet,
  { gasPrice: GasPrice.fromString("0.025untrn") }
);

const [sender] = await wallet.getAccounts();
console.log("Sender address:", sender.address);

const AMOUNT = "1000";

const transferMsg = {
    transfer: {
      recipient: RECIPIENT_ADDRESS,
      amount: AMOUNT,
    },
};

const result = await clienttwo.execute(sender.address, CONTRACT_ADDRESS, transferMsg, "auto");
console.log("Transfer transaction result:", result);


const queryMsg = {
    balance: { address: "neutron1dr9ahejux0qtgkp2576y4znta7yvu64t0k3hm3" },
  };

  // Query the contract for the token balance
  const balance = await clienttwo.queryContractSmart(CONTRACT_ADDRESS, queryMsg);
  console.log("Token balance:", balance);

  const nativeAmount = "100000"; // e.g., "1000000" untrn (adjust as needed)
  const coins = [
    {
      denom: "untrn",
      amount: nativeAmount,
    },
  ];


   // Send native tokens using the sendTokens method
   const nativeTransferResult = await clienttwo.sendTokens(
    sender.address,
    "neutron1k83rc4rd4rlawy629haagqqgzkufx24vy8mgh2",
    coins,
    "auto",
    "Native token transfer" // optional memo
  );
  console.log("Native token transfer result:", nativeTransferResult);

  const minMsg = {
    mint: {
        recipient: RECIPIENT_ADDRESS,
        amount: AMOUNT,
    },
};

const mint = await clienttwo.execute(sender.address, CONTRACT_ADDRESS, minMsg, "auto");
console.log("mint result:", mint);