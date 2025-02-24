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
exports.WithdrawalRequest = void 0;
const mongoose_1 = require("mongoose");
const mongoose_paginate_v2_1 = __importDefault(require("mongoose-paginate-v2"));
const withdrawalRequest_dto_1 = __importDefault(require("../../../../types/dtos/user/withdrawalRequest.dto"));
const withdrawRequest_response_1 = require("../../../../types/interfaces/responses/user/withdrawRequest.response");
const WithdrawalRequestSchema = new mongoose_1.Schema({
    userId: {
        type: String,
        required: true
    },
    amount: {
        type: Number,
        required: true
    },
    bank: {
        type: String,
        required: true
    },
    accountNumber: {
        type: String,
        required: true
    },
    status: {
        type: String,
        enum: withdrawRequest_response_1.WithdrawalRequestStatus
    },
    reference: {
        type: String,
        required: true
    },
    updatedAt: {
        type: String
    },
    createdAt: {
        type: String
    },
});
WithdrawalRequestSchema.plugin(mongoose_paginate_v2_1.default);
exports.WithdrawalRequest = (0, mongoose_1.model)("WithrawalRequest", WithdrawalRequestSchema);
class WithdrawalRequestModel {
    constructor() {
        this.createWithrawalRequestToDB = (details) => __awaiter(this, void 0, void 0, function* () {
            try {
                const data = yield this.WithdrawalRequest.create(details);
                if (data) {
                    return { status: true, data: new withdrawalRequest_dto_1.default(data) };
                }
                else {
                    return { status: false, error: "Couldn't create request" };
                }
            }
            catch (error) {
                return { status: false, error };
            }
        });
        this.checkIfExist = (details) => __awaiter(this, void 0, void 0, function* () {
            try {
                const data = yield this.WithdrawalRequest.findOne(details);
                if (data) {
                    return { status: true, data: new withdrawalRequest_dto_1.default(data) };
                }
                else {
                    return { status: false, error: "No request find" };
                }
            }
            catch (error) {
                return { status: false, error };
            }
        });
        this.updateWithrawalRequest = (id, details) => __awaiter(this, void 0, void 0, function* () {
            try {
                const data = yield this.WithdrawalRequest.findOneAndUpdate({ _id: id }, details, { new: true });
                if (data) {
                    return { status: true, data: new withdrawalRequest_dto_1.default(data) };
                }
                else {
                    return { status: false, error: "Unable to update request" };
                }
            }
            catch (error) {
                return { status: false, error };
            }
        });
        this.WithdrawalRequest = exports.WithdrawalRequest;
    }
}
exports.default = WithdrawalRequestModel;
