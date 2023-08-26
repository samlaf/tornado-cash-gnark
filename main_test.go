package main

import (
	"fmt"
	"testing"

	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
)

func TestMimc(t *testing.T) {
	hash := mimc.NewMiMC()

	// hash the input 35 (written as big endian byte array)
	input := [32]byte{0}
	input[31] = 35
	hash.Write(input[:])

	// get the output (which is a 32 byte slice)
	outputBytes := hash.Sum(nil)
	
	// write as a field element to get a nice string representation to output
	element := fr.Element{}
	element.SetBytes(outputBytes)
	fmt.Println(element.String())
}
