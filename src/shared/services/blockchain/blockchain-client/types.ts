export interface MsgCreateIntent {
    creator: string;
    inputToken: string;
    outputToken: string;
    intentType: string;
    memo: string;
    targetChain: string;
    minOutput: number;
    amount: number;
}

export interface MsgAcceptIntent {
    executor: string;
    id: number;
}

export interface Denom {
    owner: string;
    denom: string;
    description: string;
    ticker: string;
    precision: number;
    url: string;
    maxSupply: string;
    supply: string;
    canChangeMaxSupply: boolean;
}

export interface MsgCreateDenom {
    owner: string;
    denom: string;
    description: string;
    ticker: string;
    precision: number;
    url: string;
    maxSupply: string;
    canChangeMaxSupply: boolean;
}

export interface MsgUpdateDenom {
    owner: string;
    denom: string;
    description: string;
    maxSupply: string;
    canChangeMaxSupply: boolean;
}

export interface MsgMintTokens {
    owner: string;
    denom: string;
    amount: number;
    recipient: string;
}

export interface MsgTransferTokens {
    denom: string;
    sender: string;
    recipient: string;
    amount: number;
}

export interface MsgBurnTokens {
    signer: string;
    amount: number;
    denom: string;
}

// src/types/events.ts
export interface DenomCreatedEvent {
    owner: string;
    denom: string;
    description: string;
    ticker: string;
    precision: number;
    maxSupply: string;
    canChangeMaxSupply: boolean;
}

export interface DenomUpdatedEvent {
    owner: string;
    denom: string;
    description: string;
    maxSupply: string;
    canChangeMaxSupply: boolean;
}