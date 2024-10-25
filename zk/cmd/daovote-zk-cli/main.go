package main

import (
	"flag"

	_ "github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/mrkaurelius/ppp-daovote/zk/zk"
)

func main() {
	flag.Parse()

	zk.SetupCircuit()

	// votePower := *flag.Int("power", 0, "vote power")

	// pubKeyX := *flag.String("pubKeyX", "", "public key x")
	// pubKeyY := *flag.String("pubKeyY", "", "public key y")

	// encBCVote := *flag.String("encBCVote", "", "encBCVote")

	// vote0 := *flag.Int("vote0", 0, "vote1")
	// vote1 := *flag.Int("vote1", 0, "vote2")
	// vote2 := *flag.Int("vote2", 0, "vote3")
	// vote3 := *flag.Int("vote3", 0, "vote4")

	// zk.GenerateProof()

	// generateProof(votePower, vote0, vote1, vote2, vote3, pubKeyX, pubKeyY, encBCVote)
}
