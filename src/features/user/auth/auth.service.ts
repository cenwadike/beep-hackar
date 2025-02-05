import IUserAccountModel from "../../../shared/services/database/user/Account/type";
import EncryptionInterface from "../../../shared/services/encryption/type";
import BlockchainAccount from "../../../shared/services/blockchain/account";

class AuthService {
    private _userModel: IUserAccountModel
    private _encryptionRepo: EncryptionInterface

    private blockchain = new BlockchainAccount()

    constructor({userModel, encryptionRepo}: {
        userModel: IUserAccountModel;
        encryptionRepo: EncryptionInterface
    }){
        this._userModel = userModel
        this._encryptionRepo = encryptionRepo
    }

    public start = async (phoneNumber: string) => {
        const checkUser = await this._userModel.checkIfExist({phoneNumber})
        if (!checkUser.status) {
            return `CON Carrier info
            11. Create account
            0. Exist
            `;
        }

        if (!checkUser.data?.pin) {
            return `CON Carrier info
            12. Create pin
            0. Back`;
        }

        return `CON Carrier info
        1. Deposit
        2. Transfer
        3. Withdraw
        4. verify deposit transaction
        0. Back`;
    }

    public createAccount = async (phoneNumber: string) => {
        const checkUser = await this._userModel.checkIfExist({phoneNumber})
        if (checkUser.data) return `END You already have account`;

        const blockChainAccount = await this.blockchain.createAccount()
        const publicKey = blockChainAccount.publicKey
        const privateKey = this._encryptionRepo.encryptToken(blockChainAccount.mnemonic,process.env.ENCRYTION_KEY as string )
        
        const createAccount = await this._userModel.createAccountToDB({phoneNumber, publicKey, privateKey})
        if (!createAccount.data)  return `END Unable to create account`;

        return `CON Carrier info
        1. Create pin
        0. Back`;
    }

    public enterPin = async () => {
        return `CON Enter PIN `;
    }

    public createPin = async (phoneNumber: string, pin: string) => {
        const checkUser = await this._userModel.checkIfExist({phoneNumber})
        if (!checkUser.data) return `END Unable to get your account`;

        if (checkUser.data?.pin) return `END PIN already created`;

        if (pin.length !== 4) return `END Invalid PIN format. Please enter a 4-digit PIN.`;

        const hashPin = this._encryptionRepo.encryptPassword(pin)

        const createPin = await this._userModel.updateAccount(phoneNumber, {pin: hashPin})
        if (!createPin.data) return `END Unable to save pin`;

        return `END PIN created successfully.`;
    }


}

export default AuthService;