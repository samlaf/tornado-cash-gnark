package main

import (
	"math/big"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"

	tccircuit "github.com/samlaf/tornado-cash-gnark/circuit"
)

func main() {

	circuit := tccircuit.InitializedCircuit()
	circuitR1cs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &circuit)
	if err != nil {
		panic(err)
	}

	nullifier := big.NewInt(1)
	secret := big.NewInt(2)
	assignment := tccircuit.NewAssignment(nullifier, secret)

	witness, err := frontend.NewWitness(assignment, ecc.BN254.ScalarField())
	if err != nil {
		panic(err)
	}
	publicWitness, err := witness.Public()
	if err != nil {
		panic(err)
	}

	pk, vk, err := groth16.Setup(circuitR1cs)
	if err != nil {
		panic(err)
	}

	proof, err := groth16.Prove(circuitR1cs, pk, witness)
	if err != nil {
		panic(err)
	}
	err = groth16.Verify(proof, vk, publicWitness)
	if err != nil {
		panic("invalid proof")
	}
}
