import { Request, Response } from "express";
import UserAccountModel from "./shared/services/database/user/Account/index";
import TransactionModel from "./shared/services/database/user/transaction/index";
import EncryptionRepo from "./shared/services/encryption/index";
import AuthService from "./features/user/auth/auth.service";
import DepositService from "./features/user/deposit/deposite.service";


const encryptionRepo = new EncryptionRepo()

const userModel = new UserAccountModel()
const transactionModel = new TransactionModel()

const authService = new AuthService({userModel, encryptionRepo})
const depositService = new DepositService({userModel, transactionModel, encryptionRepo})

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
    }

    res.set('Content-Type: text/plain');
    res.send(response);
}