package types

// Event types
const (
	EventTypeCreateDenom    = "create_denom"
	EventTypeUpdateDenom    = "update_denom"
	EventTypeMintTokens     = "mint_tokens"
	EventTypeBurnTokens     = "burn_tokens"
	EventTypeTransferTokens = "transfer_tokens"
)

// Attribute keys
const (
	AttributeKeyDenomName               = "denom_name"
	AttributeKeyDenomOwner              = "denom_owner"
	AttributeKeyDenomDescription        = "denom_description"
	AttributeKeyDenomTicker             = "denom_ticker"
	AttributeKeyDenomPrecision          = "denom_precision"
	AttributeKeyDenomMaxSupply          = "denom_max_supply"
	AttributeKeyDenomCanChangeMaxSupply = "denom_can_change_max_supply"

	AttributeKeySender   = "sender"
	AttributeKeyReceiver = "receiver"

	AttributeKeyBurner      = "burner"
	AttributeKeyDenomAmount = "amount"
)
