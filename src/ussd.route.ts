import { Request, Response } from "express";
import UserAccountModel from "./shared/services/database/user/Account/index";
import AuthService from "./features/user/auth/auth.service";
import EncryptionRepo from "./shared/services/encryption/index";

const userModel = new UserAccountModel()
const encryptionRepo = new EncryptionRepo()
const authService = new AuthService({userModel, encryptionRepo})

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
    }

    res.set('Content-Type: text/plain');
    res.send(response);
}