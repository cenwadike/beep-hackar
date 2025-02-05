package simulation

import (
	"math/rand"

	"beep/x/intent/keeper"
	"beep/x/intent/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgAcceptIntent(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgAcceptIntent{
			Executor: simAccount.Address.String(),
		}

		// TODO: Handling the AcceptIntent simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "AcceptIntent simulation not implemented"), nil, nil
	}
}
