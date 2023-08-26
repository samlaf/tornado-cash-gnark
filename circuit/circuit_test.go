package circuit

import (
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend"
	"github.com/consensys/gnark/test"
)

func TestCircuit(t *testing.T) {
	assert := test.NewAssert(t)

	var mimcCircuit Circuit

	{
		assert.ProverFailed(&mimcCircuit, &Circuit{
			PreImage: 42,
			Hash:     42,
		}, test.WithBackends(backend.GROTH16), test.WithCurves(ecc.BN254))
	}

	{
		assert.ProverSucceeded(&mimcCircuit, &Circuit{
			PreImage: 35,
			Hash:     "2474112249751028531650252582366798049474486386634137916759752348728204118534",
		}, test.WithBackends(backend.GROTH16), test.WithCurves(ecc.BN254))
	}
}
