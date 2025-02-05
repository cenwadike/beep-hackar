package keeper_test

import (
	"context"
	"testing"

	keepertest "beep/testutil/keeper"
	"beep/testutil/nullify"
	"beep/x/intent/keeper"
	"beep/x/intent/types"

	"github.com/stretchr/testify/require"
)

func createNIntents(keeper keeper.Keeper, ctx context.Context, n int) []types.Intents {
	items := make([]types.Intents, n)
	for i := range items {
		items[i].Id = keeper.AppendIntents(ctx, items[i])
	}
	return items
}

func TestIntentsGet(t *testing.T) {
	keeper, ctx := keepertest.IntentKeeper(t)
	items := createNIntents(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetIntents(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestIntentsRemove(t *testing.T) {
	keeper, ctx := keepertest.IntentKeeper(t)
	items := createNIntents(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveIntents(ctx, item.Id)
		_, found := keeper.GetIntents(ctx, item.Id)
		require.False(t, found)
	}
}

func TestIntentsGetAll(t *testing.T) {
	keeper, ctx := keepertest.IntentKeeper(t)
	items := createNIntents(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllIntents(ctx)),
	)
}

func TestIntentsCount(t *testing.T) {
	keeper, ctx := keepertest.IntentKeeper(t)
	items := createNIntents(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetIntentsCount(ctx))
}
