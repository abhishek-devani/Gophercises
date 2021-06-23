package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Person struct {
	ID primitive.ObjectID
}

func main() {
	fmt.Println("starting the application...")

}
