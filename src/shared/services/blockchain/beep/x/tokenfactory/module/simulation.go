package tokenfactory

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"beep/testutil/sample"
	tokenfactorysimulation "beep/x/tokenfactory/simulation"
	"beep/x/tokenfactory/types"
)

// avoid unused import issue
var (
	_ = tokenfactorysimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateDenom = "op_weight_msg_denom"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateDenom int = 100

	opWeightMsgUpdateDenom = "op_weight_msg_denom"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateDenom int = 100

	opWeightMsgDeleteDenom = "op_weight_msg_denom"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteDenom int = 100

	opWeightMsgMintTokens = "op_weight_msg_mint_tokens"
	// TODO: Determine the simulation weight value
	defaultWeightMsgMintTokens int = 100

	opWeightMsgTransferTokens = "op_weight_msg_transfer_tokens"
	// TODO: Determine the simulation weight value
	defaultWeightMsgTransferTokens int = 100

	opWeightMsgBurnTokens = "op_weight_msg_burn_tokens"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBurnTokens int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	tokenfactoryGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		DenomList: []types.Denom{
			{
				Owner: sample.AccAddress(),
				Denom: "0",
			},
			{
				Owner: sample.AccAddress(),
				Denom: "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&tokenfactoryGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateDenom int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateDenom, &weightMsgCreateDenom, nil,
		func(_ *rand.Rand) {
			weightMsgCreateDenom = defaultWeightMsgCreateDenom
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateDenom,
		tokenfactorysimulation.SimulateMsgCreateDenom(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateDenom int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateDenom, &weightMsgUpdateDenom, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateDenom = defaultWeightMsgUpdateDenom
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateDenom,
		tokenfactorysimulation.SimulateMsgUpdateDenom(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgMintTokens int
	simState.AppParams.GetOrGenerate(opWeightMsgMintTokens, &weightMsgMintTokens, nil,
		func(_ *rand.Rand) {
			weightMsgMintTokens = defaultWeightMsgMintTokens
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMintTokens,
		tokenfactorysimulation.SimulateMsgMintTokens(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgTransferTokens int
	simState.AppParams.GetOrGenerate(opWeightMsgTransferTokens, &weightMsgTransferTokens, nil,
		func(_ *rand.Rand) {
			weightMsgTransferTokens = defaultWeightMsgTransferTokens
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgTransferTokens,
		tokenfactorysimulation.SimulateMsgTransferTokens(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgBurnTokens int
	simState.AppParams.GetOrGenerate(opWeightMsgBurnTokens, &weightMsgBurnTokens, nil,
		func(_ *rand.Rand) {
			weightMsgBurnTokens = defaultWeightMsgBurnTokens
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBurnTokens,
		tokenfactorysimulation.SimulateMsgBurnTokens(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateDenom,
			defaultWeightMsgCreateDenom,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				tokenfactorysimulation.SimulateMsgCreateDenom(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateDenom,
			defaultWeightMsgUpdateDenom,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				tokenfactorysimulation.SimulateMsgUpdateDenom(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgMintTokens,
			defaultWeightMsgMintTokens,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				tokenfactorysimulation.SimulateMsgMintTokens(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgTransferTokens,
			defaultWeightMsgTransferTokens,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				tokenfactorysimulation.SimulateMsgTransferTokens(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgBurnTokens,
			defaultWeightMsgBurnTokens,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				tokenfactorysimulation.SimulateMsgBurnTokens(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
