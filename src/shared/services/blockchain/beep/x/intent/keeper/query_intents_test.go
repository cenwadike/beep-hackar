package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "beep/testutil/keeper"
	"beep/testutil/nullify"
	"beep/x/intent/types"
)

func TestIntentsQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.IntentKeeper(t)
	msgs := createNIntents(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetIntentsRequest
		response *types.QueryGetIntentsResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetIntentsRequest{Id: msgs[0].Id},
			response: &types.QueryGetIntentsResponse{Intents: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetIntentsRequest{Id: msgs[1].Id},
			response: &types.QueryGetIntentsResponse{Intents: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetIntentsRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Intents(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestIntentsQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.IntentKeeper(t)
	msgs := createNIntents(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllIntentsRequest {
		return &types.QueryAllIntentsRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.IntentsAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Intents), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Intents),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.IntentsAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Intents), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Intents),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.IntentsAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Intents),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.IntentsAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
