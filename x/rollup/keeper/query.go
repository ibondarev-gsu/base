package keeper

import (
	"github.com/ibondarev-gsu/base/x/rollup/types"
)

var _ types.QueryServer = Keeper{}
