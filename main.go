package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	fmt.Println( //nolint:forbidigo
		primitive.NewObjectID().Hex(),
	)
}
