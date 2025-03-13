import { SigningCosmWasmClient } from "@cosmjs/cosmwasm-stargate";
import { DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { GasPrice } from "@cosmjs/stargate";
import { readFileSync } from "fs";

import dotenv from "dotenv"

dotenv.config()

const rpcEndpoint = "https://rpc-palvus.pion-1.ntrn.tech";
const mnemonic = "empower spin strong sheriff grace dash sport film staff balcony corn august" //process.env.MNEMONIC;
const wasmFilePath = "../artifacts/mono_chain_beep.wasm";

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

  console.log("Addr: ", firstAccount.address);
  // const wasmCode = readFileSync(wasmFilePath);
  // const uploadReceipt = await client.upload(firstAccount.address, wasmCode, "auto");
  // console.log("Upload successful, code ID:", uploadReceipt.codeId);

  // // Token deployment details
  // // Upload successful, code ID: 11119
  // // Test ATOM contract instantiated at: neutron17huqemgnu8r4092z74vu5jtzgm3lxg4gzqupu48648t8fz4wyzxsy7rjkf
  // // Test NGN contract instantiated at: neutron1ujaf3dgpgn7e5tg6xy2hfnx6a6aupzjgxgj8fust08jttv03059s2jv2uw

  // const initMsg = {
  //   admin: firstAccount.address, 
  //   supported_tokens: [
  //     "neutron17huqemgnu8r4092z74vu5jtzgm3lxg4gzqupu48648t8fz4wyzxsy7rjkf", // tATOM
  //     "neutron1ujaf3dgpgn7e5tg6xy2hfnx6a6aupzjgxgj8fust08jttv03059s2jv2uw" // tNGN
  //   ], 
  //   default_timeout_height: 20, // 2 minutes = 6 sec block * 20
  //   supported_protocols: ["neutron"]
  // };

  // const instantiateReceipt = await client.instantiate(firstAccount.address, uploadReceipt.codeId, initMsg, "Beep: IBC intent settlement contract", "auto");
  // console.log("Contract instantiated at:", instantiateReceipt.contractAddress);
  
  // Upload successful, code ID: 11120
  // Contract instantiated at: neutron1rvn3dawaze4qyv7p9h4mcyll5axs0tefu90fdmw0s8u5wzaxjqhqc2m3gx
}

main().catch(console.error);
