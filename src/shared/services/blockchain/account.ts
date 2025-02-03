import { DirectSecp256k1HdWallet } from "@cosmjs/proto-signing"
 

class BlockchainAccount {
     
    public createAccount = async () => {
        const wallet: DirectSecp256k1HdWallet = await DirectSecp256k1HdWallet.generate(24)
        const mnemonic = wallet.mnemonic
        const accounts = await wallet.getAccounts()
        const publicKey = accounts[0].address

        return {
            mnemonic,
            publicKey
        }
    }
}

export default BlockchainAccount;