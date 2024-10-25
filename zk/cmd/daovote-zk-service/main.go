package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mrkaurelius/ppp-daovote/zk/server"
)

func main() {
	r := gin.Default()

	// Init election
	r.POST("/election", server.ElectionHandler)

	r.POST("/vote", server.ProofRequestHandler)

	r.POST("/decrypt", server.DecryptHandler)

	r.Run("localhost:2929")
}
