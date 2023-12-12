package aesuse

import (
	"crypto/aes"
	"crypto/cipher"
	"github.com/lidenger/cryptology/padding/pkcs7"
)

// Encrypt AES/CBC/PKCS#7
func Encrypt(key, iv, data []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	padData := pkcs7.Pad(data, aes.BlockSize)
	cipherText := make([]byte, len(padData))
	mode.CryptBlocks(cipherText, padData)
	return cipherText, nil
}

// Decrypt AES/CBC/PKCS#7
func Decrypt(key, iv, cipherText []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	data := make([]byte, len(cipherText))
	mode.CryptBlocks(data, cipherText)
	originData, err := pkcs7.Unpad(data)
	if err != nil {
		return nil, err
	}
	return originData, nil
}
