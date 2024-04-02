package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func main() {
	id, err := objectID()
	if err != nil {
		fmt.Printf("ERROR -> %v\n", err)

		return
	}

	fmt.Println(id)
}

func objectID() (string, error) {
	bytes := make([]byte, 12)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}
