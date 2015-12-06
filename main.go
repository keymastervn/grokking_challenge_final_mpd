package main

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

type Command struct {
	Cmd string `json:"command"`
}

func main() {
	r := gin.Default()
	r.POST("/ledis", ledis)
	r.Run(":8080")
}

func ledis(c *gin.Context) {
	// validate content-type & content-length
	var command Command
	if c.BindJSON(&command) == nil {
		result, err := Execute(command.Cmd)
		if err != nil {
			c.JSON(200, gin.H{
				"response": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"response": result,
			})
		}
	} else {
		c.JSON(200, gin.H{
			"response": "ECOM",
		})
	}
}

func Execute(cmd string) (store map[string]interface{}, err error) {
	args := strings.Split(cmd, " ")
	if len(args) == 0 {
		err = errors.New("ECOM")
		return
	}

	c := strings.ToUpper(args[0])
	k1 := args[1]
	if k1 != strings.ToLower(k1) {
		err = errors.New("ECOM")
		return
	}

	if c == "SET" {

	}

	return
}
