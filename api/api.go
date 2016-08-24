package api

import "github.com/gin-gonic/gin"

func RunRest() {
	r := gin.Default()
	r.GET("/v1/addresses/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("0.0.0.0:8888") // listen and server on 0.0.0.0:8080
}
