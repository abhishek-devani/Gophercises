package secret

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/abhishek-devani/Gophercises/go/src/github.com/abhishek-devani/Gophercises/secret/cipher"
)

var Mock1 bool
var Mock2 bool
var Mock3 bool
var Mock4 bool
var Mock5 bool
var Mock6 bool
var Mock7 bool

// It will create a new Vault
func File(encodingKey, filepath string) *Vault {
	return &Vault{
		encodingKey: encodingKey,
		filepath:    filepath,
	}
}

// It is public struct
type Vault struct {
	encodingKey string
	filepath    string
	mutex       sync.Mutex
	keyValues   map[string]string
}

// It loads vault and if the map does not exist it creates one
func (v *Vault) Load() error {
	f, err := os.Open(v.filepath)
	if err != nil || Mock1 {
		v.keyValues = make(map[string]string)
		return err
	}
	defer f.Close()
	r, err := cipher.DecryptReader(v.encodingKey, f)
	if err != nil || Mock2 {
		fmt.Println(err)
		return err
	}
	return v.readKeyValues(r)
}

func (v *Vault) readKeyValues(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(&v.keyValues)
}

// It will write encoding key to vault
func (v *Vault) save() error {
	f, err := os.OpenFile(v.filepath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil || Mock3 {
		log.Println(err)
		return err
	}
	defer f.Close()
	w, err := cipher.EncryptWriter(v.encodingKey, f)
	if err != nil || Mock4 {
		log.Println(err)
		return err
	}
	return v.writeKeyValues(w)
}

func (v *Vault) writeKeyValues(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(v.keyValues)
}

// It is used to retrive value for specific key
func (v *Vault) Get(key string) (string, error) {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	err := v.Load()
	if err != nil || Mock5 {
		return "", err
	}
	value, ok := v.keyValues[key]
	if !ok || Mock6 {
		fmt.Println("secret: no value for that key")
		return "", nil
	}
	return value, nil
}

// It is used to store key, value pair in Vault
func (v *Vault) Set(key, value string) error {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	err := v.Load()
	if err != nil || Mock7 {
		log.Println(err)
		return err
	}
	v.keyValues[key] = value
	err = v.save()
	return err
}
