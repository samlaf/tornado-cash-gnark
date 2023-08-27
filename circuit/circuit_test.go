package circuit

import (
	"math/big"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend"
	"github.com/consensys/gnark/test"
)

func TestCircuit(t *testing.T) {
	assert := test.NewAssert(t)

	var mimcCircuit Circuit

	{
		assert.ProverFailed(
			&mimcCircuit,
			&Circuit{
				Nullifier:     1,
				Secret:        2,
				NullifierHash: 3,
				Commitment:    4,
			},
			test.WithBackends(backend.GROTH16), test.WithCurves(ecc.BN254))
	}

	{
		assert.ProverSucceeded(
			&mimcCircuit,
			NewCircuit(big.NewInt(1), big.NewInt(2)),
			test.WithBackends(backend.GROTH16), test.WithCurves(ecc.BN254))
	}
}
