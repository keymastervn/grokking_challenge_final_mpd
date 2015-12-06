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

func main() {
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
	args := strings.Split(cmd, " ")
	if len(args) == 0 {
		err = errors.New("ECOM")
		return
	}
	var store Store
	make(store)

	c := strings.ToUpper(args[0])
	k1 := args[1]
	if k1 != strings.ToLower(k1) {
		err = errors.New("EKTYP")
		return
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
		// return store.SADD(args)
	case "SMEMBERS":
		// return store.SMEMBERS(k1)
	case "SREM":
		// return store.SREM(args)
	case "SINTER":
		for _, key := range args {
			if key != strings.ToLower(key) {
				err = errors.New("EKTYP")
				return
			}
		}
		// return store.SINTER(args)
	case "SAVE":
		// return store.SAVE()
	case "RESTORE":
		// return store.RESTORE()
	case "EXPIRE":
		// return store.EXPIRE(k1, args[2])
	case "TTL":
		// return store.TTL(k1)
	case "DEL":
		// return store.DEL(k1)
	case "FLUSHDB":
		// return store.FLUSHDB()
	}

	return
}
