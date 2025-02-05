package keeper

import (
	"context"

	"beep/x/intent/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateIntent(goCtx context.Context, msg *types.MsgCreateIntent) (*types.MsgCreateIntentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateIntentResponse{}, nil
}
