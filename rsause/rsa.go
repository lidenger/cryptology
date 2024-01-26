package rsause

import (
	"crypto/rand"
	"crypto/rsa"
)

func GenerateKey(len int) *rsa.PrivateKey {
	privateKey, err := rsa.GenerateKey(rand.Reader, len)
	if err != nil {
		panic(err)
	}
	return privateKey
}
