package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mrkaurelius/agyso-zkdaovote/zk/server"
)

func main() {
	r := gin.Default()

	r.Use(CORSMiddleware())

	// Init election
	r.POST("/election/init", server.ElectionInitHandler)

	r.POST("/proof/vote", server.GenerateProofHandler)
	r.POST("/proof/submit", server.SubmitProofHandler)
	r.GET("/proof/calldata", server.GetCallDataHandler)

	r.POST("/decrypt", server.DecryptHandler)

	r.Run("localhost:2929")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
