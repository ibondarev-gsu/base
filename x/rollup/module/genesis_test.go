package rollup_test

import (
	"testing"

	keepertest "github.com/ibondarev-gsu/base/testutil/keeper"
	"github.com/ibondarev-gsu/base/testutil/nullify"
	rollup "github.com/ibondarev-gsu/base/x/rollup/module"
	"github.com/ibondarev-gsu/base/x/rollup/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RollupKeeper(t)
	rollup.InitGenesis(ctx, k, genesisState)
	got := rollup.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
