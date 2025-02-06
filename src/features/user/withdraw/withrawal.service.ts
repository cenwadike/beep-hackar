import IUserAccountModel from "../../../shared/services/database/user/Account/type";
import ITransactionModel from "../../../shared/services/database/user/transaction/type";
import EncryptionInterface from "../../../shared/services/encryption/type";
import { PaystackService } from "../../../shared/services/paystack/paystack.service";
import { sendSms } from "../../../shared/services/sms/termii";
import { TransactionStatus, TransactionTypeEnum } from "../../../shared/types/interfaces/responses/user/transaction.response";
import { modifiedPhoneNumber } from "../../../shared/constant/mobileNumberFormatter";


class WithdrawalService {
    private _userModel: IUserAccountModel
    private _transactionModel: ITransactionModel
    private _encryptionRepo: EncryptionInterface
    private paystackService = new PaystackService()

    constructor({userModel, transactionModel, encryptionRepo}: {
        userModel: IUserAccountModel;
        transactionModel: ITransactionModel;
        encryptionRepo: EncryptionInterface
    }){
        this._userModel = userModel
        this._transactionModel = transactionModel
        this._encryptionRepo = encryptionRepo
    }

    public start = async () => {
        return `CON Enter PIN `;
    }

    public verifyUser = async (phoneNumber: string, pin: string) => {
        const checkUser = await this._userModel.checkIfExist({phoneNumber})
        if (!checkUser.data) return `END Unable to get your account`;

        const veryPin = this._encryptionRepo.comparePassword(pin, checkUser.data.pin)
        if (!veryPin) return `END Incorrect PIN`;

        return `CON Enter Amount`;
    }

    public enterAccountNumber = async () => {
        return `CON Enter Account Number`;
    }

    public withdraw = async (phoneNumber: string, amount: string, accountNumber: string) => {
        const checkUser = await this._userModel.checkIfExist({phoneNumber})
        if (!checkUser.data) return `END Unable to get your account`;

        if (checkUser.data.balance < parseFloat(amount)) return `END Insufficient bNGN balance`;

        const  {id} = checkUser.data

        // to do
        // check withdrawal login
 

        return `END Transaction in progress`;
    }

     generateUniqueCode() {
        const now = new Date();
        const day = String(now.getDate()).padStart(2, '0'); // 2-digit day
        const minutes = String(now.getMinutes()).padStart(2, '0'); // 2-digit minutes
        const ms = String(now.getMilliseconds()).slice(-1); // Last digit of milliseconds
    
        return `${day}${minutes}${ms}`; // Example: "27154"
    }
}

export default WithdrawalService