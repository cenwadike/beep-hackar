package types

// Event types
const (
	EventTypeCreateIntent       = "create_intent"
	EventTypeEscrowIBCTransfer  = "ibc_transfer"
	EventTypeEscrowNativeTokens = "escrow_tokens"
	EventTypeIBCTransferFailure = "ibc_transfer_failure"
)

// Attribute keys
const (
	AttributeKeyIntentId     = "intent_id"
	AttributeKeyCreator      = "creator"
	AttributeKeyAmount       = "amount"
	AttributeKeyInputToken   = "input_token"
	AttributeKeyOutputToken  = "output_token"
	AttributeKeyTargetChain  = "target_chain"
	AttributeKeyMinOutput    = "min_output"
	AttributeKeyStatus       = "status"
	AttributeKeyExpiryHeight = "expiry_height"
	AttributeKeySender       = "sender"
	AttributeKeyReceiver     = "receiver"
	AttributeKeyDenom        = "denom"
	AttributeKeyIBCDenom     = "ibc_denom"
	AttributeKeyError        = "error"
)
