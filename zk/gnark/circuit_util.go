package zk

import (
	tedd "github.com/consensys/gnark-crypto/ecc/twistededwards"
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

	// CheckElGamalEqualityCircuit(api, insideEncVoteNew, circuit.EncVoteNew)
	CheckVoteRangeCircuit(api, circuit.VoteWeight, circuit.Vote[:])

	_ = insideEncVoteNew
	_ = curve
	return nil
}
