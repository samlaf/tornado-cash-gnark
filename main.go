package main

import (
	"fmt"
	"math/big"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"

	"github.com/samlaf/tornado-cash-gnark/circuit"
)

func main() {
	var mimcCircuit circuit.Circuit
	r1cs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &mimcCircuit)
	if err != nil {
		panic(err)
	}
	// witness
	nullifier := big.NewInt(1)
	secret := big.NewInt(2)
	assignment := circuit.NewCircuit(nullifier, secret)

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
