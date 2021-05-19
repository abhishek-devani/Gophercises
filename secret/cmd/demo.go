package main

import (
	"fmt"

	"github.com/abhishek-devani/Gophercises/go/src/github.com/abhishek-devani/Gophercises/secret"
)

func main() {
	v := secret.File("my-fake-key", ".secrets")

	// err := v.Set("demo_key1", "123")
	// if err != nil {
	// 	panic(err)
	// }

	// err = v.Set("demo_key2", "456")
	// if err != nil {
	// 	panic(err)
	// }

	// err = v.Set("demo_key3", "789")
	// if err != nil {
	// 	panic(err)
	// }

	plain, err := v.Get("demo_key1")
	if err != nil {
		panic(err)
	}

	fmt.Println("Plain:", plain)
	plain, err = v.Get("demo_key2")
	if err != nil {
		panic(err)
	}

	fmt.Println("Plain:", plain)
	plain, err = v.Get("demo_key3")
	if err != nil {
		panic(err)
	}

	fmt.Println("Plain:", plain)
}
