package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ibondarev-gsu/base/x/rollup/types"
)

func (k msgServer) RegisterVk(goCtx context.Context, msg *types.MsgRegisterVk) (*types.MsgRegisterVkResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRegisterVkResponse{}, nil
}
