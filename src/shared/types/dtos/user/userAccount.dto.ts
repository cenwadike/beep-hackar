import { IUserAccount } from "../../interfaces/responses/user/userAccount.response";

class UserAccountDto implements IUserAccount {
    public id?: string;
    public phoneNumber: string;
    public pin: string;
    public publicKey: string ;
    public privateKey: string;
    public balance: number;
    public updatedAt?: Date;
    public createdAt?: Date;
    
    constructor(userAccount: IUserAccount) {
      this.id = userAccount._id;
      this.phoneNumber = userAccount.phoneNumber;
      this.pin = userAccount.pin;
      this.publicKey = userAccount.publicKey;
      this.privateKey = userAccount.privateKey;
      this.balance = userAccount.balance
      this.updatedAt = userAccount.updatedAt;
      this.createdAt = userAccount.createdAt;

    }
  
    get getModel() {
      return {
        _id: this.id,
        phoneNumber: this.phoneNumber,
        pin: this.pin,
        publicKey: this.publicKey,
        privateKey: this.privateKey,
        balance: this.balance,
        updatedAt: this.updatedAt ? new Date(this.updatedAt): undefined,
        createdAt: this.createdAt ? new Date(this.createdAt): undefined,
      } as IUserAccount
    }

    get getSecureRespons() {
        return {
          _id: this.id,
          phoneNumber: this.phoneNumber,
          publicKey: this.publicKey,
          balance: this.balance,
          updatedAt: this.updatedAt ? new Date(this.updatedAt): undefined,
          createdAt: this.createdAt ? new Date(this.createdAt): undefined,
        } as IUserAccount
      }
  
}

export default UserAccountDto;