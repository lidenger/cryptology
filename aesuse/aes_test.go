package aesuse

import (
	"encoding/base64"
	"github.com/lidenger/cryptology/testdata"
	"testing"
	"time"
)

var (
	// 模拟一个手机号
	data = []byte("15100932122")
	// 长度可选：16|24|32 分别对应AES 128|192|256
	key128 = []byte("1234567890123456")
	key192 = []byte("123456789012345612345678")
	key256 = []byte("12345678901234561234567890123456")
	// 注意长度需要是aes分组块大小
	iv = []byte("1234567890123456")
)

func TestEncrypt(t *testing.T) {
	cipherText, err := Encrypt(key128, iv, data)
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
	origin, err := Decrypt(key128, iv, cipher)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("密文:%s,解密后:%s", cipherStr, string(origin))
}

func encryptAndDecrypt(key, d []byte) {
	cipherText, err := Encrypt(key, iv, d)
	if err != nil {
		panic(err)
	}
	_, err = Decrypt(key, iv, cipherText)
	if err != nil {
		panic(err)
	}
}

func TestKeySizeEncryptAndDecrypt(t *testing.T) {
	start := time.Now().UnixMilli()
	for i := 0; i < 100; i++ {
		encryptAndDecrypt(key256, testdata.GetTestFile())
	}
	tim := time.Now().UnixMilli() - start
	t.Logf("耗时:%d", tim)
}

func TestKeySizeEncryptAndDecrypt2(t *testing.T) {
	start := time.Now().UnixMilli()
	for i := 0; i < 100000; i++ {
		arr := testdata.GetTestPhoneArr()
		for _, phone := range arr {
			encryptAndDecrypt(key256, []byte(phone))
		}
	}
	tim := time.Now().UnixMilli() - start
	t.Logf("耗时:%d", tim)
}

func TestKeySizeEncryptAndDecrypt3(t *testing.T) {
	start := time.Now().UnixMilli()
	for i := 0; i < 100000; i++ {
		arr := testdata.GetTestBankCardNumArr()
		for _, bcn := range arr {
			encryptAndDecrypt(key256, []byte(bcn))
		}
	}
	tim := time.Now().UnixMilli() - start
	t.Logf("耗时:%d", tim)
}
