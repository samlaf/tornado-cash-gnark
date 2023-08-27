package circuit

import (
	"math/big"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/test"
)

func TestCircuit(t *testing.T) {
	assert := test.NewAssert(t)

	circuit := InitializedCircuit()

	{
		assert.ProverFailed(
			&circuit,
			&Circuit{
				Nullifier:          1,
				Secret:             2,
				Commitment:         3,
				NullifierHash:      4,
				MerkleRootHash:     5,
				MerklePathElements: []frontend.Variable{6, 7, 8, 9},
				LeafIdx:            0,
			},
			test.WithBackends(backend.GROTH16), test.WithCurves(ecc.BN254))
	}

	{
		assert.ProverSucceeded(
			&circuit,
			NewAssignment(big.NewInt(1), big.NewInt(2)),
			test.WithBackends(backend.GROTH16), test.WithCurves(ecc.BN254))
	}
}
