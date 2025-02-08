# Creating Intents

This guide will walk you through setting up the intent based blockchains, `beep`.

This guide demonstrate how to create and accept intents and verify that the transaction was successful.

## Step 1: Build blockchain

```sh
    ignite chain build
```

## Step 2: Start blockchain with default config

```sh
    ignite chain serve
```

## Step 3: Create intent transaction

```sh
    beepd tx intent create-intent token 20 stake 2 "SWAP" "beep" "test again" --from bob
```

Accept transaction to sign and submit the transaction.

## Step 4: Verify intent transaction was successful

```sh
    beepd query tx <TX_HASH>
```

## Step 5: Check intent exist in storage

```sh
    beepd query intent list-intents --chain-id "beep"
```

## Step 6: Accept intent

```sh
    beepd tx intent accept-intent <INTENT_ID> --from alice
```

## Step 7: Verify intent transaction was successful

```sh
    beepd query tx <TX_HASH>
```

## Step 8: Confirm status change

```sh
    beepd query intent list-intents --chain-id "beep"
```
