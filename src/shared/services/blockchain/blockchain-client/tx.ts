import { SigningStargateClient } from "@cosmjs/stargate";
import { MsgAcceptIntent, MsgBurnTokens, MsgCreateDenom, MsgCreateIntent, MsgMintTokens, MsgTransferTokens, MsgUpdateDenom } from "./types";

export class BeepTxClient {
    private signingClient: SigningStargateClient;

    constructor(signingClient: SigningStargateClient) {
        this.signingClient = signingClient;
    }

    async createIntent(
        creator: string,
        params: Omit<MsgCreateIntent, "creator">
    ) {
        const msg = {
            creator: creator,
            ...params,
            amount: params.amount.toString()
        };

        return await this.signingClient.signAndBroadcast(
            creator,
            [{
                typeUrl: "/beep.intent.MsgCreateIntent",
                value: msg
            }],
            "auto"
        );
    }

    async acceptIntent(
        executor: string,
        params: Omit<MsgAcceptIntent, "executor">
    ) {
        const msg = {
            executor: executor,
            ...params,
            id: params.id.toString()
        };

        return await this.signingClient.signAndBroadcast(
            executor,
            [{
                typeUrl: "/beep.intent.MsgAcceptIntent",
                value: msg
            }],
            "auto"
        );
    }

    async createDenom(
        creator: string,
        params: Omit<MsgCreateDenom, "owner">
    ) {
        const msg = {
            owner: creator,
            ...params,
            maxSupply: params.maxSupply.toString()
        };

        return await this.signingClient.signAndBroadcast(
            creator,
            [{
                typeUrl: "/beep.tokenfactory.MsgCreateDenom",
                value: msg
            }],
            "auto"
        );
    }

    async updateDenom(
        owner: string,
        params: Omit<MsgUpdateDenom, "owner">
    ) {
        const msg = {
            owner,
            ...params,
            maxSupply: params.maxSupply.toString()
        };

        return await this.signingClient.signAndBroadcast(
            owner,
            [{
                typeUrl: "/beep.tokenfactory.MsgUpdateDenom",
                value: msg
            }],
            "auto"
        );
    }

    async mintToken(
        owner: string,
        params: Omit<MsgMintTokens, "owner">
    ) {
        const msg = {
            owner,
            ...params,
            recipient: params.recipient.toString()
        };

        return await this.signingClient.signAndBroadcast(
            owner,
            [{
                typeUrl: "/beep.tokenfactory.MsgMintTokens",
                value: msg
            }],
            "auto"
        );
    }

    async transferToken(
        sender: string,
        params: Omit<MsgTransferTokens, "sender">
    ) {
        const msg = {
            sender,
            ...params,
            amount: params.amount.toString()
        };

        return await this.signingClient.signAndBroadcast(
            sender,
            [{
                typeUrl: "/beep.tokenfactory.MsgTransferTokens",
                value: msg
            }],
            "auto"
        );
    }

    async burnToken(
        signer: string,
        params: Omit<MsgBurnTokens, "signer">
    ) {
        const msg = {
            signer,
            ...params,
            denom: params.denom.toString()
        };

        return await this.signingClient.signAndBroadcast(
            signer,
            [{
                typeUrl: "/beep.tokenfactory.MsgBurnTokens",
                value: msg
            }],
            "auto"
        );
    }
}
