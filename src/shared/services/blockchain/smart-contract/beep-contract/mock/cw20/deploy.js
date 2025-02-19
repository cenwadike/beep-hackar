import { SigningCosmWasmClient } from "@cosmjs/cosmwasm-stargate";
import { DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { GasPrice } from "@cosmjs/stargate";
import { readFileSync } from "fs";

import dotenv from "dotenv"

dotenv.config()

const rpcEndpoint = "https://rpc-palvus.pion-1.ntrn.tech";
console.log(1)
const mnemonic = process.env.MNEMONIC;
// const wasmFilePath = "./artifacts/first_token_cw20contract.wasm";
const wasmFilePath = "target/wasm32-unknown-unknown/release/first_token_cw20contract.wasm";

console.log(2)

async function main() {
    console.log(3)
  const wallet = await DirectSecp256k1HdWallet.fromMnemonic(mnemonic, {
    prefix: "neutron",
  });

  console.log(4)

  const [firstAccount] = await wallet.getAccounts();

  console.log(5)

  const client = await SigningCosmWasmClient.connectWithSigner(
    rpcEndpoint,
    wallet,
    {
      gasPrice: GasPrice.fromString("0.025untrn"),
    }
  );

  console.log(6)

  const wasmCode = readFileSync(wasmFilePath);
  console.log(7)
  const uploadReceipt = await client.upload(firstAccount.address, wasmCode, "auto");
  console.log(8)
  console.log("Upload successful, code ID:", uploadReceipt.codeId);
  console.log(9)

  const initMsg = {
    name: "Test ATOM",
    symbol: "tATOM",
    decimals: 6,
    initial_balances: [
      {
        "address": "neutron107nhk9pqhp446fr0fc83z0v82rg9guy8runkuz",
        "amount": "10000000"
      },
      {
        "address": "neutron1tn5uf2q6n5ks8a40vkf2j2tkz0c9asd0udq6t4",
        "amount": "10000000"
      }
    ]
  };

  console.log(10)

  const instantiateReceipt = await client.instantiate(firstAccount.address, uploadReceipt.codeId, initMsg, "CW Token", "auto");
  console.log("Contract instantiated at:", instantiateReceipt.contractAddress);
}

main().catch(console.error);


