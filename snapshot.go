package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type KeyInfo struct {
	Timeout   int
	CreatedAt time.Time
}

func (s *Store) SAVE() string {
	file, err := os.Open("/tmp/snapshot")
	if err != nil {
		fmt.Println("error:", err)
	}
	defer file.Close()

	b, err := json.Marshal(*s)
	if err != nil {
		fmt.Println("error:", err)
	}
	file.Write(b)
	return OK
}

func RESTORE() *Store {
	file, err := os.Open("/tmp/snapshot")
	if err != nil {
		fmt.Println("error:", err)
	}
	defer file.Close()

	var b []byte
	file.Read(b)
	var store = make(Store)
	err = json.Unmarshal(b, &store)
	if err != nil {
		fmt.Println("error:", err)
	}
	return &store
}

func (s *Store) DEL(key string) string {
	delete(*s, key)
	return OK
}
