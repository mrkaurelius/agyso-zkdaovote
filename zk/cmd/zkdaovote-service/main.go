package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mrkaurelius/agyso-zkdaovote/zk/server"
)

func main() {
	r := gin.Default()

	// Init election
	r.POST("/election/init", server.ElectionInitHandler)

	r.POST("/proof/vote", server.GenerateProofHandler)
	r.POST("/proof/submit", server.SubmitProofHandler)
	r.GET("/proof/calldata", server.GetCallDataHandler)

	r.POST("/decrypt", server.DecryptHandler)

	r.Run("localhost:2929")
}
