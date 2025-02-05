package intent_test

import (
	"testing"

	keepertest "beep/testutil/keeper"
	"beep/testutil/nullify"
	intent "beep/x/intent/module"
	"beep/x/intent/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		IntentsList: []types.Intents{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		IntentsCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IntentKeeper(t)
	intent.InitGenesis(ctx, k, genesisState)
	got := intent.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	require.ElementsMatch(t, genesisState.IntentsList, got.IntentsList)
	require.Equal(t, genesisState.IntentsCount, got.IntentsCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
