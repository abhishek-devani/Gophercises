package secret

import (
	"log"
	"path/filepath"
	"testing"

	"github.com/mitchellh/go-homedir"
)

func secretpath() string {
	home, _ := homedir.Dir()
	return filepath.Join(home, ".secrets")
}

func TestSet(t *testing.T) {
	// m7, m1, m2, m3, m4
	file := secretpath()
	Vault := File("", file)

	Mock1 = true
	err := Vault.Set("key", "value")
	Mock1 = false
	ErrorCheck(err)

	Mock2 = true
	err = Vault.Set("key", "value")
	Mock2 = false
	ErrorCheck(err)

	Mock3 = true
	err = Vault.Set("key", "value")
	Mock3 = false
	ErrorCheck(err)

	Mock7 = true
	err = Vault.Set("key", "value")
	Mock7 = false
	ErrorCheck(err)

	// Mock4 = true
	// err = Vault.Set("key", "value")
	// Mock4 = false
	// ErrorCheck(err)

	err = Vault.Set("key", "value")

	if err != nil {
		t.Error("Unexpected: ", err)
	}

}

func TestGet(t *testing.T) {
	file := secretpath()
	vault := File("", file)

	Mock5 = true
	_, err := vault.Get("key")
	Mock5 = false
	ErrorCheck(err)

	Mock6 = true
	_, err = vault.Get("key")
	Mock6 = false
	ErrorCheck(err)

	_, err = vault.Get("key")

	if err != nil {
		t.Error("Expected nil but got", err)
	}
}

func ErrorCheck(err error) {
	log.Println(err)
}
