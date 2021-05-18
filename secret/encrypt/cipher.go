package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

// Encrypt will take in a key and plaintext and return a hex representation
// of the excrypted value.
// this code is based on the standard library examples at:
//	- https://golang.org/pkg/crypto/cipher/#NewCBCEncrypter

func Encrypt(key, plaintext string) (string, error) {

	block, err := newCipherBlock(key)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], []byte(plaintext))

	return fmt.Sprintf("%x", ciphertext), nil
}

// Decrypt will take in a key and a cipherHex (hax representation of the
// ciphertext) and decrypt it.
// this code is based on the standard library examples at:
//	- https://golang.org/pkg/crypto/cipher/#NewCBCDecrypter
func Decrypt(key, cipherHex string) (string, error) {

	block, err := newCipherBlock(key)
	if err != nil {
		return "", err
	}

	ciphertext, err := hex.DecodeString(cipherHex)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("encrypt: cipher too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(ciphertext, ciphertext)

	return string(ciphertext), nil
}

func newCipherBlock(key string) (cipher.Block, error) {
	hasher := md5.New()
	fmt.Fprint(hasher, key)
	cipherKey := hasher.Sum(nil)
	// fmt.Println(cipherKey)
	return aes.NewCipher(cipherKey)

}
