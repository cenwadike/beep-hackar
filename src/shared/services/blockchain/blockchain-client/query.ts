import { QueryClient } from "@cosmjs/stargate";
import { Denom } from "./types";

export class BeepQueryClient {
    private queryClient: QueryClient;

    constructor(queryClient: QueryClient) {
        this.queryClient = queryClient;
    }

    async getDenom(denom: string): Promise<Denom> {
        const path = "/beep.tokenfactory.Query/Denom";
        const request = new Uint8Array(Buffer.from(JSON.stringify({ denom })));
        const response = await this.queryClient.queryAbci(path, request);
        return JSON.parse(new TextDecoder().decode(response.value));
    }

    async listDenoms(): Promise<Denom[]> {
        const path = "/beep.tokenfactory.Query/DenomAll";
        const request = new Uint8Array();
        const response = await this.queryClient.queryAbci(path, request);
        return JSON.parse(new TextDecoder().decode(response.value)).denoms;
    }

    async getIntent(id: string): Promise<Denom> {
        const path = "/beep.intent.Query/Intents";
        const request = new Uint8Array(Buffer.from(JSON.stringify({ id })));
        const response = await this.queryClient.queryAbci(path, request);
        return JSON.parse(new TextDecoder().decode(response.value));
    }

    async listIntents(): Promise<Denom[]> {
        const path = "/beep.intent.Query/IntentsAll";
        const request = new Uint8Array();
        const response = await this.queryClient.queryAbci(path, request);
        return JSON.parse(new TextDecoder().decode(response.value)).intents;
    }
}
