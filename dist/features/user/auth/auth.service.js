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
const termii_1 = require("../../../shared/services/sms/termii");
const mobileNumberFormatter_1 = require("../../../shared/constant/mobileNumberFormatter");
const dotenv_1 = __importDefault(require("dotenv"));
// import BlockchainAccount from "../../../shared/services/blockchain/account";
const index_1 = require("../../../shared/services/blockchain/blockchain-client-two/index");
const tx_1 = require("../../../shared/services/blockchain/blockchain-client-two/tx");
dotenv_1.default.config();
class AuthService {
    constructor({ userModel, encryptionRepo }) {
        this.tokenFactoryClient = new index_1.TokenFactoryClient(process.env.RPC, process.env.TOKEN_CONTRACT_ADDRESS);
        this.beepTxClient = new tx_1.BeepTxClient();
        this.start = (phoneNumber) => __awaiter(this, void 0, void 0, function* () {
            var _a;
            const checkUser = yield this._userModel.checkIfExist({ phoneNumber });
            if (!checkUser.status) {
                return `CON Carrier info
            11. Create account
            0. Exist
            `;
            }
            if (!((_a = checkUser.data) === null || _a === void 0 ? void 0 : _a.pin)) {
                return `CON Carrier info
            12. Create pin
            0. Back`;
            }
            return `CON Carrier info
        1. Deposit
        2. Transfer Crypto
        3. Withdraw Naira
        4. Verify Deposit 
        5. Convert Naira to Crypto
        6. Convert Crypto to Naira
        7. Get Balance
        0. Back`;
        });
        this.createAccount = (phoneNumber) => __awaiter(this, void 0, void 0, function* () {
            const checkUser = yield this._userModel.checkIfExist({ phoneNumber });
            if (checkUser.data)
                return `END You already have account`;
            const blockChainAccount = yield this.tokenFactoryClient.createAccount();
            const publicKey = blockChainAccount.publicKey;
            const privateKey = this._encryptionRepo.encryptToken(blockChainAccount.mnemonic, process.env.ENCRYTION_KEY);
            const createAccount = yield this._userModel.createAccountToDB({ phoneNumber, publicKey, privateKey });
            if (!createAccount.data)
                return `END Unable to create account`;
            return `CON Carrier info
        1. Create pin
        0. Back`;
        });
        this.enterPin = () => __awaiter(this, void 0, void 0, function* () {
            return `CON Enter PIN`;
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
        this.createPin = (phoneNumber, pin) => __awaiter(this, void 0, void 0, function* () {
            var _a;
            const checkUser = yield this._userModel.checkIfExist({ phoneNumber });
            if (!checkUser.data)
                return `END Unable to get your account`;
            if ((_a = checkUser.data) === null || _a === void 0 ? void 0 : _a.pin)
                return `END PIN already created`;
            if (pin.length !== 4)
                return `END Invalid PIN format. Please enter a 4-digit PIN.`;
            const hashPin = this._encryptionRepo.encryptPassword(pin);
            const createPin = yield this._userModel.updateAccount(phoneNumber, { pin: hashPin });
            if (!createPin.data)
                return `END Unable to save pin`;
            return `END PIN created successfully.`;
        });
        this.getBalance = (phoneNumber, pin) => __awaiter(this, void 0, void 0, function* () {
            const checkUser = yield this._userModel.checkIfExist({ phoneNumber });
            if (!checkUser.data)
                return `END Unable to get your account`;
            const veryPin = this._encryptionRepo.comparePassword(pin, checkUser.data.pin);
            if (!veryPin)
                return `END Incorrect PIN`;
            //get the real bToken balance from blockchain
            const nativeTokenBalance = yield this.tokenFactoryClient.getNativeTokenBal(checkUser.data.publicKey);
            const mnemonic = this._encryptionRepo.decryptToken(checkUser.data.privateKey, process.env.ENCRYTION_KEY);
            const connectWallet = yield this.tokenFactoryClient.connectWallet(mnemonic);
            const balanceMsg = yield this.beepTxClient.balance(checkUser.data.publicKey);
            const beepTokenBalance = yield this.tokenFactoryClient.query(connectWallet.client, balanceMsg);
            if (!beepTokenBalance.status)
                return `END Unable to get balance`;
            const bNGNBalance = checkUser.data.balance;
            let mobileNumber = (0, mobileNumberFormatter_1.modifiedPhoneNumber)(phoneNumber);
            const text = `NGN Balance: ${beepTokenBalance.result.balance}, ATOM Balance: ${beepTokenBalance.result.balance}`;
            (0, termii_1.sendSms)(mobileNumber, text);
            return `END NGN Balance: ${beepTokenBalance.result.balance}
        ATOM Balance: ${beepTokenBalance.result.balance}`;
        });
        this._userModel = userModel;
        this._encryptionRepo = encryptionRepo;
    }
}
exports.default = AuthService;
