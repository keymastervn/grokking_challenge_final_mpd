package main

import (
	"errors"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Command struct {
	Cmd string `json:"command"`
}

var store Store

func main() {
	store = make(Store)
	r := gin.Default()

	r.POST("/ledis", ledis)
	r.Run(":80")
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

func Execute(cmd string) (result interface{}, err error) {
	if len(cmd) == 0 {
		err = errors.New("ECOM")
		return
	}

	args := strings.Split(cmd, " ")
	if len(args) == 0 {
		err = errors.New("ECOM")
		return
	}
	var k1 string

	c := strings.ToUpper(args[0])
	if c != "SAVE" || c != "RESTORE" {
		k1 = args[1]

		if k1 != strings.ToLower(k1) {
			err = errors.New("EKTYP")
			return
		}
	}

	switch c {
	case "SET":
		result = store.StringSet(k1, args[2])
	case "GET":
		result = store.StringGet(k1)
	case "LLEN":
		result = store.LLEN(k1)
	case "RPUSH":
		result = store.RPUSH(k1, args[2:]...)
	case "LPOP":
		result = store.LPOP(k1)
	case "RPOP":
		result = store.RPOP(k1)
	case "LRANGE":
		start, _ := strconv.Atoi(args[2])
		stop, _ := strconv.Atoi(args[3])
		result, err = store.LRANGE(k1, start, stop)
	case "SADD":
		result = store.SADD(k1, args[2:]...)
		// return store.SADD(args)
	case "SMEMBERS":
		result = store.SMEMBERS(k1)
	case "SREM":
		// return store.SREM(args)
	case "SINTER":
		for _, k := range args[2:] {
			if k != strings.ToLower(k) {
				err = errors.New("EKTYP")
				return
			}
		}

		result = store.SINTER(args[1:]...)
	case "SAVE":
		result = store.SAVE()
	case "RESTORE":
		store = *RESTORE()
	case "EXPIRE":
		// return store.EXPIRE(k1, args[2])
	case "TTL":
		// return store.TTL(k1)
	case "DEL":
		result = store.DEL(k1)
	case "FLUSHDB":
		store = make(Store)
		result = OK
	}

	return
}
