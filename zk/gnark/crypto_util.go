package zk

import (
	"math/big"

	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	bn254 "github.com/consensys/gnark-crypto/ecc/bn254/twistededwards"
)

var Base = bn254.GetEdwardsCurve().Base

type ElGamal struct {
	Left  *bn254.PointAffine
	Right *bn254.PointAffine
}

type Votes struct {
	ElGamals [COUNT]*ElGamal
}

func NewElGamal(p1, p2 *bn254.PointAffine) *ElGamal {
	p := new(ElGamal)
	p.Left = p1
	p.Right = p2

	return p
}

func NewPoint(x, y *fr.Element) *bn254.PointAffine {
	p := new(bn254.PointAffine)
	p.X = *x
	p.Y = *y

	return p
}

func CreateElGamal(message *big.Int, publicKey *bn254.PointAffine, random *big.Int) *ElGamal {
	left := new(bn254.PointAffine).ScalarMultiplication(&Base, random)
	right := new(bn254.PointAffine).Add(
		new(bn254.PointAffine).ScalarMultiplication(&Base, message),
		new(bn254.PointAffine).ScalarMultiplication(publicKey, random))

	return NewElGamal(left, right)
}

func CreateVotes(message, random []*big.Int, publicKey *bn254.PointAffine) *Votes {
	votes := new(Votes)
	for i := 0; i < COUNT; i++ {
		votes.ElGamals[i] = CreateElGamal(message[i], publicKey, random[i])
	}

	return votes
}

func AddVotes(oldVotes, addVotes *Votes) *Votes {
	newVotes := new(Votes)
	for i := 0; i < COUNT; i++ {
		newVotes.ElGamals[i] = new(ElGamal)
		newVotes.ElGamals[i].Left = new(bn254.PointAffine).Add(oldVotes.ElGamals[i].Left, addVotes.ElGamals[i].Left)
		newVotes.ElGamals[i].Right = new(bn254.PointAffine).Add(oldVotes.ElGamals[i].Right, addVotes.ElGamals[i].Right)
	}

	return newVotes
}
