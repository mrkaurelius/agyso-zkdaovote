package zk

import (
	"encoding/json"
	"io"
	"math/big"
	"os"
	"path/filepath"

	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	bn254 "github.com/consensys/gnark-crypto/ecc/bn254/twistededwards"
)

var Base = bn254.GetEdwardsCurve().Base
var electionKey KeyData

type ElGamal struct {
	Left  *bn254.PointAffine
	Right *bn254.PointAffine
}

type Votes struct {
	ElGamals [COUNT]*ElGamal
}

type KeyData struct {
	PrivateKey string `json:"privateKey"`
	PublicKeyX string `json:"publicKeyX"`
	PublicKeyY string `json:"publicKeyY"`
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

func GetElectionKeys() (key *KeyData, err error) {

	keyFile, err := os.Open(filepath.Join(circuitSetupPath, "keyfile.json"))
	if err != nil {
		return nil, err
	}
	defer keyFile.Close()

	keyFileBytes, err := io.ReadAll(keyFile)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(keyFileBytes, &key); err != nil {
		return nil, err
	}

	return key, nil
}

func DecryptEncryprtedBallotBox(encryped string, secretStr string) []int {
	secret, _ := new(big.Int).SetString(secretStr, 10)
	return DecryptEncryptedBallotsFromStr(encryped, secret)
}

func DecryptEncryptedBallotsFromStr(str string, sec *big.Int) []int {
	votes := StringToVotesSolidity(str)

	return DecryptEncryptedBallotsFrom(votes, sec)

}

func DecryptEncryptedBallotsFrom(ballots *Votes, sec *big.Int) []int {
	decVote := make([]int, 4)
	for i := 0; i < 4; i++ {
		decVote[i] = DecryptElgamalBrute(ballots.ElGamals[i], sec)
	}

	return decVote
}

// (c1,c2) = (g^r, g^m*pk^r)
func DecryptElgamalBrute(enc *ElGamal, sec *big.Int) int {

	dec := new(bn254.PointAffine).Add(
		enc.Right,
		new(bn254.PointAffine).ScalarMultiplication(enc.Left, new(big.Int).Neg(sec)))

	//TODO
	for i := 0; i < 1000; i++ {
		if new(bn254.PointAffine).ScalarMultiplication(&Base, big.NewInt(int64(i))).X.Equal(&dec.X) {
			return int(i)
		}
	}

	return 1111
}
