package zk

import (
	"bytes"
	"encoding/hex"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	bn254 "github.com/consensys/gnark-crypto/ecc/bn254/twistededwards"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/backend/plonk"
	"github.com/consensys/gnark/constraint"
	"os"
)

func StringToPointUncompress(strX, strY string) *bn254.PointAffine {

	x, _ := new(fr.Element).SetString(strX)
	y, _ := new(fr.Element).SetString(strY)
	pub := new(bn254.PointAffine)
	pub.X = *x
	pub.Y = *y

	return pub
}

func StringsToElGamalUncompress(str1X, str1Y, str2X, str2Y string) *ElGamal {
	left := StringToPointUncompress(str1X, str1Y)
	right := StringToPointUncompress(str2X, str2Y)

	cipher := new(ElGamal)
	cipher.Left = left
	cipher.Right = right

	return cipher
}

func StringsToVotesUncompress(str1X, str1Y, str2X, str2Y, str3X, str3Y, str4X, str4Y, str5X, str5Y, str6X, str6Y, str7X, str7Y, str8X, str8Y string) *Votes {
	votes := new(Votes)
	votes.ElGamals[0] = StringsToElGamalUncompress(str1X, str1Y, str2X, str2Y)
	votes.ElGamals[1] = StringsToElGamalUncompress(str3X, str3Y, str4X, str4Y)
	votes.ElGamals[2] = StringsToElGamalUncompress(str5X, str5Y, str6X, str6Y)
	votes.ElGamals[3] = StringsToElGamalUncompress(str7X, str7Y, str8X, str8Y)

	return votes
}

func StringToPoint(str string) *bn254.PointAffine {
	pubArray, _ := hex.DecodeString(str)
	pub := new(bn254.PointAffine)
	pub.Unmarshal(pubArray)

	return pub
}

func StringsToElGamal(str1, str2 string) *ElGamal {
	left := StringToPoint(str1)
	right := StringToPoint(str2)

	cipher := new(ElGamal)
	cipher.Left = left
	cipher.Right = right

	return cipher
}

func StringsToVotes(str1, str2, str3, str4, str5, str6, str7, str8 string) *Votes {
	votes := new(Votes)
	votes.ElGamals[0] = StringsToElGamal(str1, str2)
	votes.ElGamals[1] = StringsToElGamal(str3, str4)
	votes.ElGamals[2] = StringsToElGamal(str5, str6)
	votes.ElGamals[3] = StringsToElGamal(str7, str8)

	return votes
}

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

func GetCCSPlonk() constraint.ConstraintSystem {
	byteArray, _ := os.ReadFile("../data/ccs.bin")
	ccs := plonk.NewCS(ecc.BN254)
	ccs.ReadFrom(bytes.NewReader(byteArray))
	//r1cs := ccs.(*cs.SparseR1CS)
	//srs, srsLagrangeInterpolation, _ := unsafekzg.NewSRS(r1cs)

	return ccs
}

func GetPKPlonk() plonk.ProvingKey {
	byteArray, _ := os.ReadFile("../data/pk.bin")
	pk := plonk.NewProvingKey(ecc.BN254)
	pk.ReadFrom(bytes.NewReader(byteArray))

	return pk
}
