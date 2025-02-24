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
Object.defineProperty(exports, "__esModule", { value: true });
exports.BeepTxClient = void 0;
class BeepTxClient {
    transfer(recipient, amount) {
        return __awaiter(this, void 0, void 0, function* () {
            return {
                transfer: {
                    recipient: recipient,
                    amount: amount,
                },
            };
        });
    }
    mint(recipient, amount) {
        return __awaiter(this, void 0, void 0, function* () {
            return {
                mint: {
                    recipient: recipient,
                    amount: amount,
                },
            };
        });
    }
    burn(amount) {
        return __awaiter(this, void 0, void 0, function* () {
            return {
                burn: {
                    amount: amount,
                },
            };
        });
    }
    burnFrom(owner, amount) {
        return __awaiter(this, void 0, void 0, function* () {
            return {
                burn_from: {
                    owner: owner,
                    amount: amount,
                },
            };
        });
    }
    balance(address) {
        return __awaiter(this, void 0, void 0, function* () {
            return {
                balance: {
                    address: address
                },
            };
        });
    }
}
exports.BeepTxClient = BeepTxClient;
