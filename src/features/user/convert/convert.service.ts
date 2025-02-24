import IUserAccountModel from "../../../shared/services/database/user/Account/type";
import ITransactionModel from "../../../shared/services/database/user/transaction/type";
import EncryptionInterface from "../../../shared/services/encryption/type";
import { PaystackService } from "../../../shared/services/paystack/paystack.service";
import { sendSms } from "../../../shared/services/sms/termii";
import { TransactionStatus, TransactionTypeEnum } from "../../../shared/types/interfaces/responses/user/transaction.response";
import { modifiedPhoneNumber } from "../../../shared/constant/mobileNumberFormatter";
import { TokenFactoryClient } from "../../../shared/services/blockchain/blockchain-client-two/index";
import { BeepTxClient } from "../../../shared/services/blockchain/blockchain-client-two/tx";
import dotenv from "dotenv";

dotenv.config();

class ConvertService {
    private _userModel: IUserAccountModel
    private _transactionModel: ITransactionModel
    private _encryptionRepo: EncryptionInterface
    private paystackService = new PaystackService()
    private tokenFactoryClient = new TokenFactoryClient(process.env.RPC as string, process.env.TOKEN_CONTRACT_ADDRESS as string)
    private beepTxClient = new BeepTxClient()

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

    public convertBNGNToBToken = async (phoneNumber: string, amount: string) => {
        const checkUser = await this._userModel.checkIfExist({phoneNumber})
        if (!checkUser.data) return `END Unable to get your account`;

        if (checkUser.data.balance < parseFloat(amount)) return `END Insufficient bNGN balance`;

        const  {id} = checkUser.data

        const newBalance = checkUser.data.balance - parseFloat(amount)

        const reference = this.generateUniqueCode()
        const newTransaction = await this._transactionModel.createTransactionToDB({userId: id, amount: parseFloat(amount), reference: reference, type: TransactionTypeEnum.DEBIT, status: TransactionStatus.COMPLETED})
        if (!newTransaction.data)  return `END Unable to create transaction`;

        const updateBalance = await this._userModel.updateAccount(phoneNumber, {balance: newBalance})
        if (!updateBalance.data) return `END Unable to carry out Transaction`;

        // const mnemonic =  this._encryptionRepo.decryptToken(checkUser.data.privateKey, process.env.ENCRYTION_KEY as string )
        const adminMnemonic =  process.env.ADMIN_MNEMONIC as string 

        const adminConnectWallet = await this.tokenFactoryClient.connectWallet(adminMnemonic)

        const mintMsg = await this.beepTxClient.mint(checkUser.data.publicKey, amount)

        const mintToken = await this.tokenFactoryClient.tx(adminConnectWallet.client, adminConnectWallet.sender,  mintMsg)
        if (!mintToken.status) return `END Unable to carry out Transaction`;

        return `END Transaction in progress`;
      
    }

    public convertBTokenToBNGN = async (phoneNumber: string, amount: string) => {
        const checkUser = await this._userModel.checkIfExist({phoneNumber})
        if (!checkUser.data) return `END Unable to get your account`;

        // to do
        // check blockchain balnce
        const  {id} = checkUser.data

        const mnemonic =  this._encryptionRepo.decryptToken(checkUser.data.privateKey, process.env.ENCRYTION_KEY as string )

        const connectWallet = await this.tokenFactoryClient.connectWallet(mnemonic)

        const balanceMsg = await this.beepTxClient.balance(checkUser.data.publicKey)
        const burnMsg = await this.beepTxClient.burn(amount)

        const getBeepTokenBalance = await this.tokenFactoryClient.query(connectWallet.client, balanceMsg)
        if (!getBeepTokenBalance.status) return `END Unable to carry out Transaction`;

        if (getBeepTokenBalance.result.balance < amount) return `END Insufficient balance`;

        const burnBeepToken = await this.tokenFactoryClient.tx(connectWallet.client, connectWallet.sender, burnMsg)
        if (!burnBeepToken.status) return `END Unable to create transaction`;

        console.log('getBalance', getBeepTokenBalance.result.balance)

        // to do
        // do login to convert bNaira to bToken

        const newBalance = checkUser.data.balance + parseFloat(amount)

        const reference = this.generateUniqueCode()
        const newTransaction = await this._transactionModel.createTransactionToDB({userId: id, amount: parseFloat(amount), reference: reference, type: TransactionTypeEnum.CREDIT, status: TransactionStatus.COMPLETED})
        if (!newTransaction.data)  return `END Unable to create transaction`;

        const updateBalance = await this._userModel.updateAccount(phoneNumber, {balance: newBalance})
        if (!updateBalance.data) return `END Unable to carry ou Transaction`;

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

export default ConvertService