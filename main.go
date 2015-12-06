package main

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Command struct {
	Cmd string `json:"command"`
}

var store Store
var keyInfos map[string]*KeyInfo

func main() {
	store = make(Store)
	r := gin.Default()

	r.POST("/ledis", ledis)
	timeout := make(chan bool, 1)

	go func() {
		time.Sleep(1 * time.Second)
		timeout <- true
	}()

	go func() {
		select {
		case <-timeout:
			for k, v := range keyInfos {
				expired := v.CreatedAt
				for i := 0; i < v.Timeout; i++ {
					expired = expired.Add(1 * time.Second)
				}

				if time.Now().Sub(expired).Seconds() >= 0 {
					store.DEL(k)
				}

			}
			// the read from ch has timed out
		}
	}()

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

	c := strings.ToUpper(args[0])
	k1 := ""
	if c != "SAVE" && c != "RESTORE" && c != "FLUSHDB" {
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
	case "SCARD":
		result = store.SCARD(k1)
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
		result = OK
	case "EXPIRE":
		var timeOut int64
		timeOut, err = strconv.ParseInt(args[2], 10, 64)
		if err != nil {
			err = errors.New("EKTYP")
			return
		}
		keyInfo := KeyInfo{
			Timeout:   int(timeOut),
			CreatedAt: time.Now(),
		}
		if _, ok := keyInfos[k1]; ok {
			result = 1
		} else {
			result = 0
		}
		keyInfos[k1] = &keyInfo
		// return store.EXPIRE(k1, args[2])
	case "TTL":
		if keyInfo, ok := keyInfos[k1]; ok {
			result = keyInfo.Timeout
		} else {
			result = 0
		}
	case "DEL":
		result = store.DEL(k1)
		delete(keyInfos, k1)
	case "FLUSHDB":
		store = make(Store)
		delete(keyInfos, k1)
		result = OK
	}

	return
}
