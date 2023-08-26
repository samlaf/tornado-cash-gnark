package circuit

import (
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/hash/mimc"
)

// Circuit defines a pre-image knowledge proof
// mimc(secret preImage) = public hash
type Circuit struct {
	PreImage frontend.Variable
	Hash     frontend.Variable `gnark:",public"`
}

// Define declares the circuit's constraints
func (circuit *Circuit) Define(api frontend.API) error {
	// hash function
	mimc, err := mimc.NewMiMC(api)
	if err != nil {
		return err
	}

	mimc.Write(circuit.PreImage)
	computedHash := mimc.Sum()

	// specify constraints
	// mimc(preImage) == hash
	api.AssertIsEqual(circuit.Hash, computedHash)

	return nil
}
