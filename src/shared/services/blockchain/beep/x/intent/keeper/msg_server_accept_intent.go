package keeper

import (
	"context"

	"beep/x/intent/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AcceptIntent(goCtx context.Context, msg *types.MsgAcceptIntent) (*types.MsgAcceptIntentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgAcceptIntentResponse{}, nil
}
