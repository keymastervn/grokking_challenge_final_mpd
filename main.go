package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.POST("/ledis", ledis)
	r.Run(":8080")
}

func ledis(c *gin.Context) {
	c.JSON(200, gin.H{
		"command": "OK",
	})
}
