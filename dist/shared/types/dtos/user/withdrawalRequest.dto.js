"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
class WithdrawalRequestDto {
    constructor(userwithdrwalRequest) {
        this.id = userwithdrwalRequest._id;
        this.userId = userwithdrwalRequest.userId;
        this.accountNumber = userwithdrwalRequest.accountNumber;
        this.bank = userwithdrwalRequest.bank;
        this.amount = userwithdrwalRequest.amount;
        this.status = userwithdrwalRequest.status;
        this.reference = userwithdrwalRequest.reference;
        this.updatedAt = userwithdrwalRequest.updatedAt;
        this.createdAt = userwithdrwalRequest.createdAt;
    }
    get getModel() {
        return {
            _id: this.id,
            userId: this.userId,
            accountNumber: this.accountNumber,
            bank: this.bank,
            amount: this.amount,
            status: this.status,
            reference: this.reference,
            updatedAt: this.updatedAt ? new Date(this.updatedAt) : undefined,
            createdAt: this.createdAt ? new Date(this.createdAt) : undefined,
        };
    }
}
exports.default = WithdrawalRequestDto;
