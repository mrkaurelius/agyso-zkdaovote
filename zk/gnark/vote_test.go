package zk

import (
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	bn254 "github.com/consensys/gnark-crypto/ecc/bn254/twistededwards"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
)

func TestCircuit(t *testing.T) {
	//create pair
	priv, _ := rand.Int(rand.Reader, ecc.BN254.ScalarField())
	pub := new(bn254.PointAffine).ScalarMultiplication(&Base, priv)

	// weight
	weight := new(big.Int).SetInt64(10)

	// create current bc Votes
	votes := []*big.Int{new(big.Int).SetInt64(10), new(big.Int).SetInt64(20), new(big.Int).SetInt64(30), new(big.Int).SetInt64(40)}
	randoms := []*big.Int{new(big.Int).SetInt64(1111), new(big.Int).SetInt64(2222), new(big.Int).SetInt64(3333), new(big.Int).SetInt64(4444)}
	currentEncVotes := CreateVotes(votes, randoms, pub)

	// create add Votes
	addVotes := []*big.Int{new(big.Int).SetInt64(1), new(big.Int).SetInt64(2), new(big.Int).SetInt64(3), new(big.Int).SetInt64(4)}
	addRandoms := []*big.Int{new(big.Int).SetInt64(111), new(big.Int).SetInt64(222), new(big.Int).SetInt64(333), new(big.Int).SetInt64(444)}
	addEncVotes := CreateVotes(addVotes, addRandoms, pub)

	// create new stage of Votes
	newEncVotes := AddVotes(currentEncVotes, addEncVotes)
	//

	var circuit CircuitMain
	ccs, _ := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &circuit)

	pk, vk, _ := groth16.Setup(ccs)

	_ = vk
	assignment := CircuitMain{}

	assignment.Weight = weight
	assignment.MasterPubKey.X = pub.X
	assignment.MasterPubKey.Y = pub.Y
	for i := 0; i < COUNT; i++ {
		assignment.Randoms[i] = randoms[i]
		assignment.Vote[i] = votes[i]

		assignment.EncVoteNew.ElGamals[i].Left.X = newEncVotes.ElGamals[i].Left.X
		assignment.EncVoteNew.ElGamals[i].Left.Y = newEncVotes.ElGamals[i].Left.Y
		assignment.EncVoteNew.ElGamals[i].Right.X = newEncVotes.ElGamals[i].Right.X
		assignment.EncVoteNew.ElGamals[i].Right.Y = newEncVotes.ElGamals[i].Right.Y

		assignment.EncVoteOld.ElGamals[i].Left.X = currentEncVotes.ElGamals[i].Left.X
		assignment.EncVoteOld.ElGamals[i].Left.Y = currentEncVotes.ElGamals[i].Left.Y
		assignment.EncVoteOld.ElGamals[i].Right.X = currentEncVotes.ElGamals[i].Right.X
		assignment.EncVoteOld.ElGamals[i].Right.Y = currentEncVotes.ElGamals[i].Right.Y
	}

	witness, _ := frontend.NewWitness(&assignment, ecc.BN254.ScalarField())
	publicWitness, _ := witness.Public()

	// create proof
	proof, error := groth16.Prove(ccs, pk, witness)

	print(error)
	_ = publicWitness
	_ = proof
}
