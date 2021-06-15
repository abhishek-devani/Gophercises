package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"io"
)

var Mock1 bool
var Mock2 bool
var Mock3 bool
var Mock4 bool
var Mock5 bool
var Mock6 bool
var Mock7 bool

func encryptStream(key string, iv []byte) (cipher.Stream, error) {
	block, err := newCipherBlock(key)
	if err != nil || Mock1 {
		return nil, err
	}
	return cipher.NewCFBEncrypter(block, iv), nil
}

// EncryptWriter will return a writer that will write encrypted data to
// the original writer.
func EncryptWriter(key string, w io.Writer) (*cipher.StreamWriter, error) {
	iv := make([]byte, aes.BlockSize)
	_, err := io.ReadFull(rand.Reader, iv)
	if err != nil || Mock2 {
		return nil, err
	}
	stream, err := encryptStream(key, iv)
	if err != nil || Mock3 {
		return nil, err
	}
	n, err := w.Write(iv)
	if n != len(iv) || err != nil || Mock4 {
		return nil, err
	}
	return &cipher.StreamWriter{S: stream, W: w}, nil
}

func decryptStream(key string, iv []byte) (cipher.Stream, error) {
	block, err := newCipherBlock(key)
	if err != nil || Mock5 {
		return nil, err
	}
	return cipher.NewCFBDecrypter(block, iv), nil
}

// DecryptReader will return a reader that will decrypt data from the
// provided reader and give the user a way to read that data as it if was
// not encrypted.
func DecryptReader(key string, r io.Reader) (*cipher.StreamReader, error) {
	iv := make([]byte, aes.BlockSize)
	n, err := r.Read(iv)
	if n < len(iv) || err != nil || Mock6 {
		// fmt.Println("hahaha")
		return nil, err
		// errors.New("encrypt: unable to read the full iv")
	}
	stream, err := decryptStream(key, iv)
	if err != nil || Mock7 {
		return nil, err
	}
	return &cipher.StreamReader{S: stream, R: r}, nil
}

// This function return cipher block containing hashed version of key
func newCipherBlock(key string) (cipher.Block, error) {
	hasher := md5.New()
	fmt.Fprint(hasher, key)
	cipherKey := hasher.Sum(nil)
	return aes.NewCipher(cipherKey)
}
