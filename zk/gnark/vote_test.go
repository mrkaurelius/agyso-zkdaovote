package zk

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"strings"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	bn254 "github.com/consensys/gnark-crypto/ecc/bn254/twistededwards"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/backend/plonk"
	wit "github.com/consensys/gnark/backend/witness"
	cs "github.com/consensys/gnark/constraint/bn254"
	"github.com/consensys/gnark/frontend"
	_ "github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/consensys/gnark/frontend/cs/scs"
	"github.com/consensys/gnark/test/unsafekzg"
)

// // run carefully!!
func TestCCSandPKandVK(t *testing.T) {

	return
	var circuit CircuitMain
	ccs, _ := frontend.Compile(ecc.BN254.ScalarField(), scs.NewBuilder, &circuit)
	r1cs := ccs.(*cs.SparseR1CS)
	srs, srsLagrangeInterpolation, _ := unsafekzg.NewSRS(r1cs)
	pk, vk, _ := plonk.Setup(ccs, srs, srsLagrangeInterpolation)

	var buffCCS bytes.Buffer
	ccs.WriteTo(&buffCCS)
	file0, _ := os.Create("ccs.bin")
	defer file0.Close()
	file0.Write(buffCCS.Bytes())

	var buffPK bytes.Buffer
	pk.WriteTo(&buffPK)
	file1, _ := os.Create("pk.bin")
	defer file1.Close()
	file1.Write(buffPK.Bytes())

	var buffSRS bytes.Buffer
	srs.WriteTo(&buffSRS)
	file2, _ := os.Create("srs.bin")
	defer file2.Close()
	file2.Write(buffSRS.Bytes())

	var buffSRSLag bytes.Buffer
	srsLagrangeInterpolation.WriteTo(&buffSRSLag)
	file3, _ := os.Create("srsLag.bin")
	defer file3.Close()
	file3.Write(buffSRSLag.Bytes())

	var buffVK bytes.Buffer
	vk.WriteTo(&buffVK)
	strVK := strings.ToUpper(hex.EncodeToString(buffVK.Bytes()))
	fmt.Println(strVK)

}

func TestCircuit(t *testing.T) {
	return
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

	ccs := GetCCS()
	pk := GetPK()

	vkBytes, _ := hex.DecodeString(VKStr)
	vk := groth16.NewVerifyingKey(ecc.BN254)
	vk.ReadFrom(bytes.NewReader(vkBytes))

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
	if err != nil {
		t.Fatalf("public witnes err %s", err)
	}

	// create proof
	proof, err := groth16.Prove(ccs, pk, witness)
	if err != nil {
		t.Fatalf("prove error: %s", err)
	}
	t.Logf("proof: %+v", proof)

	err = groth16.Verify(proof, vk, publicWitness)
	if err != nil {
		t.Fatalf("verify error: %s", err)
	}

	proofBasePath := "/var/tmp/agyso-daovote/proof/groth16"

	proofPath := fmt.Sprintf("%s/groth16.proof", proofBasePath)
	proofFile, err := os.Create(proofPath)
	if err != nil {
		t.Fatal(err)
	}

	vkFilePath := fmt.Sprintf("%s/groth16.vk", proofBasePath)
	vkFile, err := os.Create(vkFilePath)
	if err != nil {
		t.Fatal(err)
	}

	witnessFilePath := fmt.Sprintf("%s/groth16_pub_input.pub", proofBasePath)
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

	// >>>>>> Proof serialisation

	witnessByte, _ := os.ReadFile(witnessFilePath)
	witnessPrime, _ := wit.New(ecc.BN254.ScalarField()) // Serialised witness
	witnessPrime.ReadFrom(bytes.NewReader(witnessByte))
	fmt.Printf("witnessPrime: %+v\n", witnessPrime)

	proofByte, _ := os.ReadFile(proofPath)
	proofPrime := groth16.NewProof(ecc.BN254) // Serialised witness
	proofPrime.ReadFrom(bytes.NewReader(proofByte))
	fmt.Printf("proofPrime: %+v\n", proofPrime)

	vkByte, _ := os.ReadFile(vkFilePath)
	vkPrime := groth16.NewVerifyingKey(ecc.BN254) // Serialised witness
	vkPrime.ReadFrom(bytes.NewReader(vkByte))
	fmt.Printf("vkPrime: %+v\n", vkPrime)

	err = groth16.Verify(proofPrime, vkPrime, witnessPrime)
	if err != nil {
		t.Fatalf("serialised verify error: %s", err)
	}

	t.Log("Serialised proof verification succeeded")
}

