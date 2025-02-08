# Beep Events

This is a highlight of important beep events useful for auditors and off-chain intent executors.

## Create Intent Tranasction Log

- To get logs, run ```beepd query tx <TX_HASH>```

```sh
code: 0
codespace: ""
data: 12260A242F626565702E696E74656E742E4D7367437265617465496E74656E74526573706F6E7365
events:
- attributes:
  - index: true
    key: fee
    value: ""
  - index: true
    key: fee_payer
    value: beep17hm22tvhv0grtzg3qq9u2w037c0f4z08l74sr7
  type: tx
- attributes:
  - index: true
    key: acc_seq
    value: beep17hm22tvhv0grtzg3qq9u2w037c0f4z08l74sr7/6
  type: tx
- attributes:
  - index: true
    key: signature
    value: t00zWtsNH2OVUsI5qYYveDuhT93rfewoZn0ZQdGzTjBdfGhDwZmpcdyzpCBh7QXJTcVzbh+N/Vk3NSAMN7+4PQ==
  type: tx
- attributes:
  - index: true
    key: action
    value: /beep.intent.MsgCreateIntent
  - index: true
    key: sender
    value: beep17hm22tvhv0grtzg3qq9u2w037c0f4z08l74sr7
  - index: true
    key: module
    value: intent
  - index: true
    key: msg_index
    value: "0"
  type: message
- attributes:
  - index: true
    key: intent_id
    value: "4"
  - index: true
    key: creator
    value: beep17hm22tvhv0grtzg3qq9u2w037c0f4z08l74sr7
  - index: true
    key: amount
    value: "20"
  - index: true
    key: input_token
    value: token
  - index: true
    key: output_token
    value: stake
  - index: true
    key: target_chain
    value: beep
  - index: true
    key: min_output
    value: "2"
  - index: true
    key: status
    value: OPEN
  - index: true
    key: expiry_height
    value: "1739006417"
  - index: true
    key: msg_index
    value: "0"
  type: create_intent
- attributes:
  - index: true
    key: spender
    value: beep17hm22tvhv0grtzg3qq9u2w037c0f4z08l74sr7
  - index: true
    key: amount
    value: 20token
  - index: true
    key: msg_index
    value: "0"
  type: coin_spent
- attributes:
  - index: true
    key: receiver
    value: beep19q4uhsls5d9g5jkx7qxzwm70vm8nw4arm57dud
  - index: true
    key: amount
    value: 20token
  - index: true
    key: msg_index
    value: "0"
  type: coin_received
- attributes:
  - index: true
    key: recipient
    value: beep19q4uhsls5d9g5jkx7qxzwm70vm8nw4arm57dud
  - index: true
    key: sender
    value: beep17hm22tvhv0grtzg3qq9u2w037c0f4z08l74sr7
  - index: true
    key: amount
    value: 20token
  - index: true
    key: msg_index
    value: "0"
  type: transfer
- attributes:
  - index: true
    key: sender
    value: beep17hm22tvhv0grtzg3qq9u2w037c0f4z08l74sr7
  - index: true
    key: msg_index
    value: "0"
  type: message
- attributes:
  - index: true
    key: intent_id
    value: "4"
  - index: true
    key: sender
    value: beep17hm22tvhv0grtzg3qq9u2w037c0f4z08l74sr7
  - index: true
    key: receiver
    value: beep19q4uhsls5d9g5jkx7qxzwm70vm8nw4arm57dud
  - index: true
    key: denom
    value: token
  - index: true
    key: amount
    value: "20"
  - index: true
    key: msg_index
    value: "0"
  type: escrow_native_tokens
gas_used: "56387"
gas_wanted: "200000"
height: "2203"
info: ""
logs: []
raw_log: ""
timestamp: "2025-02-08T09:19:17Z"
tx:
  '@type': /cosmos.tx.v1beta1.Tx
  auth_info:
    fee:
      amount: []
      gas_limit: "200000"
      granter: ""
      payer: ""
    signer_infos:
    - mode_info:
        single:
          mode: SIGN_MODE_DIRECT
      public_key:
        '@type': /cosmos.crypto.secp256k1.PubKey
        key: AvPj7iSvqMG3TN8Z+qn+3MA5Bur+b/B9tD+i1b2JQoeH
      sequence: "6"
    tip: null
  body:
    extension_options: []
    memo: ""
    messages:
    - '@type': /beep.intent.MsgCreateIntent
      amount: 20
      creator: beep17hm22tvhv0grtzg3qq9u2w037c0f4z08l74sr7
      inputToken: token
      intentType: SWAP
      memo: test again
      minOutput: 2
      outputToken: stake
      targetChain: beep
    non_critical_extension_options: []
    timeout_height: "0"
  signatures:
  - t00zWtsNH2OVUsI5qYYveDuhT93rfewoZn0ZQdGzTjBdfGhDwZmpcdyzpCBh7QXJTcVzbh+N/Vk3NSAMN7+4PQ==
txhash: EED4CE7537FAAB728437F2302187626D3A031E5CDA8A6D2901A276B9105BCA40
```

