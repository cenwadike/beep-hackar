# Setting Up Hermes Relayer for IBC

This guide will walk you through setting up the Hermes relayer for IBC between two chains, `beep` and `boop`.

## Step 1: Delete previous Hermes configuration

```sh
    rm -rf ~/.ignite/relayer
```

## Step 2: Copy Hermes configuration

```sh
    cp ./config.toml ~/.hermes/config.toml
```

## Step 3: Define Configuration Files

First, we need to modify the configuration files for both chains.

### Beep Configuration (`beep.yml`)

```yaml
    version: 1
    build:
    proto:
        path: proto
        third_party_paths:
        - third_party/proto
        - proto_vendor
    accounts:
    - name: alice
    coins:
    - 1000token
    - 1000000000stake
    - name: bob
    coins:
    - 500token
    - 100000000stake
    faucet:
    name: bob
    coins:
    - 5token
    - 100000stake
    host: :4501
    genesis:
    chain_id: beep
    validators:
    - name: alice
    bonded: 100000000stake
    app:
        api:
        address: :1318
        grpc:
        address: :9092
        grpc-web:
        address: :9093
    config:
        p2p:
        laddr: :26658
        rpc:
        laddr: :26659
        pprof_laddr: :6061
    home: $HOME/.beep
  ```

### Boop Configuration (`boop.yml`)

```yaml
    version: 1
    build:
    proto:
        path: proto
        third_party_paths:
        - third_party/proto
        - proto_vendor
    accounts:
    - name: alice
    coins:
    - 1000token
    - 1000000000stake
    - name: bob
    coins:
    - 500token
    - 100000000stake
    faucet:
    name: bob
    coins:
    - 5token
    - 100000stake
    host: :4501
    genesis:
    chain_id: boop
    validators:
    - name: alice
    bonded: 100000000stake
    app:
        api:
        address: :1318
        grpc:
        address: :9092
        grpc-web:
        address: :9093
    config:
        p2p:
        laddr: :26658
        rpc:
        laddr: :26659
        pprof_laddr: :6061
    home: $HOME/.boop
```

## Step 4: Configure the Relayer Connection

Run the following command to configure the relayer connection:

```sh
    ignite relayer hermes configure "boop" "http://localhost:26657" "http://localhost:9090" "beep" "http://localhost:26659" "http://localhost:9092" --chain-a-faucet "http://0.0.0.0:4500" --chain-b-faucet "http://0.0.0.0:4501" --chain-a-port-id "intent" --chain-b-port-id "intent" --channel-version "intent-1" --chain-a-account-prefix "beep" --chain-b-account-prefix "beep"
```

![alt text](<./images/Screenshot 2025-02-07 at 09.04.59.png>)

## Step 5: Create the Channel

Use the following command to create the channel:

```sh
    hermes create channel --a-chain beep --b-chain boop --a-port "intent" --b-port "intent" --order ordered --chan-version "intent-1" --new-client-connection
```

## Send IBC transaction

```sh
    beepd tx intent send-intent-packet "intent" "channel-1" "swap" "test" beep 0 --from alice --chain-id boop --home ~/.boop
```
![alt text](<./images/Screenshot 2025-02-07 at 09.29.29.png>)

- Verify transaction was successful with `code 0`

![alt text](<./images/Screenshot 2025-02-07 at 09.29.29.png>)