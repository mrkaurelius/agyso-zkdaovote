package zk

import (
	"crypto/rand"
	"errors"
	"math/big"
	"strings"

	"github.com/consensys/gnark-crypto/ecc"
	tedd "github.com/consensys/gnark-crypto/ecc/twistededwards"
	"github.com/consensys/gnark/backend/plonk"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/algebra/native/twistededwards"
)

type CircuitMain struct {
	VoteWeight   frontend.Variable    `gnark:",public"`
	MasterPubKey twistededwards.Point `gnark:",public"`
	Vote         [4]frontend.Variable
	Randoms      [4]frontend.Variable
	EncVoteOld   VotesCircuit `gnark:",public"`
	EncVoteNew   VotesCircuit `gnark:",public"`
}

const COUNT = 4

type ElGamalCircuit struct {
	Left  twistededwards.Point
	Right twistededwards.Point
}

type VotesCircuit struct {
	ElGamals [4]ElGamalCircuit
}

func CreateElGamalCircuit(curve twistededwards.Curve, value, random frontend.Variable, base, publicKey twistededwards.Point) ElGamalCircuit {
	return ElGamalCircuit{curve.ScalarMul(base, random), curve.Add(curve.ScalarMul(base, value), curve.ScalarMul(publicKey, random))}
}

func CreateVotesCircuit(curve twistededwards.Curve, value, randoms []frontend.Variable, base, publicKey twistededwards.Point) VotesCircuit {

	votes := VotesCircuit{}
	for i := 0; i < COUNT; i++ {
		votes.ElGamals[i] = CreateElGamalCircuit(curve, value[i], randoms[i], base, publicKey)
	}

	return votes
}

func AddElGamalCircuit(curve twistededwards.Curve, oldEncVote, addEncVote ElGamalCircuit) ElGamalCircuit {
	return ElGamalCircuit{curve.Add(oldEncVote.Left, addEncVote.Left), curve.Add(oldEncVote.Right, addEncVote.Right)}
}

func AddVotesCircuit(curve twistededwards.Curve, oldEncVote, addEncVote VotesCircuit) VotesCircuit {
	newVotes := VotesCircuit{}
	for i := 0; i < COUNT; i++ {
		newVotes.ElGamals[i] = AddElGamalCircuit(curve, oldEncVote.ElGamals[i], addEncVote.ElGamals[i])
	}

	return newVotes
}

func CheckVoteRangeCircuit(api frontend.API, weight frontend.Variable, vote []frontend.Variable) {
	var sum frontend.Variable = 0
	for i := 0; i < len(vote); i++ {
		api.AssertIsLessOrEqual(0, vote[i])
		sum = api.Add(sum, vote[i])
	}
	api.AssertIsEqual(sum, weight)
}

func CheckElGamalEqualityCircuit(api frontend.API, enc1, enc2 VotesCircuit) {
	for i := 0; i < COUNT; i++ {
		api.AssertIsEqual(enc1.ElGamals[i].Left.X, enc2.ElGamals[i].Left.X)
		api.AssertIsEqual(enc1.ElGamals[i].Left.Y, enc2.ElGamals[i].Left.Y)
		api.AssertIsEqual(enc1.ElGamals[i].Right.X, enc2.ElGamals[i].Right.X)
		api.AssertIsEqual(enc1.ElGamals[i].Right.Y, enc2.ElGamals[i].Right.Y)
	}
}

func (circuit *CircuitMain) Define(api frontend.API) error {
	curve, err := twistededwards.NewEdCurve(api, tedd.BN254)
	if err != nil {
		return err
	}

	base := twistededwards.Point{}
	base.X = Base.X
	base.Y = Base.Y

	insideEncAddVote := CreateVotesCircuit(curve, circuit.Vote[:], circuit.Randoms[:], base, circuit.MasterPubKey)
	insideEncVoteNew := AddVotesCircuit(curve, circuit.EncVoteOld, insideEncAddVote)

	CheckElGamalEqualityCircuit(api, insideEncVoteNew, circuit.EncVoteNew)
	CheckVoteRangeCircuit(api, circuit.VoteWeight, circuit.Vote[:])

	_ = insideEncVoteNew
	_ = curve
	return nil
}

func GenerateProof(power, vote0, vote1, vote2, vote3 int, pubKeyX, pubKeyY, encBCVote string) error {

	bcEncVotes := new(Votes)
	if strings.Compare(encBCVote, "0") == 0 {
		bcEncVotes = zeroToVote()
	} else {
		bcEncVotes = StringToVotesSolidity(encBCVote)
	}

	pub := StringToPointUncompress(pubKeyX, pubKeyY)

	votes := []*big.Int{
		new(big.Int).SetInt64(int64(vote0)),
		new(big.Int).SetInt64(int64(vote1)),
		new(big.Int).SetInt64(int64(vote2)),
		new(big.Int).SetInt64(int64(vote3))}

	rand0, _ := rand.Int(rand.Reader, ecc.BN254.ScalarField())
	rand1, _ := rand.Int(rand.Reader, ecc.BN254.ScalarField())
	rand2, _ := rand.Int(rand.Reader, ecc.BN254.ScalarField())
	rand3, _ := rand.Int(rand.Reader, ecc.BN254.ScalarField())
	randoms := []*big.Int{rand0, rand1, rand2, rand3}

	addEncVotes := CreateVotes(votes, randoms, pub)
	newEncVotes := AddVotes(bcEncVotes, addEncVotes)

	for i := 0; i < COUNT; i++ {
		println(bcEncVotes.ElGamals[i].Left.X.String())
		println(bcEncVotes.ElGamals[i].Left.Y.String())
		println(bcEncVotes.ElGamals[i].Right.X.String())
		println(bcEncVotes.ElGamals[i].Right.Y.String())
	}

	ccs := GetCCSPlonk()
	pk := GetPKPlonk()
	vk := GetVKPlonk()

	assignment := CircuitMain{}

	assignment.VoteWeight = power
	assignment.MasterPubKey.X = pub.X
	assignment.MasterPubKey.Y = pub.Y

	for i := 0; i < COUNT; i++ {
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

	witness, err := frontend.NewWitness(&assignment, ecc.BN254.ScalarField())
	if err != nil {
		return err
	}
	publicWitness, err := witness.Public()
	if err != nil {
		return err
	}
	proof, err := plonk.Prove(ccs, pk, witness)
	if err != nil {
		return err
	}

	result := plonk.Verify(proof, vk, publicWitness)
	if result == nil {
		return nil
	} else {
		return errors.New("proof verification error")
	}
}