- Create Intent Event definition is found in **beep/x/intent/keepers/msg_server_create_intent.go**
- Create Intent Event Type definition is found in **beep/x/intent/types/events.go**

## Accept Intent Tranaction Log

- To get logs, run ```beepd query tx <TX_HASH>```

```sh
code: 0
codespace: ""
data: 12260A242F626565702E696E74656E742E4D7367416363657074496E74656E74526573706F6E7365
events:
- attributes:
  - index: true
    key: fee
    value: ""
  - index: true
    key: fee_payer
    value: beep16zzglylfxvp2wuwa3hqeelr6hat5egvtjkfxpv
  type: tx
- attributes:
  - index: true
    key: acc_seq
    value: beep16zzglylfxvp2wuwa3hqeelr6hat5egvtjkfxpv/1
  type: tx
- attributes:
  - index: true
    key: signature
    value: E3z69zGxTJxJqJuuekghdWQVjJ92UnduY4EpRx70QB4CaMO440HVmFL7FgWY+DYpiKSX3/Pds9h5Sg0al9MPZw==
  type: tx
- attributes:
  - index: true
    key: action
    value: /beep.intent.MsgAcceptIntent
  - index: true
    key: sender
    value: beep16zzglylfxvp2wuwa3hqeelr6hat5egvtjkfxpv
  - index: true
    key: module
    value: intent
  - index: true
    key: msg_index
    value: "0"
  type: message
- attributes:
  - index: true
    key: intent_id
    value: "4"
  - index: true
    key: creator
    value: beep17hm22tvhv0grtzg3qq9u2w037c0f4z08l74sr7
  - index: true
    key: executor
    value: beep16zzglylfxvp2wuwa3hqeelr6hat5egvtjkfxpv
  - index: true
    key: input_token
    value: token
  - index: true
    key: output_token
    value: stake
  - index: true
    key: amount
    value: "20"
  - index: true
    key: min_output
    value: "2"
  - index: true
    key: msg_index
    value: "0"
  type: accept_intent
- attributes:
  - index: true
    key: spender
    value: beep16zzglylfxvp2wuwa3hqeelr6hat5egvtjkfxpv
  - index: true
    key: amount
    value: 20stake
  - index: true
    key: msg_index
    value: "0"
  type: coin_spent
- attributes:
  - index: true
    key: receiver
    value: beep19q4uhsls5d9g5jkx7qxzwm70vm8nw4arm57dud
  - index: true
    key: amount
    value: 20stake
  - index: true
    key: msg_index
    value: "0"
  type: coin_received
- attributes:
  - index: true
    key: recipient
    value: beep19q4uhsls5d9g5jkx7qxzwm70vm8nw4arm57dud
  - index: true
    key: sender
    value: beep16zzglylfxvp2wuwa3hqeelr6hat5egvtjkfxpv
  - index: true
    key: amount
    value: 20stake
  - index: true
    key: msg_index
    value: "0"
  type: transfer
- attributes:
  - index: true
    key: sender
    value: beep16zzglylfxvp2wuwa3hqeelr6hat5egvtjkfxpv
  - index: true
    key: msg_index
    value: "0"
  type: message
- attributes:
  - index: true
    key: intent_id
    value: "4"
  - index: true
    key: sender
    value: beep16zzglylfxvp2wuwa3hqeelr6hat5egvtjkfxpv
  - index: true
    key: receiver
    value: beep19q4uhsls5d9g5jkx7qxzwm70vm8nw4arm57dud
  - index: true
    key: denom
    value: stake
  - index: true
    key: amount
    value: "20"
  - index: true
    key: msg_index
    value: "0"
  type: executor_escrow
- attributes:
  - index: true
    key: spender
    value: beep19q4uhsls5d9g5jkx7qxzwm70vm8nw4arm57dud
  - index: true
    key: amount
    value: 20stake
  - index: true
    key: msg_index
    value: "0"
  type: coin_spent
- attributes:
  - index: true
    key: receiver
    value: beep17hm22tvhv0grtzg3qq9u2w037c0f4z08l74sr7
  - index: true
    key: amount
    value: 20stake
  - index: true
    key: msg_index
    value: "0"
  type: coin_received
- attributes:
  - index: true
    key: recipient
    value: beep17hm22tvhv0grtzg3qq9u2w037c0f4z08l74sr7
  - index: true
    key: sender
    value: beep19q4uhsls5d9g5jkx7qxzwm70vm8nw4arm57dud
  - index: true
    key: amount
    value: 20stake
  - index: true
    key: msg_index
    value: "0"
  type: transfer
- attributes:
  - index: true
    key: sender
    value: beep19q4uhsls5d9g5jkx7qxzwm70vm8nw4arm57dud
  - index: true
    key: msg_index
    value: "0"
  type: message
- attributes:
  - index: true
    key: intent_id
    value: "4"
  - index: true
    key: sender
    value: beep19q4uhsls5d9g5jkx7qxzwm70vm8nw4arm57dud
  - index: true
    key: receiver
    value: beep17hm22tvhv0grtzg3qq9u2w037c0f4z08l74sr7
  - index: true
    key: denom
    value: stake
  - index: true
    key: amount
    value: "20"
  - index: true
    key: msg_index
    value: "0"
  type: creator_payout
- attributes:
  - index: true
    key: spender
    value: beep19q4uhsls5d9g5jkx7qxzwm70vm8nw4arm57dud
  - index: true
    key: amount
    value: 20token
  - index: true
    key: msg_index
    value: "0"
  type: coin_spent
- attributes:
  - index: true
    key: receiver
    value: beep16zzglylfxvp2wuwa3hqeelr6hat5egvtjkfxpv
  - index: true
    key: amount
    value: 20token
  - index: true
    key: msg_index
    value: "0"
  type: coin_received
- attributes:
  - index: true
    key: recipient
    value: beep16zzglylfxvp2wuwa3hqeelr6hat5egvtjkfxpv
  - index: true
    key: sender
    value: beep19q4uhsls5d9g5jkx7qxzwm70vm8nw4arm57dud
  - index: true
    key: amount
    value: 20token
  - index: true
    key: msg_index
    value: "0"
  type: transfer
- attributes:
  - index: true
    key: sender
    value: beep19q4uhsls5d9g5jkx7qxzwm70vm8nw4arm57dud
  - index: true
    key: msg_index
    value: "0"
  type: message
- attributes:
  - index: true
    key: intent_id
    value: "4"
  - index: true
    key: sender
    value: beep19q4uhsls5d9g5jkx7qxzwm70vm8nw4arm57dud
  - index: true
    key: receiver
    value: beep16zzglylfxvp2wuwa3hqeelr6hat5egvtjkfxpv
  - index: true
    key: denom
    value: token
  - index: true
    key: amount
    value: "20"
  - index: true
    key: msg_index
    value: "0"
  type: executor_payout
- attributes:
  - index: true
    key: intent_id
    value: "4"
  - index: true
    key: creator
    value: beep17hm22tvhv0grtzg3qq9u2w037c0f4z08l74sr7
  - index: true
    key: executor
    value: beep16zzglylfxvp2wuwa3hqeelr6hat5egvtjkfxpv
  - index: true
    key: status
    value: EXECUTED
  - index: true
    key: msg_index
    value: "0"
  type: intent_executed
gas_used: "81347"
gas_wanted: "200000"
height: "2601"
info: ""
logs: []
raw_log: ""
timestamp: "2025-02-08T09:26:07Z"
tx:
  '@type': /cosmos.tx.v1beta1.Tx
  auth_info:
    fee:
      amount: []
      gas_limit: "200000"
      granter: ""
      payer: ""
    signer_infos:
    - mode_info:
        single:
          mode: SIGN_MODE_DIRECT
      public_key:
        '@type': /cosmos.crypto.secp256k1.PubKey
        key: A2mrfo9s4kr4chpxEK4rMKTtm/1Q7Sy32ggEv2IzynTc
      sequence: "1"
    tip: null
  body:
    extension_options: []
    memo: ""
    messages:
    - '@type': /beep.intent.MsgAcceptIntent
      executor: beep16zzglylfxvp2wuwa3hqeelr6hat5egvtjkfxpv
      id: "4"
    non_critical_extension_options: []
    timeout_height: "0"
  signatures:
  - E3z69zGxTJxJqJuuekghdWQVjJ92UnduY4EpRx70QB4CaMO440HVmFL7FgWY+DYpiKSX3/Pds9h5Sg0al9MPZw==
txhash: 31E9C7A29F8BD21004402277C5A2A4DE89416C79C42ACE26E490F717BF34D79B
```

- Create Intent Event definition is found in **beep/x/intent/keepers/msg_server_accept_intent.go**
- Create Intent Event Type definition is found in **beep/x/intent/types/events.go**