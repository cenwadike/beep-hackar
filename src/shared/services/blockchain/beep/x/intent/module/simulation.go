package intent

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"beep/testutil/sample"
	intentsimulation "beep/x/intent/simulation"
	"beep/x/intent/types"
)

// avoid unused import issue
var (
	_ = intentsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateIntent = "op_weight_msg_create_intent"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateIntent int = 100

	opWeightMsgAcceptIntent = "op_weight_msg_accept_intent"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAcceptIntent int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	intentGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&intentGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateIntent int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateIntent, &weightMsgCreateIntent, nil,
		func(_ *rand.Rand) {
			weightMsgCreateIntent = defaultWeightMsgCreateIntent
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateIntent,
		intentsimulation.SimulateMsgCreateIntent(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAcceptIntent int
	simState.AppParams.GetOrGenerate(opWeightMsgAcceptIntent, &weightMsgAcceptIntent, nil,
		func(_ *rand.Rand) {
			weightMsgAcceptIntent = defaultWeightMsgAcceptIntent
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAcceptIntent,
		intentsimulation.SimulateMsgAcceptIntent(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateIntent,
			defaultWeightMsgCreateIntent,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				intentsimulation.SimulateMsgCreateIntent(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgAcceptIntent,
			defaultWeightMsgAcceptIntent,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				intentsimulation.SimulateMsgAcceptIntent(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
