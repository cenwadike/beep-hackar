export interface IUserAccount {
    _id?: string;
    phoneNumber: string;
    pin: string;
    publicKey: string;
    privateKey: string;
    updatedAt?: Date;
    createdAt?: Date;
}