func TestPlonkCircuit(t *testing.T) {
	//create pair
	//priv, _ := rand.Int(rand.Reader, ecc.BN254.ScalarField())
	priv := new(big.Int).SetInt64(100)
	pub := new(bn254.PointAffine).ScalarMultiplication(&Base, priv)
	pubArray := pub.Bytes()
	pubStr := hex.EncodeToString(pubArray[:])

	println("pubStr: " + pubStr)
	// weight
	weight := new(big.Int).SetInt64(10)

	// create current bc Votes
	votes := []*big.Int{new(big.Int).SetInt64(1), new(big.Int).SetInt64(1), new(big.Int).SetInt64(1), new(big.Int).SetInt64(1)}
	randoms := []*big.Int{new(big.Int).SetInt64(1111), new(big.Int).SetInt64(1111), new(big.Int).SetInt64(1111), new(big.Int).SetInt64(1111)}
	currentEncVotes := CreateVotes(votes, randoms, pub)

	// println("pub.X.String(): \n", pub.X.String())
	// println("pub.Y.String(): \n", pub.Y.String())
	// println("currentEncVotes.ElGamals[0].Left.X.String(): \n", currentEncVotes.ElGamals[0].Left.X.String())
	// println("currentEncVotes.ElGamals[0].Left.Y.String(): \n", currentEncVotes.ElGamals[0].Left.Y.String())
	// println("currentEncVotes.ElGamals[0].Right.X.String(): \n", currentEncVotes.ElGamals[0].Right.X.String())
	// println("currentEncVotes.ElGamals[0].Right.Y.String(): \n", currentEncVotes.ElGamals[0].Right.Y.String())

	// create add Votes
	addVotes := []*big.Int{new(big.Int).SetInt64(1), new(big.Int).SetInt64(2), new(big.Int).SetInt64(3), new(big.Int).SetInt64(4)}
	addRandoms := []*big.Int{new(big.Int).SetInt64(111), new(big.Int).SetInt64(222), new(big.Int).SetInt64(333), new(big.Int).SetInt64(444)}
	addEncVotes := CreateVotes(addVotes, addRandoms, pub)

	// create new stage of Votes
	newEncVotes := AddVotes(currentEncVotes, addEncVotes)

	//
	arrayLX := newEncVotes.ElGamals[0].Left.X.Bytes()
	fmt.Printf("newEncVotes.ElGamals[0].Left.X.String(): %v\n", hex.EncodeToString(arrayLX[:]))
	arrayLY := newEncVotes.ElGamals[0].Left.Y.Bytes()
	fmt.Printf("newEncVotes.ElGamals[0].Left.Y.String(): %v\n", hex.EncodeToString(arrayLY[:]))
	arrayRX := newEncVotes.ElGamals[0].Right.X.Bytes()
	fmt.Printf("newEncVotes.ElGamals[0].Right.X.String(): %v\n", hex.EncodeToString(arrayRX[:]))
	arrayRY := newEncVotes.ElGamals[0].Right.Y.Bytes()
	fmt.Printf("newEncVotes.ElGamals[0].Right.Y.String(): %v\n", hex.EncodeToString(arrayRY[:]))

	ccs := GetCCSPlonk()
	pk := GetPKPlonk()

	vkBytes, _ := hex.DecodeString(VKStr)
	vk := plonk.NewVerifyingKey(ecc.BN254)
	vk.ReadFrom(bytes.NewReader(vkBytes))

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
