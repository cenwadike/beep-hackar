# Working with tokens

This guide will walk you through setting up the intent based blockchains, `beep`.

This guide demonstrate how to create denom, update denom, mint tokens, transfer tokens and burn tokens.


## Step 1: Build blockchain

```sh
    ignite chain build
```

## Step 2: Start blockchain with default config

```sh
    ignite chain serve
```

## Step 3: Create denom

```sh
    beepd tx tokenfactory create-denom "bATOM" "Beep ATOM" "bATOM" 6 "url" 1000000 true --from bob
```

Accept transaction to sign and submit the transaction.

## Step 3: Update denom

```sh
    beepd tx tokenfactory update-denom "bATOM" "Beep Protocol ATOM" "new_url" 1000000000 true --from bob
```

Accept transaction to sign and submit the transaction.

## Step 4: Mint Tokens

```sh
    beepd tx tokenfactory mint-tokens "bATOM" 100 beep1... --from bob
```

## Step 4: Transfer Tokens

```sh
    beepd tx tokenfactory transfer-tokens "bATOM" beep1... 10 --from bob
```

## Step 5: Burn Tokens

```sh
    beepd tx tokenfactory burn-tokens "bATOM" 10 --from bob
```
