package keeper

import (
	"bytes"
	"context"
	"cosmossdk.io/errors"
	"fmt"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/backend/witness"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ibondarev-gsu/base/x/rollup/types"
)

func (k Keeper) LoadVerificationRegisterKey(goCtx context.Context) ([]byte, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)
	vk, _ := store.Get([]byte("vk"))
	if vk == nil {
		return nil, errors.New("verification key not found", 10, "")
	}
	return vk, nil
}

// VerifyProof проверяет ZK-Proof
func (k Keeper) VerifyProof(goCtx sdk.Context, proof []byte, txData []byte) error {
	// 1. Загружаем Verification Key (VK)
	ctx := sdk.UnwrapSDKContext(goCtx)
	vkBytes, err := k.LoadVerificationRegisterKey(ctx)
	if err != nil {
		return err
	}

	// 2. Восстанавливаем верификационный ключ
	var vk groth16.VerifyingKey

	if _, err := vk.UnsafeReadFrom(bytes.NewReader(vkBytes)); err != nil {
		return fmt.Errorf("failed to unmarshal VK: %w", err)
	}

	// 3. Загружаем доказательство
	var proofStruct groth16.Proof
	if _, err := proofStruct.ReadFrom(bytes.NewReader(proof)); err != nil {
		return fmt.Errorf("failed to unmarshal proof: %w", err)
	}

	// 4. Формируем публичные входные данные
	publicInputs := struct {
		TxData frontend.Variable `gnark:",public"`
	}{
		TxData: txData,
	}

	// Формируем publicWitness
	publicWitness, err := witness.New(&publicInputs, groth16.CurveID)
	if err != nil {
		return nil, fmt.Errorf("failed to create public witness: %w", err)
	}

	// 5. Проверяем proof
	if err := groth16.Verify(proofStruct, vk, publicWitness); err != nil {
		return fmt.Errorf("invalid zero-knowledge proof: %w", err)
	}

	return nil
}

// Имитация проверки доказательства (замени на реальную ZK-библиотеку)
func zkVerify(vk []byte, proof []byte, txData []byte) bool {
	// Здесь должна быть реальная ZK-проверка (например, через gnark, circom, или Halo2)
	// Сейчас просто заглушка для примера
	return len(proof) > 0 && len(vk) > 0 && len(txData) > 0
}
