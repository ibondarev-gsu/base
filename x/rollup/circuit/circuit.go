package circuit

import (
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/hash/mimc"
)

// RollupCircuit описывает ZK-циркуит
type RollupCircuit struct {
	TxData frontend.Variable `gnark:",public"` // Хеш транзакции (публичный)
	Secret frontend.Variable // Секретный ввод (приватный)
}

// Define — основная логика циркуита
func (circuit *RollupCircuit) Define(api frontend.API) error {
	// Хешируем секретный ввод
	hasher, _ := mimc.NewMiMC(api)
	hasher.Write(circuit.Secret)
	computedHash := hasher.Sum()

	// Проверяем, что хеш равен публичному TxData
	api.AssertIsEqual(computedHash, circuit.TxData)

	return nil
}
