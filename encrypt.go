package egu


import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/ebar-go/egu/aes"
	"github.com/ebar-go/egu/rsa"
)

//Md5 return the encrypt string by md5 algorithm
func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// Sha1 return hash string
func Sha1(s string) string {
	Sha1Inst := sha1.New()
	Sha1Inst.Write([]byte(s))
	result := Sha1Inst.Sum([]byte(""))
	return fmt.Sprintf("%x", result)
}


// Encrypt
type Encrypt interface {
	// encrypt
	Encrypt(src []byte) ([]byte, error)

	// decrypt
	Decrypt(src []byte) ([]byte, error)
}

// Aes encrypt
func Aes(key []byte) Encrypt {
	return aes.New(key)
}

// Rsa encrypt
func Rsa(public, private []byte) Encrypt  {
	return rsa.New(public, private)
}
