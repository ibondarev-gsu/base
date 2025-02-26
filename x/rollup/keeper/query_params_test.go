package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/ibondarev-gsu/base/testutil/keeper"
	"github.com/ibondarev-gsu/base/x/rollup/types"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := keepertest.RollupKeeper(t)
	params := types.DefaultParams()
	require.NoError(t, keeper.SetParams(ctx, params))

	response, err := keeper.Params(ctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
