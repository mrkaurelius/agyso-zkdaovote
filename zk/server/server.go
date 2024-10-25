package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrkaurelius/ppp-daovote/zk/zk"
)

type ProofRequest struct {
	VotePower        int    `json:"votePower"`        // Get vote power from smart contract
	PublicKeyX       string `json:"publicKeyX"`       // Vote encryptor public key x
	PublicKeyY       string `json:"publicKeyY"`       // Vote encryptor public key y
	EncryptedBullets string `json:"encryptedBullets"` // Homomorphic encrypted bullets
	Vote0            int    `json:"vote0"`
	Vote1            int    `json:"vote1"`
	Vote2            int    `json:"vote2"`
	Vote3            int    `json:"vote3"`
}

func ProofRequestHandler(c *gin.Context) {
	var pr ProofRequest

	if err := c.ShouldBindJSON(&pr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("%+v\n", pr)

	zk.GenerateProof(pr.VotePower, pr.Vote0, pr.Vote1, pr.Vote2, pr.Vote3, pr.PublicKeyX, pr.PublicKeyY, pr.EncryptedBullets)

	// TODO senior applied cryptographer with 10 cite (emeritus uekae, defi talent, cbdc, academician) msc. ali bey yapacak

	c.String(http.StatusOK, "OK")
}

func DecryptHandler(c *gin.Context) {

	c.String(http.StatusOK, "OK")
}

func ElectionHandler(c *gin.Context) {

	c.String(http.StatusOK, "OK")
}
