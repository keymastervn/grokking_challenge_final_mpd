package main

import "github.com/gin-gonic/gin"

func checkContentType(c *gin.Context) {
	if len(c.requestHeader("Content-Type")) == 0 {
		return false
	}
	return true
}
