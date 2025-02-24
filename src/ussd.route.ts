import { Request, Response } from "express";
import UserAccountModel from "./shared/services/database/user/Account/index";
import TransactionModel from "./shared/services/database/user/transaction/index";
import WithdrawalRequestModel from "./shared/services/database/user/withdrawalRequest/index";
import EncryptionRepo from "./shared/services/encryption/index";
import AuthService from "./features/user/auth/auth.service";
import DepositService from "./features/user/deposit/deposite.service";
import ConvertService from "./features/user/convert/convert.service";
import TransferService from "./features/user/transfer/transfer.service";
import WithdrawalService from "./features/user/withdraw/withrawal.service";


const encryptionRepo = new EncryptionRepo()

const userModel = new UserAccountModel()
const transactionModel = new TransactionModel()
const withdrawalRequestModel = new WithdrawalRequestModel()

const authService = new AuthService({userModel, encryptionRepo})
const depositService = new DepositService({userModel, transactionModel, encryptionRepo})
const convertService = new ConvertService({userModel, transactionModel, encryptionRepo})
const transferService = new TransferService({userModel, transactionModel, encryptionRepo})
const withdrawalService = new WithdrawalService({userModel, transactionModel, withdrawalRequestModel, encryptionRepo})

export const ussdRoute  = async(req: Request, res: Response) => {
    const {
        sessionId,
        serviceCode,
        phoneNumber,
        text,
    } = req.body;

    console.log('sessionId', sessionId)
    console.log('serviceCode', serviceCode)
    console.log('phoneNumber', phoneNumber)
    console.log('text', text)
  
    let response = '';

    if (text == '') {
        response = await authService.start(phoneNumber);
    }else if ( text == '11') {
        response = await authService.createAccount(phoneNumber)
    }else if ( text == '11*1') {
        response = await authService.enterPin()
    }else if(text.startsWith('11*1*')){
        let pin = text.split('*')[2];
        response = await authService.createPin(phoneNumber, pin)   
    }else if(text == '12'){
        response = await authService.enterPin()
    }else if(text.startsWith('12*')){
        let pin = text.split('*')[1];
        response = await authService.createPin(phoneNumber, pin)
    }else if(text == '1'){
        response = await depositService.start()
    }else if(text.startsWith('1*')){
        let parts = text.split('*');

        if (parts.length == 2) {     
            let pin = text.split('*')[1];
            response = await depositService.verifyUser(phoneNumber, pin)
        } else if (parts.length === 3) {
            let amount = text.split('*')[2];
            response = await depositService.initializeDeposit(phoneNumber, amount)
        }
    }else if(text == '4'){
        response = await depositService.enterreference()
    }else if(text.startsWith('4*')){
        let reference = text.split('*')[1];
        response = await depositService.verifyDeposit(phoneNumber, reference)
    }else if(text == '7'){
        response = await authService.enterPin()
    }else if(text.startsWith('7*')){
        let pin = text.split('*')[1];
        response = await authService.getBalance(phoneNumber, pin)
    }else if(text == '5'){
        response = await convertService.start()
    }if(text.startsWith('5*')){
        let parts = text.split('*');

        if (parts.length == 2) {     
            let pin = text.split('*')[1];
            response = await convertService.verifyUser(phoneNumber, pin)
        } else if (parts.length === 3) {
            let amount = text.split('*')[2];
            response = await convertService.convertBNGNToBToken(phoneNumber, amount)
        }
    }else if(text == '6'){
        response = await convertService.start()
    }if(text.startsWith('6*')){
        let parts = text.split('*');

        if (parts.length == 2) {     
            let pin = text.split('*')[1];
            response = await convertService.verifyUser(phoneNumber, pin)
        } else if (parts.length === 3) {
            let amount = text.split('*')[2];
            response = await convertService.convertBTokenToBNGN(phoneNumber, amount)
        }
    }else if(text == '2'){
        response = await transferService.start()
    }if(text.startsWith('2*')){
        let parts = text.split('*');

        if (parts.length == 2) {     
            let pin = text.split('*')[1];
            response = await transferService.verifyUser(phoneNumber, pin)
        } else if (parts.length === 3) {
            response = await transferService.enterAddress()
        }else if (parts.length === 4) {
            let amount = text.split('*')[2];
            let address = text.split('*')[3];
            response = await transferService.transfer(phoneNumber, amount, address)
        }
    }else if(text == '3'){
        response = await withdrawalService.start()
    }if(text.startsWith('3*')){
        let parts = text.split('*');

        if (parts.length == 2) {     
            let pin = text.split('*')[1];
            response = await withdrawalService.verifyUser(phoneNumber, pin)
        } else if (parts.length === 3) {
            response = await withdrawalService.enterAccountNumber()
        }else if (parts.length === 4) {
            let amount = text.split('*')[2];
            let account = text.split('*')[3];
            response = await withdrawalService.enterBankName()
        }else if (parts.length === 5) {
            let amount = text.split('*')[2];
            let account = text.split('*')[3];
            let bank = text.split('*')[4];
            response = await withdrawalService.withdraw(phoneNumber, amount, account, bank)
        }
    }

    res.set('Content-Type: text/plain');
    res.send(response);
}