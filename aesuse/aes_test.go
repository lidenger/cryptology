package aesuse

import (
	"encoding/base64"
	"testing"
)

var (
	// 模拟一个手机号
	data = []byte("15100932122")
	// 长度可选：16|32|64 分别对应AES 128|256|512
	key = []byte("1234567890123456")
	// 注意长度需要是aes分组块大小
	iv = []byte("1234567890123456")
)

func TestEncrypt(t *testing.T) {
	cipherText, err := Encrypt(key, iv, data)
	if err != nil {
		t.Fatal(err)
	}
	cipherStr := base64.RawURLEncoding.EncodeToString(cipherText)
	t.Logf("data:%s,加密后:%s\n", string(data), cipherStr)
}

func TestDecrypt(t *testing.T) {
	cipherStr := "vVpKSF-8FmFszFmKVA_HxA"
	cipher, err := base64.RawURLEncoding.DecodeString(cipherStr)
	if err != nil {
		t.Fatal(err)
	}
	origin, err := Decrypt(key, iv, cipher)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("密文:%s,解密后:%s", cipherStr, string(origin))
}
