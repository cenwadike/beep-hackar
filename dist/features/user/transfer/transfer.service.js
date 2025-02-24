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
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const paystack_service_1 = require("../../../shared/services/paystack/paystack.service");
const dotenv_1 = __importDefault(require("dotenv"));
const index_1 = require("../../../shared/services/blockchain/blockchain-client-two/index");
const tx_1 = require("../../../shared/services/blockchain/blockchain-client-two/tx");
dotenv_1.default.config();
class TransferService {
    constructor({ userModel, transactionModel, encryptionRepo }) {
        this.paystackService = new paystack_service_1.PaystackService();
        this.tokenFactoryClient = new index_1.TokenFactoryClient(process.env.RPC, process.env.TOKEN_CONTRACT_ADDRESS);
        this.beepTxClient = new tx_1.BeepTxClient();
        this.start = () => __awaiter(this, void 0, void 0, function* () {
            return `CON Enter PIN `;
        });
        this.verifyUser = (phoneNumber, pin) => __awaiter(this, void 0, void 0, function* () {
            const checkUser = yield this._userModel.checkIfExist({ phoneNumber });
            if (!checkUser.data)
                return `END Unable to get your account`;
            const veryPin = this._encryptionRepo.comparePassword(pin, checkUser.data.pin);
            if (!veryPin)
                return `END Incorrect PIN`;
            return `CON Enter Amount`;
        });
        this.enterAddress = () => __awaiter(this, void 0, void 0, function* () {
            return `CON Enter wallet Address `;
        });
        this.transfer = (phoneNumber, amount, address) => __awaiter(this, void 0, void 0, function* () {
            const checkUser = yield this._userModel.checkIfExist({ phoneNumber });
            if (!checkUser.data)
                return `END Unable to get your account`;
            const mnemonic = this._encryptionRepo.decryptToken(checkUser.data.privateKey, process.env.ENCRYTION_KEY);
            const connectWallet = yield this.tokenFactoryClient.connectWallet(mnemonic);
            const balanceMsg = yield this.beepTxClient.balance(checkUser.data.publicKey);
            const transferMsg = yield this.beepTxClient.transfer(address, (parseFloat(amount) * 1000000).toString());
            const getBeepTokenBalance = yield this.tokenFactoryClient.query(connectWallet.client, balanceMsg);
            if (!getBeepTokenBalance.status)
                return `END Unable to carry out Transaction`;
            if ((getBeepTokenBalance.result.balance / 1000000) < parseFloat(amount))
                return `END Insufficient balance`;
            const transferToken = yield this.tokenFactoryClient.tx(connectWallet.client, connectWallet.sender, transferMsg);
            if (!transferToken.status)
                return `END Unable to create transaction`;
            return `END Transaction in progress`;
        });
        this._userModel = userModel;
        this._transactionModel = transactionModel;
        this._encryptionRepo = encryptionRepo;
    }
    generateUniqueCode() {
        const now = new Date();
        const day = String(now.getDate()).padStart(2, '0'); // 2-digit day
        const minutes = String(now.getMinutes()).padStart(2, '0'); // 2-digit minutes
        const ms = String(now.getMilliseconds()).slice(-1); // Last digit of milliseconds
        return `${day}${minutes}${ms}`; // Example: "27154"
    }
}
exports.default = TransferService;
