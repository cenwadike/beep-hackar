import IUserAccountModel from "../../../shared/services/database/user/Account/type";
import EncryptionInterface from "../../../shared/services/encryption/type";
import { sendSms } from "../../../shared/services/sms/termii";
import { modifiedPhoneNumber } from "../../../shared/constant/mobileNumberFormatter";
// import BlockchainAccount from "../../../shared/services/blockchain/account";

class AuthService {
    private _userModel: IUserAccountModel
    private _encryptionRepo: EncryptionInterface

    // private blockchain = new BlockchainAccount()

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
        2. Transfer bTokn
        3. Withdraw bNGN
        4. verify deposit transaction
        5. convert bNGN to bToken
        6. convert bToken to bNGN
        7. balance
        0. Back`;
    }

    public createAccount = async (phoneNumber: string) => {
        const checkUser = await this._userModel.checkIfExist({phoneNumber})
        if (checkUser.data) return `END You already have account`;

        // const blockChainAccount = await this.blockchain.createAccount()
        const publicKey = "generalpublickeururu"
        const privateKey = this._encryptionRepo.encryptToken("general994848900044", process.env.ENCRYTION_KEY as string )
        
        const createAccount = await this._userModel.createAccountToDB({phoneNumber, publicKey, privateKey})
        if (!createAccount.data)  return `END Unable to create account`;

        return `CON Carrier info
        1. Create pin
        0. Back`;
    }

    public enterPin = async () => {
        return `CON Enter PIN`;
    }

    public verifyUser = async (phoneNumber: string, pin: string) => {
        const checkUser = await this._userModel.checkIfExist({phoneNumber})
        if (!checkUser.data) return `END Unable to get your account`;

        const veryPin = this._encryptionRepo.comparePassword(pin, checkUser.data.pin)
        if (!veryPin) return `END Incorrect PIN`;

        return `CON Enter Amount`;
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

    public getBalance = async (phoneNumber: string, pin: string) => {
        const checkUser = await this._userModel.checkIfExist({phoneNumber})
        if (!checkUser.data) return `END Unable to get your account`;

        const veryPin = this._encryptionRepo.comparePassword(pin, checkUser.data.pin)
        if (!veryPin) return `END Incorrect PIN`;

        // to do 
        //get the real bToken balance from blockchain
        const bTokenBalance = 20

        const bNGNBalance = checkUser.data.balance

        let mobileNumber = modifiedPhoneNumber(phoneNumber);

        const text = `bNGN balance: ${bNGNBalance}, bToken balance: ${bTokenBalance}`

        sendSms(mobileNumber, text)

        return `END bNGN balance: ${bNGNBalance}
        bToken balance: ${bTokenBalance}`;
    }


}

export default AuthService;