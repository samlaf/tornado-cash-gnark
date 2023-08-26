package main

import (
	"fmt"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	
	"github.com/samlaf/tornado-cash-gnark/circuit"

	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
)

func main() {
	var mimcCircuit circuit.Circuit
	r1cs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &mimcCircuit)
	if err != nil {
		panic(err)
	}
	// witness
	preImage := uint64(35)
	hash := calculateMimcOutput(preImage)
	assignment := &circuit.Circuit{
		Hash:     hash,
		PreImage: preImage,
	}

	witness, _ := frontend.NewWitness(assignment, ecc.BN254.ScalarField())
	publicWitness, _ := witness.Public()
	fmt.Println("witness:", witness.Vector())
	pk, vk, err := groth16.Setup(r1cs)
	if err != nil {
		panic(err)
	}

	proof, err := groth16.Prove(r1cs, pk, witness)
	if err != nil {
		panic(err)
	}
	err = groth16.Verify(proof, vk, publicWitness)
	if err != nil {
		panic("invalid proof")
	}
}

func calculateMimcOutput(inputNum uint64) string {
	hash := mimc.NewMiMC()

	// hash the input (written as big endian byte array)
	input := fr.Element{}
	input.SetUint64(inputNum)
	inputBytes32 := input.Bytes()
	hash.Write(inputBytes32[:])

	// get the output (which is a 32 byte slice)
	outputBytes := hash.Sum(nil)

	// write as a field element to get a nice string representation to output
	element := fr.Element{}
	element.SetBytes(outputBytes)
	return element.String()
}
