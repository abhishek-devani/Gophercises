package main

import (
	"testing"
)

func TestMain(t *testing.T) {

	main()
	Mock = true
	main()
}
