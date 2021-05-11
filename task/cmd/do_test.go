package cmd

import "testing"

func TestDo(t *testing.T) {

	db := startDB()
	defer db.Close()

	a := []string{"1"}
	doCmd.Run(doCmd, a)

}

func TestMock1(t *testing.T) {

	db := startDB()
	defer db.Close()

	a := []string{"1"}
	MockDo1 = true
	doCmd.Run(doCmd, a)
	MockDo1 = false

	MockDo2 = true
	doCmd.Run(doCmd, a)
	MockDo2 = false

	MockDo3 = true
	doCmd.Run(doCmd, a)
	MockDo3 = false

	b := []string{"5"}
	doCmd.Run(doCmd, b)
}
