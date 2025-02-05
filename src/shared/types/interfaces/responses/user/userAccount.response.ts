export interface IUserAccount {
    _id?: string;
    phoneNumber: string;
    pin: string;
    publicKey: string;
    privateKey: string;
    balance: number;
    updatedAt?: Date;
    createdAt?: Date;
}