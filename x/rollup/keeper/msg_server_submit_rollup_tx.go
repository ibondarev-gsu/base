package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ibondarev-gsu/base/x/rollup/types"
)

func (k msgServer) SubmitRollupTx(goCtx context.Context, msg *types.MsgSubmitRollupTx) (*types.MsgSubmitRollupTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSubmitRollupTxResponse{}, nil
}
