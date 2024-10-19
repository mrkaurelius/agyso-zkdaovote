package zk

import (
	"bytes"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/constraint"
	"os"
)

func GetCCS() constraint.ConstraintSystem {
	byteArray, _ := os.ReadFile("../data/ccs.bin")
	ccs := groth16.NewCS(ecc.BN254)
	ccs.ReadFrom(bytes.NewReader(byteArray))

	return ccs
}

func GetPK() groth16.ProvingKey {
	byteArray, _ := os.ReadFile("../data/pk.bin")
	pk := groth16.NewProvingKey(ecc.BN254)
	pk.ReadFrom(bytes.NewReader(byteArray))

	return pk
}