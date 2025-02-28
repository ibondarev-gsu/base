package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ibondarev-gsu/base/x/rollup/types"
)

func (k msgServer) RegisterVk(goCtx context.Context, msg *types.MsgRegisterVk) (*types.MsgRegisterVkResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Сохраняем VK в хранилище
	store := k.storeService.OpenKVStore(ctx)
	err := store.Set([]byte("vk"), []byte(msg.Vk))
	if err != nil {
		k.Logger().Error("Key go", err)
		return nil, err
	}

	return &types.MsgRegisterVkResponse{}, nil
}
