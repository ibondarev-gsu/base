package rollup

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/ibondarev-gsu/base/testutil/sample"
	rollupsimulation "github.com/ibondarev-gsu/base/x/rollup/simulation"
	"github.com/ibondarev-gsu/base/x/rollup/types"
)

// avoid unused import issue
var (
	_ = rollupsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgSubmitRollupTx = "op_weight_msg_submit_rollup_tx"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubmitRollupTx int = 100

	opWeightMsgRegisterVk = "op_weight_msg_register_vk"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterVk int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	rollupGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&rollupGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgSubmitRollupTx int
	simState.AppParams.GetOrGenerate(opWeightMsgSubmitRollupTx, &weightMsgSubmitRollupTx, nil,
		func(_ *rand.Rand) {
			weightMsgSubmitRollupTx = defaultWeightMsgSubmitRollupTx
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubmitRollupTx,
		rollupsimulation.SimulateMsgSubmitRollupTx(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRegisterVk int
	simState.AppParams.GetOrGenerate(opWeightMsgRegisterVk, &weightMsgRegisterVk, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterVk = defaultWeightMsgRegisterVk
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterVk,
		rollupsimulation.SimulateMsgRegisterVk(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgSubmitRollupTx,
			defaultWeightMsgSubmitRollupTx,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				rollupsimulation.SimulateMsgSubmitRollupTx(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgRegisterVk,
			defaultWeightMsgRegisterVk,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				rollupsimulation.SimulateMsgRegisterVk(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
