package cmd

import (
	"testing"
)

func TestList(t *testing.T) {
	db := startDB()
	defer db.Close()

	a := []string{""}
	listCmd.Run(listCmd, a)
}

func TestMockList(t *testing.T) {

	db := startDB()
	defer db.Close()

	a := []string{""}
	MockList1 = true
	listCmd.Run(listCmd, a)
	MockList1 = false

	MockList2 = true
	listCmd.Run(listCmd, a)
	MockList2 = false
}
