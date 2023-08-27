package circuit

import (
	"math/big"

	"github.com/consensys/gnark/frontend"
	snarkmimc "github.com/consensys/gnark/std/hash/mimc"

	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
)

// Commitment is stored in the merkle tree for each deposit
// NullifierHash is to make sure that the same deposit commitment is not used for more than one withdrawal
type Circuit struct {
	// Input
	Nullifier frontend.Variable // k
	Secret    frontend.Variable // r
	// Output
	NullifierHash frontend.Variable `gnark:",public"` // H(k)
	Commitment    frontend.Variable `gnark:",public"` // H(k, r)
	// we also store these to send them to the contract
	nullifierHash *big.Int
	commitment    *big.Int
}

func NewCircuit(nullifier, secret *big.Int) *Circuit {
	nullifierHash := calculateMimcOutput(nullifier)
	commitment := calculateMimcOutput(nullifier, secret)
	assignment := &Circuit{
		Nullifier:     nullifier,
		Secret:        secret,
		NullifierHash: nullifierHash,
		Commitment:    commitment,
		nullifierHash: nullifierHash,
		commitment:    commitment,
	}
	return assignment
}

func (c *Circuit) GetPublicOutputs() (*big.Int, *big.Int) {
	return c.nullifierHash, c.commitment
}

// Define declares the circuit's constraints
func (c *Circuit) Define(api frontend.API) error {
	// hash function
	mimc, err := snarkmimc.NewMiMC(api)
	if err != nil {
		return err
	}

	mimc.Write(c.Nullifier)
	computedNulliferHash := mimc.Sum()
	api.AssertIsEqual(c.NullifierHash, computedNulliferHash)

	mimc.Reset()
	mimc.Write(c.Nullifier, c.Secret)
	computedCommitment := mimc.Sum()
	api.AssertIsEqual(c.Commitment, computedCommitment)

	return nil
}

func calculateMimcOutput(inputNum ...*big.Int) *big.Int {
	hash := mimc.NewMiMC()

	// hash the input (written as big endian byte array)
	var input fr.Element
	for _, num := range inputNum {
		input.SetBigInt(num)
		inputBytes32 := input.Bytes()
		hash.Write(inputBytes32[:])
	}

	// get the output (which is a 32 byte slice)
	outputBytes := hash.Sum(nil)

	// write as a field element to get a nice string representation to output
	element := fr.Element{}
	element.SetBytes(outputBytes)
	return element.BigInt(big.NewInt(0))
}
