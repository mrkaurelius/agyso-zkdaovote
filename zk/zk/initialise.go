package zk

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/consensys/gnark-crypto/ecc"
	bn254 "github.com/consensys/gnark-crypto/ecc/bn254/twistededwards"
	"github.com/consensys/gnark/backend/plonk"
	cs "github.com/consensys/gnark/constraint/bn254"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/scs"
	"github.com/consensys/gnark/test/unsafekzg"
)

const circuitSetupPath = "/var/tmp/agyso-daovote/circuit"

// Setup circuit stuff
func SetupCircuit() {

	err := os.MkdirAll(circuitSetupPath, os.ModePerm)
	if err != nil {
		panic(err)
	}

	priv, _ := rand.Int(rand.Reader, ecc.BN254.ScalarField())
	privateKey := priv.String()

	pub := new(bn254.PointAffine).ScalarMultiplication(&Base, priv)
	pubXArr := pub.X.Bytes()
	pubYArr := pub.Y.Bytes()

	pubXEncoded := hex.EncodeToString(pubXArr[:])
	pubYEncoded := hex.EncodeToString(pubYArr[:])

	keyData := KeyData{PrivateKey: privateKey, PublicKeyX: pubXEncoded, PublicKeyY: pubYEncoded}

	keyDataBytes, err := json.Marshal(keyData)
	if err != nil {
		panic(err)
	}

	keyFile, _ := os.Create(filepath.Join(circuitSetupPath, "keyfile.json"))
	defer keyFile.Close()
	keyFile.Write(keyDataBytes)

	var circuit CircuitMain
	ccs, _ := frontend.Compile(ecc.BN254.ScalarField(), scs.NewBuilder, &circuit)
	r1cs := ccs.(*cs.SparseR1CS)
	srs, srsLagrangeInterpolation, _ := unsafekzg.NewSRS(r1cs)
	pk, vk, _ := plonk.Setup(ccs, srs, srsLagrangeInterpolation)

	var buffCCS bytes.Buffer
	ccs.WriteTo(&buffCCS)
	file0, _ := os.Create(filepath.Join(circuitSetupPath, "ccs.bin"))
	defer file0.Close()
	file0.Write(buffCCS.Bytes())

	var buffPK bytes.Buffer
	pk.WriteTo(&buffPK)
	file1, _ := os.Create(filepath.Join(circuitSetupPath, "pk.bin"))
	defer file1.Close()
	file1.Write(buffPK.Bytes())

	var buffSRS bytes.Buffer
	srs.WriteTo(&buffSRS)
	file2, _ := os.Create(filepath.Join(circuitSetupPath, "srs.bin"))
	defer file2.Close()
	file2.Write(buffSRS.Bytes())

	var buffSRSLag bytes.Buffer
	srsLagrangeInterpolation.WriteTo(&buffSRSLag)
	file3, _ := os.Create(filepath.Join(circuitSetupPath, "srsLag.bin"))
	defer file3.Close()
	file3.Write(buffSRSLag.Bytes())

	var buffVK bytes.Buffer
	vk.WriteTo(&buffVK)

	file4, _ := os.Create(filepath.Join(circuitSetupPath, "vk.bin"))
	defer file4.Close()
	file4.Write(buffVK.Bytes())
}
