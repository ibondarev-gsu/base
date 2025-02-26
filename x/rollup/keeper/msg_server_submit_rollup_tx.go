package keeper

import (
	"context"
	"encoding/binary"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ibondarev-gsu/base/x/rollup/types"
)

func (k msgServer) SubmitRollupTx(goCtx context.Context, msg *types.MsgSubmitRollupTx) (*types.MsgSubmitRollupTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Логируем полученные данные
	k.Logger().Info("Received Rollup Tx", "data", msg.Data)

	// Открываем KVStore через KVStoreService
	store := k.storeService.OpenKVStore(ctx)

	// Получаем текущий счетчик транзакций
	txCountBz, err := store.Get([]byte("tx_count"))
	if err != nil {
		return nil, err
	}

	var txCount int64
	if txCountBz == nil {
		txCount = 0
	} else {
		txCount = int64(binary.BigEndian.Uint64(txCountBz))
	}

	// Увеличиваем счетчик и сохраняем новую транзакцию
	txCount++
	err = store.Set([]byte(fmt.Sprintf("rollup_tx_%d", txCount)), []byte(msg.Data))
	if err != nil {
		return nil, err
	}

	// Сохраняем обновленный счетчик транзакций
	txCountBz = make([]byte, 8)
	binary.BigEndian.PutUint64(txCountBz, uint64(txCount))
	err = store.Set([]byte("tx_count"), txCountBz)
	if err != nil {
		return nil, err
	}

	return &types.MsgSubmitRollupTxResponse{}, nil
}
