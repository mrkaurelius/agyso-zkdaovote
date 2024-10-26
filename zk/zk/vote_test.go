package zk

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	bn254 "github.com/consensys/gnark-crypto/ecc/bn254/twistededwards"
	"github.com/consensys/gnark/backend/plonk"
	"github.com/consensys/gnark/frontend"
	_ "github.com/consensys/gnark/frontend/cs/r1cs"
)

func TestPlonkCircuit(t *testing.T) {
	//create pair
	priv := new(big.Int).SetInt64(100)
	pub := new(bn254.PointAffine).ScalarMultiplication(&Base, priv)
	// priv := new(big.Int).SetInt64(100)

	pubArray := pub.Bytes()
	pubStr := hex.EncodeToString(pubArray[:])

	println("pubStr: " + pubStr)
	// weight
	weight := new(big.Int).SetInt64(10)

	fmt.Printf("pub.X.String(): %v\n", pub.X.String())
	fmt.Printf("pub.Y.String(): %v\n", pub.Y.String())

	// create current bc Votes
	votes := []*big.Int{new(big.Int).SetInt64(1), new(big.Int).SetInt64(1), new(big.Int).SetInt64(1), new(big.Int).SetInt64(1)}
	randoms := []*big.Int{new(big.Int).SetInt64(1111), new(big.Int).SetInt64(1111), new(big.Int).SetInt64(1111), new(big.Int).SetInt64(1111)}
	currentEncVotes := CreateVotes(votes, randoms, pub)

	currentEncVotes = StringToVotesSolidity("248f64260528cc106604cca088a01ac1aa92783dd2bf131f3ab28025a95e5ea91d225d2c97bdb6b7701d78d8ef08e05acecda1e7e2ff88d6bdf69b2cfe7f2fa80f94cbffa2a1b4571ea97160b0dc21c70f3beaab10265e3655ae6dd795f3f7df0aecb8871b679eb5e1500874931d1ac13ac652c684f1ea94b80ad3c8a0adc1be248f64260528cc106604cca088a01ac1aa92783dd2bf131f3ab28025a95e5ea91d225d2c97bdb6b7701d78d8ef08e05acecda1e7e2ff88d6bdf69b2cfe7f2fa80f94cbffa2a1b4571ea97160b0dc21c70f3beaab10265e3655ae6dd795f3f7df0aecb8871b679eb5e1500874931d1ac13ac652c684f1ea94b80ad3c8a0adc1be248f64260528cc106604cca088a01ac1aa92783dd2bf131f3ab28025a95e5ea91d225d2c97bdb6b7701d78d8ef08e05acecda1e7e2ff88d6bdf69b2cfe7f2fa80f94cbffa2a1b4571ea97160b0dc21c70f3beaab10265e3655ae6dd795f3f7df0aecb8871b679eb5e1500874931d1ac13ac652c684f1ea94b80ad3c8a0adc1be248f64260528cc106604cca088a01ac1aa92783dd2bf131f3ab28025a95e5ea91d225d2c97bdb6b7701d78d8ef08e05acecda1e7e2ff88d6bdf69b2cfe7f2fa80f94cbffa2a1b4571ea97160b0dc21c70f3beaab10265e3655ae6dd795f3f7df0aecb8871b679eb5e1500874931d1ac13ac652c684f1ea94b80ad3c8a0adc1be")
	// println("pub.X.String(): \n", pub.X.String())
	// println("pub.Y.String(): \n", pub.Y.String())
	// println("currentEncVotes.ElGamals[0].Left.X.String(): \n", currentEncVotes.ElGamals[0].Left.X.String())
	// println("currentEncVotes.ElGamals[0].Left.Y.String(): \n", currentEncVotes.ElGamals[0].Left.Y.String())
	// println("currentEncVotes.ElGamals[0].Right.X.String(): \n", currentEncVotes.ElGamals[0].Right.X.String())
	// println("currentEncVotes.ElGamals[0].Right.Y.String(): \n", currentEncVotes.ElGamals[0].Right.Y.String())

	for i := 0; i < COUNT; i++ {
		println(currentEncVotes.ElGamals[i].Left.X.String())
		println(currentEncVotes.ElGamals[i].Left.Y.String())
		println(currentEncVotes.ElGamals[i].Right.X.String())
		println(currentEncVotes.ElGamals[i].Right.Y.String())
	}

	// create add Votes
	addVotes := []*big.Int{new(big.Int).SetInt64(1), new(big.Int).SetInt64(2), new(big.Int).SetInt64(3), new(big.Int).SetInt64(4)}
	addRandoms := []*big.Int{new(big.Int).SetInt64(111), new(big.Int).SetInt64(222), new(big.Int).SetInt64(333), new(big.Int).SetInt64(444)}
	addEncVotes := CreateVotes(addVotes, addRandoms, pub)

	// create new stage of Votes
	newEncVotes := AddVotes(currentEncVotes, addEncVotes)

	//
	arrayLX := currentEncVotes.ElGamals[0].Left.X.Bytes()
	fmt.Printf("newEncVotes.ElGamals[0].Left.X.String(): %v\n", hex.EncodeToString(arrayLX[:]))
	arrayLY := currentEncVotes.ElGamals[0].Left.Y.Bytes()
	fmt.Printf("newEncVotes.ElGamals[0].Left.Y.String(): %v\n", hex.EncodeToString(arrayLY[:]))
	arrayRX := currentEncVotes.ElGamals[0].Right.X.Bytes()
	fmt.Printf("newEncVotes.ElGamals[0].Right.X.String(): %v\n", hex.EncodeToString(arrayRX[:]))
	arrayRY := currentEncVotes.ElGamals[0].Right.Y.Bytes()
	fmt.Printf("newEncVotes.ElGamals[0].Right.Y.String(): %v\n", hex.EncodeToString(arrayRY[:]))

	ccs := GetCCSPlonk()
	pk := GetPKPlonk()
	vk := GetVKPlonk()

	assignment := CircuitMain{}

	assignment.VoteWeight = weight
	assignment.MasterPubKey.X = pub.X
	assignment.MasterPubKey.Y = pub.Y

	for i := 0; i < COUNT; i++ {
		assignment.Randoms[i] = addRandoms[i]
		assignment.Vote[i] = addVotes[i]

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
	publicWitness, err := witness.Public()

	var buffxxx bytes.Buffer
	publicWitness.WriteTo(&buffxxx)
	fmt.Printf("hex.EncodeToString(buffxxx.Bytes()): %v\n", hex.EncodeToString(buffxxx.Bytes()))

	fmt.Printf("ahk")

	xxx, _ := publicWitness.MarshalBinary()
	xxxStr := hex.EncodeToString(xxx)
	votesPrime := StringToVotesSolidity(xxxStr[1240 : 1240+16*64])
	_ = votesPrime
	if err != nil {
		t.Fatalf("public witnes err %s", err)
	}

	// create proof
	proof, err := plonk.Prove(ccs, pk, witness)
	if err != nil {
		t.Fatalf("prove error: %s", err)
	}
	//t.Logf("proof: %+v", proof)

	err = plonk.Verify(proof, vk, publicWitness)
	if err != nil {
		t.Fatalf("verify error: %s", err)
	}

	/*
		proofBasePath := "/var/tmp/agyso-daovote/proof/plonk"

		proofPath := fmt.Sprintf("%s/plonk.proof", proofBasePath)
		proofFile, err := os.Create(proofPath)
		if err != nil {
			t.Fatal(err)
		}

		vkFilePath := fmt.Sprintf("%s/plonk.vk", proofBasePath)
		vkFile, err := os.Create(vkFilePath)
		if err != nil {
			t.Fatal(err)
		}

		witnessFilePath := fmt.Sprintf("%s/plonk_pub_input.pub", proofBasePath)
		witnessFile, err := os.Create(witnessFilePath)
		if err != nil {
			t.Fatal(err)
		}

		defer proofFile.Close()
		defer vkFile.Close()
		defer witnessFile.Close()

		_, err = proof.WriteTo(proofFile)
		if err != nil {
			t.Fatal("could not serialize proof into file")
		}
		_, err = vk.WriteTo(vkFile)
		if err != nil {
			t.Fatal("could not serialize verification key into file")
		}
		_, err = publicWitness.WriteTo(witnessFile)
		if err != nil {
			t.Fatal("could not serialize proof into file")
		}

		t.Logf("proof written into %s\n", proofPath)
		t.Logf("verification key into %s\n", vkFilePath)
		t.Logf("public witness written into %s\n", witnessFilePath)

		t.Log("Proof verification succeeded")
	*/
}

func TestStringToVotesSolidity(t *testing.T) {
	priv := new(big.Int).SetInt64(100)
	pub := new(bn254.PointAffine).ScalarMultiplication(&Base, priv)
	print(pub.X.String())
}

func TestZero(t *testing.T) {

	priv := new(big.Int).SetInt64(100)
	pub := new(bn254.PointAffine).ScalarMultiplication(&Base, priv)

	pubXArray := pub.X.Bytes()
	pubYArray := pub.Y.Bytes()

	err := GenerateProof(10, 1, 2, 3, 4, hex.EncodeToString(pubXArray[:]), hex.EncodeToString(pubYArray[:]), "248f64260528cc106604cca088a01ac1aa92783dd2bf131f3ab28025a95e5ea91d225d2c97bdb6b7701d78d8ef08e05acecda1e7e2ff88d6bdf69b2cfe7f2fa80f94cbffa2a1b4571ea97160b0dc21c70f3beaab10265e3655ae6dd795f3f7df0aecb8871b679eb5e1500874931d1ac13ac652c684f1ea94b80ad3c8a0adc1be248f64260528cc106604cca088a01ac1aa92783dd2bf131f3ab28025a95e5ea91d225d2c97bdb6b7701d78d8ef08e05acecda1e7e2ff88d6bdf69b2cfe7f2fa80f94cbffa2a1b4571ea97160b0dc21c70f3beaab10265e3655ae6dd795f3f7df0aecb8871b679eb5e1500874931d1ac13ac652c684f1ea94b80ad3c8a0adc1be248f64260528cc106604cca088a01ac1aa92783dd2bf131f3ab28025a95e5ea91d225d2c97bdb6b7701d78d8ef08e05acecda1e7e2ff88d6bdf69b2cfe7f2fa80f94cbffa2a1b4571ea97160b0dc21c70f3beaab10265e3655ae6dd795f3f7df0aecb8871b679eb5e1500874931d1ac13ac652c684f1ea94b80ad3c8a0adc1be248f64260528cc106604cca088a01ac1aa92783dd2bf131f3ab28025a95e5ea91d225d2c97bdb6b7701d78d8ef08e05acecda1e7e2ff88d6bdf69b2cfe7f2fa80f94cbffa2a1b4571ea97160b0dc21c70f3beaab10265e3655ae6dd795f3f7df0aecb8871b679eb5e1500874931d1ac13ac652c684f1ea94b80ad3c8a0adc1be")

	fmt.Printf("err: %v\n", err)
}

func TestStringToDec(t *testing.T) {
	str := "248f64260528cc106604cca088a01ac1aa92783dd2bf131f3ab28025a95e5ea91d225d2c97bdb6b7701d78d8ef08e05acecda1e7e2ff88d6bdf69b2cfe7f2fa80f94cbffa2a1b4571ea97160b0dc21c70f3beaab10265e3655ae6dd795f3f7df0aecb8871b679eb5e1500874931d1ac13ac652c684f1ea94b80ad3c8a0adc1be248f64260528cc106604cca088a01ac1aa92783dd2bf131f3ab28025a95e5ea91d225d2c97bdb6b7701d78d8ef08e05acecda1e7e2ff88d6bdf69b2cfe7f2fa80f94cbffa2a1b4571ea97160b0dc21c70f3beaab10265e3655ae6dd795f3f7df0aecb8871b679eb5e1500874931d1ac13ac652c684f1ea94b80ad3c8a0adc1be248f64260528cc106604cca088a01ac1aa92783dd2bf131f3ab28025a95e5ea91d225d2c97bdb6b7701d78d8ef08e05acecda1e7e2ff88d6bdf69b2cfe7f2fa80f94cbffa2a1b4571ea97160b0dc21c70f3beaab10265e3655ae6dd795f3f7df0aecb8871b679eb5e1500874931d1ac13ac652c684f1ea94b80ad3c8a0adc1be248f64260528cc106604cca088a01ac1aa92783dd2bf131f3ab28025a95e5ea91d225d2c97bdb6b7701d78d8ef08e05acecda1e7e2ff88d6bdf69b2cfe7f2fa80f94cbffa2a1b4571ea97160b0dc21c70f3beaab10265e3655ae6dd795f3f7df0aecb8871b679eb5e1500874931d1ac13ac652c684f1ea94b80ad3c8a0adc1be"
	priv := new(big.Int).SetInt64(100)

	DecryptEncryptedBulletsFromStr(str, priv)
}
