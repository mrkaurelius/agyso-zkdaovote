package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrkaurelius/agyso-zkdaovote/zk/zk"
)

type ProofRequest struct {
	VotePower        int    `json:"votePower"`        // Get vote power from smart contract
	EncryptedBallots string `json:"encryptedBallots"` // Homomorphic encrypted ballots
	Vote0            int    `json:"vote0"`
	Vote1            int    `json:"vote1"`
	Vote2            int    `json:"vote2"`
	Vote3            int    `json:"vote3"`
}

type DecryptRequest struct {
	EncrytedVotes string `json:"encryptedVotes"`
}

func GenerateProofHandler(c *gin.Context) {
	var pr ProofRequest

	if err := c.ShouldBindJSON(&pr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("%+v\n", pr)

	electionKeys, err := zk.GetElectionKeys()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	err = zk.GenerateProof(pr.VotePower, pr.Vote0, pr.Vote1, pr.Vote2, pr.Vote3, electionKeys.PublicKeyX,
		electionKeys.PublicKeyY, pr.EncryptedBallots)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "proof generated"})
}

func SubmitProofHandler(c *gin.Context) {

	calldata, err := zk.ExecAgysoDaoVoteRs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "proof submitted to aligned and verified", "calldata": calldata})
}

func GetCallDataHandler(c *gin.Context) {
	calldata, err := zk.GetCallData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "calldata": calldata})
}

func DecryptHandler(c *gin.Context) {

	var dr DecryptRequest

	if err := c.ShouldBindJSON(&dr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	key, err := zk.GetElectionKeys()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	decryptedVotes := zk.DecryptEncryprtedBallotBox(dr.EncrytedVotes, key.PrivateKey)

	c.JSON(http.StatusOK, gin.H{"status": "ok", "decryptedVotes": decryptedVotes})
}

func ElectionInitHandler(c *gin.Context) {

	c.String(http.StatusOK, "OK")
}
