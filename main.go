package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
)

func main() {
	encrypt()
}

func encrypt() {
	key := genRandomBytes(16)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	iv := genRandomBytes(aes.BlockSize)
	data := genRandomBytes(32)

	mode := cipher.NewCBCEncrypter(block, iv)
	cipher := make([]byte, len(data))
	mode.CryptBlocks(cipher, data)
	base64.URLEncoding.EncodeToString(cipher)
	fmt.Println(base64.RawURLEncoding.EncodeToString(cipher))
}

func genRandomBytes(num int) []byte {
	id := uuid.NewString()
	part := id[:num]
	return []byte(part)
}
