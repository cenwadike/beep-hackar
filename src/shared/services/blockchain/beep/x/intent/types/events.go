package types

// Event types
const (
	EventTypeCreateIntent        = "create_intent"
	EventTypeEscrowIBCTransfer   = "escrow_ibc_transfer"
	EventTypeEscrowNativeTokens  = "escrow_native_tokens"
	EventTypeIBCTransferFailure  = "ibc_transfer_failure"
	EventTypeAcceptIntent        = "accept_intent"
	EventTypeAcceptIntentFailure = "accept_intent_failure"
	EventTypeExecutorEscrow      = "executor_escrow"
	EventTypeCreatorPayout       = "creator_payout"
	EventTypeExecutorPayout      = "executor_payout"
	EventTypeIntentExecuted      = "intent_executed"
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
	AttributeKeyExecutor     = "executor"
	AttributeKeyStage        = "stage"
)
