package cipher

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/mitchellh/go-homedir"
)

type TestValut struct {
	encodingKey string
	filepath    string
	keyValues   map[string]string
}

func testFile(key, path string) *TestValut {
	return &TestValut{
		encodingKey: key,
		filepath:    path,
	}
}

func TestEncryptWriter(t *testing.T) {
	key := "testing string"
	file := FilePath()

	vault := testFile(key, file)

	f, err := os.OpenFile(vault.filepath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		t.Error(err)
	}

	Mock1 = true
	_, err = EncryptWriter(vault.encodingKey, f)
	Mock1 = false
	checkError(err)

	Mock2 = true
	_, err = EncryptWriter(vault.encodingKey, f)
	Mock2 = false
	checkError(err)

	Mock3 = true
	_, err = EncryptWriter(vault.encodingKey, f)
	Mock3 = false
	checkError(err)

	Mock4 = true
	_, err = EncryptWriter(vault.encodingKey, f)
	Mock4 = false
	checkError(err)

}

func TestDecryptReader(t *testing.T) {
	key := "testing string"

	file := FilePath()

	vault := testFile(key, file)

	f, err := os.Open(file)
	if err != nil {
		vault.keyValues = make(map[string]string)
	}
	defer f.Close()

	Mock5 = true
	_, err = DecryptReader(vault.encodingKey, f)
	Mock5 = false
	checkError(err)

	Mock6 = true
	_, err = DecryptReader(vault.encodingKey, f)
	Mock6 = false
	checkError(err)

	Mock7 = true
	_, err = DecryptReader(vault.encodingKey, f)
	Mock7 = false
	checkError(err)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func FilePath() string {
	home, _ := homedir.Dir()
	return filepath.Join(home, ".test")
}
