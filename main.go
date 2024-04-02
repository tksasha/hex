package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

const (
	objectIDLength = 12
)

func main() {
	id, err := objectID()
	if err != nil {
		fmt.Printf("ERROR -> %v\n", err) //nolint:forbidigo

		return
	}

	fmt.Println(id) //nolint:forbidigo
}

func objectID() (string, error) {
	bytes := make([]byte, objectIDLength)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}
