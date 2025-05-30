package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "beep/testutil/keeper"
	"beep/x/intent/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.IntentKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
