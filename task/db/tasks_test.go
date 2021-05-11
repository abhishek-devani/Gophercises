package db

import (
	"path/filepath"
	"testing"

	"github.com/mitchellh/go-homedir"
)

func TestInit(t *testing.T) {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")

	_, err := Init(dbPath)

	if err != nil {
		t.Fatal(err)
	}

}

func TestCreateTask(t *testing.T) {
	err := CreateTask("Task Create")
	if err != nil {
		panic(err)
	}
}

func TestAllTask(t *testing.T) {
	tasks, err := AllTasks()
	if err != nil {
		t.Fatal(err)
	}
	if tasks == nil {
		t.Fatal("Unexpected Task")
	}
}

func TestDeleteTask(t *testing.T) {
	err := DeleteTasks(1)
	if err != nil {
		t.Fatal(err)
	}
}

var CovertedByte []byte

func TestItob(t *testing.T) {
	CovertedByte = itob(4)
	if CovertedByte == nil {
		t.Fatal("Conversion Failed")
	}
}

func TestBtoi(t *testing.T) {
	CovertedInt := btoi(CovertedByte)
	if CovertedInt == 0 {
		t.Fatal("Conversion Failed")
	}
}

func TestError(t *testing.T) {

	CreateMock = true
	CreateTask("Task Create")

	AllMock = true
	AllTasks()

	InitMock = true
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	_, _ = Init(dbPath)

}
