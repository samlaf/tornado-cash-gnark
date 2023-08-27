package circuit

import (
	"bytes"
	"math"
	"math/big"

	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/accumulator/merkle"
	snarkmimc "github.com/consensys/gnark/std/hash/mimc"

	"github.com/consensys/gnark-crypto/accumulator/merkletree"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
)

// hardcode small tree for now while developing
const DEPTH = 2

var NUM_LEAVES = int(math.Pow(2, DEPTH))

// Commitment is stored in the merkle tree for each deposit
// NullifierHash is to make sure that the same deposit commitment is not used for more than one withdrawal
// Note that circuit is hardcoded for bn254 (both by mimc used and curve mod in tree elements)
type Circuit struct {
	// ====== private inputs =======
	Nullifier          frontend.Variable   // k
	Secret             frontend.Variable   // r
	MerklePathElements []frontend.Variable // path elements from the merkle tree
	// leafIdx should be between 0 and 2^depth-1
	LeafIdx frontend.Variable // index of leaf node for which we are proving membership

	// ====== public inputs =======
	NullifierHash  frontend.Variable `gnark:",public"` // H(k)
	Commitment     frontend.Variable `gnark:",public"` // H(k, r)
	MerkleRootHash frontend.Variable `gnark:",public"` // root of the merkle tree

	// we also store the public inputs as big.Ints so that we can return them when generating
	// the call to verifier smart contract
	nullifierHash  *big.Int
	commitment     *big.Int
	merkleRootHash *big.Int
}

// InitializedCircuit returns a circuit with all variables initialized
// this is only relevant for array variables
// Not sure why this is needed but calls to frontend.Compile always take an unassigned circuit
// and an assignment (NewCircuit) is only later computed before calling frontend.NewWitness
func InitializedCircuit() Circuit {
	var circuit Circuit
	circuit.MerklePathElements = make([]frontend.Variable, DEPTH+1)
	return circuit
}

func NewAssignment(nullifier, secret *big.Int) *Circuit {
	nullifierHash := calculateMimcOutput(nullifier)
	commitment := calculateMimcOutput(nullifier, secret)

	mod := ecc.BN254.ScalarField()
	modNumBytes := len(mod.Bytes())
	leafIdx := uint64(0)

	// we only write the single leaf H(k,r) for now
	var buf bytes.Buffer
	commitmentBytes := commitment.Bytes()
	buf.Write(make([]byte, modNumBytes-len(commitmentBytes)))
	buf.Write(commitmentBytes)
	for i := 0; i < NUM_LEAVES-1; i++ {
		buf.Write(make([]byte, modNumBytes))
	}

	merkleRoot, proofPath, _, err := merkletree.BuildReaderProof(&buf, mimc.NewMiMC(), modNumBytes, leafIdx)
	if err != nil {
		panic(err)
	}

	merklePathElements := make([]frontend.Variable, len(proofPath))
	for i, pathElement := range proofPath {
		merklePathElements[i] = pathElement
	}

	var merkleRootBigInt big.Int
	merkleRootBigInt.SetBytes(merkleRoot)

	circuit := &Circuit{
		Nullifier:     nullifier,
		Secret:        secret,
		NullifierHash: nullifierHash,
		Commitment:    commitment,
		nullifierHash: nullifierHash,
		commitment:    commitment,

		MerkleRootHash:     merkleRoot,
		MerklePathElements: merklePathElements,
		LeafIdx:            leafIdx,
		merkleRootHash:     &merkleRootBigInt,
	}
	return circuit
}

func (c *Circuit) GetPublicOutputs() (*big.Int, *big.Int, *big.Int) {
	return c.nullifierHash, c.commitment, c.merkleRootHash
}

// Define declares the circuit's constraints
func (c *Circuit) Define(api frontend.API) error {
	// 1. we first verify nullifier and commitment
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

	// then we verify the merkle proof
	merkleProof := merkle.MerkleProof{
		RootHash: c.MerkleRootHash,
		Path:     c.MerklePathElements,
	}
	mimc.Reset()
	merkleProof.VerifyProof(api, &mimc, c.LeafIdx)

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
