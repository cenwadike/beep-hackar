import { Schema, model, PaginateModel } from "mongoose";
import mongoosePaginate from 'mongoose-paginate-v2'
import { IUserAccount } from "../../../../types/interfaces/responses/user/userAccount.response";
import IUserAccountModel from "./type";
import UserAccountDto from "../../../../types/dtos/user/userAccount.dto";


const UserAccountSchema = new Schema<IUserAccount>({
    phoneNumber: {
      type: String,
      unique: true,
      required: true
    },
    pin: {
      type: String
    },
    publicKey: {
        type: String,
    },
    privateKey: {
        type: String,
    },
    updatedAt: {
      type: String
    },
    createdAt: {
      type: String
    },
  });
  
  UserAccountSchema.plugin(mongoosePaginate);
  
  export const UserAccount = model<IUserAccount, PaginateModel<IUserAccount>>("UserAccount", UserAccountSchema)

  class  UserAccountModel implements  IUserAccountModel {
    UserAccount: typeof UserAccount;
    
    constructor() {
        this.UserAccount =  UserAccount;
    }

    createAccountToDB = async (details: Partial<IUserAccount>) => {
        try {
            const data = await this.UserAccount.create(details);
            if (data) {
              return {status: true, data: new UserAccountDto(data)};
            } else {
              return {status: false, error: "Couldn't create account"};
            }
        } catch (error) {
            return {status: false, error };
        }
    }

    checkIfExist = async (details: Partial<IUserAccount>) => {
        try {
            const data = await this.UserAccount.findOne(details);
            if (data) {
              return {status: true, data: new UserAccountDto(data)};
            } else {
              return {status: false, error: "no account find"};
            }
        } catch (error) {
            return {status: false, error };
        }
    }

    updateAccount = async (phoneNumber: string, details: Partial<IUserAccount>) => {
      try {
          const data = await this.UserAccount.findOneAndUpdate({phoneNumber}, details, {new: true});
          if (data) {
            return {status: true, data: new UserAccountDto(data)};
          } else {
            return {status: false, error: "Unable to update configuration"};
          }
      } catch (error) {
          return {status: false, error };
      }
  }

}

export default UserAccountModel;