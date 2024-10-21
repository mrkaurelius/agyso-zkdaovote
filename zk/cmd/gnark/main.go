package main

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/plonk"
	"github.com/consensys/gnark/frontend"
	_ "github.com/consensys/gnark/frontend/cs/r1cs"
	zk "github.com/mrkaurelius/ppp-daovote/zk/gnark"
	"math/big"
)

func main() {
	votePowerPtr := flag.Int("power", 0, "vote power")
	pubKeyPtr := flag.String("pubKey", "", "public key")
	encBCVote1Left := flag.String("encBCVote1Left", "", "encBCVote1Left")
	encBCVote1Right := flag.String("encBCVote1Right", "", "encBCVote1Right")
	encBCVote2Left := flag.String("encBCVote2Left", "", "encBCVote2Left")
	encBCVote2Right := flag.String("encBCVote2Right", "", "encBCVote2Right")
	encBCVote3Left := flag.String("encBCVote3Left", "", "encBCVote3Left")
	encBCVote3Right := flag.String("encBCVote3Right", "", "encBCVote3Right")
	encBCVote4Left := flag.String("encBCVote4Left", "", "encBCVote4Left")
	encBCVote4Right := flag.String("encBCVote4Right", "", "encBCVote4Right")
	vote0 := flag.Int("vote0", 0, "vote1")
	vote1 := flag.Int("vote1", 0, "vote2")
	vote2 := flag.Int("vote2", 0, "vote3")
	vote3 := flag.Int("vote3", 0, "vote3")

	flag.Parse()
	pub := zk.StringToPoint(*pubKeyPtr)

	votes := []*big.Int{
		new(big.Int).SetInt64(int64(*vote0)),
		new(big.Int).SetInt64(int64(*vote1)),
		new(big.Int).SetInt64(int64(*vote2)),
		new(big.Int).SetInt64(int64(*vote3))}

	//TODO
	rand0, _ := rand.Int(rand.Reader, ecc.BN254.ScalarField())
	rand1, _ := rand.Int(rand.Reader, ecc.BN254.ScalarField())
	rand2, _ := rand.Int(rand.Reader, ecc.BN254.ScalarField())
	rand3, _ := rand.Int(rand.Reader, ecc.BN254.ScalarField())
	randoms := []*big.Int{rand0, rand1, rand2, rand3}

	addEncVotes := zk.CreateVotes(votes, randoms, pub)

	bcEncVotes := zk.StringsToVotes(
		*encBCVote1Left, *encBCVote1Right,
		*encBCVote2Left, *encBCVote2Right,
		*encBCVote3Left, *encBCVote3Right,
		*encBCVote4Left, *encBCVote4Right)

	newEncVotes := zk.AddVotes(bcEncVotes, addEncVotes)

	ccs := zk.GetCCSPlonk()
	pk := zk.GetPKPlonk()

	vkBytes, _ := hex.DecodeString(zk.VKStr)
	vk := plonk.NewVerifyingKey(ecc.BN254)
	vk.ReadFrom(bytes.NewReader(vkBytes))

	assignment := zk.CircuitMain{}

	assignment.VoteWeight = *votePowerPtr
	assignment.MasterPubKey.X = pub.X
	assignment.MasterPubKey.Y = pub.Y

	for i := 0; i < zk.COUNT; i++ {
		assignment.Randoms[i] = randoms[i]
		assignment.Vote[i] = votes[i]

		assignment.EncVoteNew.ElGamals[i].Left.X = newEncVotes.ElGamals[i].Left.X
		assignment.EncVoteNew.ElGamals[i].Left.Y = newEncVotes.ElGamals[i].Left.Y
		assignment.EncVoteNew.ElGamals[i].Right.X = newEncVotes.ElGamals[i].Right.X
		assignment.EncVoteNew.ElGamals[i].Right.Y = newEncVotes.ElGamals[i].Right.Y

		assignment.EncVoteOld.ElGamals[i].Left.X = bcEncVotes.ElGamals[i].Left.X
		assignment.EncVoteOld.ElGamals[i].Left.Y = bcEncVotes.ElGamals[i].Left.Y
		assignment.EncVoteOld.ElGamals[i].Right.X = bcEncVotes.ElGamals[i].Right.X
		assignment.EncVoteOld.ElGamals[i].Right.Y = bcEncVotes.ElGamals[i].Right.Y
	}

	witness, _ := frontend.NewWitness(&assignment, ecc.BN254.ScalarField())
	publicWitness, _ := witness.Public()

	proof, _ := plonk.Prove(ccs, pk, witness)

	result := plonk.Verify(proof, vk, publicWitness)
	if result == nil {
		fmt.Print("OK")
	}

}
