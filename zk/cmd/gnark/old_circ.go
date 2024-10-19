package main

import (
	"fmt"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
)

const COUNT = 4

/*
Public tag:
`gnark:",public"`
*/

///////////////////

type Circuit struct {
	X frontend.Variable
	Y frontend.Variable `gnark:",public"`
}

func (circuit *Circuit) Define(api frontend.API) error {
	x3 := api.Mul(circuit.X, circuit.X, circuit.X)
	x3Plus := api.Add(x3, circuit.X, 5)

	api.AssertIsEqual(x3Plus, circuit.Y)

	return nil

}

// f(x) = x^3 + x + 5 = y
func cirMain() {

	var poly Circuit

	r1cs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &poly)

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	pk, vk, _ := groth16.Setup(r1cs)

	_ = vk
	fmt.Printf("pk.NbG2(): %v\n", pk.NbG2())
	fmt.Printf("pk.NbG1(): %v\n", pk.NbG1())

	assignment := &Circuit{
		X: 3,
		Y: 35,
	}

	witness, _ := frontend.NewWitness(assignment, ecc.BN254.ScalarField())

	proof, err := groth16.Prove(r1cs, pk, witness)

	_ = proof
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	pubWitness, _ := witness.Public()

	result := groth16.Verify(proof, vk, pubWitness)

	if result == nil {
		fmt.Printf("\"sonuc\": %v\n", "basarili")
	}

	//proof.(*bn254.Proof).Ar
}
