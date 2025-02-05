package keeper

import (
	"context"

	"beep/x/intent/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) IntentsAll(ctx context.Context, req *types.QueryAllIntentsRequest) (*types.QueryAllIntentsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var intentss []types.Intents

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	intentsStore := prefix.NewStore(store, types.KeyPrefix(types.IntentsKey))

	pageRes, err := query.Paginate(intentsStore, req.Pagination, func(key []byte, value []byte) error {
		var intents types.Intents
		if err := k.cdc.Unmarshal(value, &intents); err != nil {
			return err
		}

		intentss = append(intentss, intents)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllIntentsResponse{Intents: intentss, Pagination: pageRes}, nil
}

func (k Keeper) Intents(ctx context.Context, req *types.QueryGetIntentsRequest) (*types.QueryGetIntentsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	intents, found := k.GetIntents(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetIntentsResponse{Intents: intents}, nil
}
