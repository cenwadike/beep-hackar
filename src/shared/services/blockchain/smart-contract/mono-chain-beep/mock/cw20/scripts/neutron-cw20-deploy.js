import { SigningCosmWasmClient } from "@cosmjs/cosmwasm-stargate";
import { DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { GasPrice } from "@cosmjs/stargate";
import { readFileSync } from "fs";

import dotenv from "dotenv"

dotenv.config()

const rpcEndpoint = "https://rpc-palvus.pion-1.ntrn.tech";
const mnemonic = "empower spin strong sheriff grace dash sport film staff balcony corn august" // process.env.MNEMONIC;
const wasmFilePath = "../artifacts/first_token_cw20contract.wasm";

async function main() {
  const wallet = await DirectSecp256k1HdWallet.fromMnemonic(mnemonic, {
    prefix: "neutron",
  });

  const [firstAccount] = await wallet.getAccounts();

  const client = await SigningCosmWasmClient.connectWithSigner(
    rpcEndpoint,
    wallet,
    {
      gasPrice: GasPrice.fromString("0.025untrn"),
    }
  );

  const wasmCode = readFileSync(wasmFilePath);
  const uploadReceipt = await client.upload(firstAccount.address, wasmCode, "auto");
  console.log("Upload successful, code ID:", uploadReceipt.codeId);

  const initMsgAtom = {
    name: "Test ATOM",
    symbol: "tATOM",
    decimals: 6,
    initial_balances: [
      {
        "address": "neutron107nhk9pqhp446fr0fc83z0v82rg9guy8runkuz",
        "amount": "1000000000000"
      }
    ],
    mint: {
      minter: "neutron107nhk9pqhp446fr0fc83z0v82rg9guy8runkuz",
      cap: "1000000000000000000"
    }
  };

  const initMsgNgn = {
    name: "Test NGN",
    symbol: "tNGN",
    decimals: 6,
    initial_balances: [
      {
        "address": "neutron107nhk9pqhp446fr0fc83z0v82rg9guy8runkuz",
        "amount": "1000000000000"
      }
    ],
    mint: {
      minter: "neutron107nhk9pqhp446fr0fc83z0v82rg9guy8runkuz",
      cap: "1000000000000000000"
    }
  };

  const instantiateReceiptAtom = await client.instantiate(firstAccount.address, uploadReceipt.codeId, initMsgAtom, "Test ATOM Token", "auto");
  console.log("Test ATOM contract instantiated at:", instantiateReceiptAtom.contractAddress);

  const instantiateReceiptNgn = await client.instantiate(firstAccount.address, uploadReceipt.codeId, initMsgNgn, "Test NGN Token", "auto");
  console.log("Test NGN contract instantiated at:", instantiateReceiptNgn.contractAddress);

  // Upload successful, code ID: 11119
  // Test ATOM contract instantiated at: neutron17huqemgnu8r4092z74vu5jtzgm3lxg4gzqupu48648t8fz4wyzxsy7rjkf
  // Test NGN contract instantiated at: neutron1ujaf3dgpgn7e5tg6xy2hfnx6a6aupzjgxgj8fust08jttv03059s2jv2uw
}

main().catch(console.error);
