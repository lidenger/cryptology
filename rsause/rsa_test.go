package rsause

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"testing"
	"time"
)

func TestGenerateKey(t *testing.T) {
	privateKey := GenerateKey(1024)
	t.Logf("1024=>%d", privateKey.Size()*8)

	privateKey = GenerateKey(1027)
	t.Logf("1027=>%d", privateKey.Size()*8)

	privateKey = GenerateKey(2048)
	t.Logf("2048=>%d", privateKey.Size()*8)

	privateKey = GenerateKey(3000)
	t.Logf("3000=>%d", privateKey.Size()*8)

	privateKey = GenerateKey(3073)
	t.Logf("3073=>%d", privateKey.Size()*8)

	privateKey = GenerateKey(4096)
	t.Logf("4096=>%d", privateKey.Size()*8)

	privateKey = GenerateKey(8192)
	t.Logf("8192=>%d", privateKey.Size()*8)

	privateKey = GenerateKey(12345)
	t.Logf("12345=>%d", privateKey.Size()*8)
}

func TestPKCS1(t *testing.T) {
	// 创建一个2048的key
	privateKey := GenerateKey(2048)
	testData := []byte("123abc")
	// 从私钥中拿到公钥进行加密
	cipher, err := rsa.EncryptPKCS1v15(rand.Reader, &privateKey.PublicKey, testData)
	if err != nil {
		t.Fatalf("RSA-PKCS1v15加密失败%s", err)
	}
	t.Logf("RSA-PKCS1v15密文为:%s", hex.EncodeToString(cipher))
	// 使用私钥解密
	origin, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipher)
	if err != nil {
		t.Fatalf("RSA-PKCS1v15解密失败%s", err)
	}
	t.Logf("RSA-PKCS1v15解密结果:%s", string(origin))
}

func TestOAEP(t *testing.T) {
	privateKey := GenerateKey(2048)
	testData := []byte("123abc")
	// 设置一个label，可以是任意数据，加解密保持一致即可
	label := []byte("OAEP test encrypt and decrypt label")
	cipher, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, &privateKey.PublicKey, testData, label)
	if err != nil {
		t.Fatalf("RSA-OAEP加密失败%s", err)
	}
	t.Logf("RSA-OAEP密文为:%s", hex.EncodeToString(cipher))
	origin, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, cipher, label)
	if err != nil {
		t.Fatalf("RSA-OAEP解密失败%s", err)
	}
	t.Logf("RSA-OAEP解密结果:%s", string(origin))
}

func PKCS1EncryptAndDecrypt(privateKey *rsa.PrivateKey, testData string) error {
	cipher, err := rsa.EncryptPKCS1v15(rand.Reader, &privateKey.PublicKey, []byte(testData))
	if err != nil {
		return err
	}
	origin, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipher)
	if err != nil {
		return err
	}
	if string(origin) != testData {
		return errors.New("解密结果不合符预期")
	}
	return nil
}

func OAEPEncryptAndDecrypt(privateKey *rsa.PrivateKey, testData string) error {
	label := []byte("OAEP test encrypt and decrypt label")
	cipher, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, &privateKey.PublicKey, []byte(testData), label)
	if err != nil {
		return err
	}
	origin, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, cipher, label)
	if err != nil {
		return err
	}
	if string(origin) != testData {
		return errors.New("解密结果不合符预期")
	}
	return nil
}

func TestKeySizeEncryptAndDecrypt2(t *testing.T) {
	testData := "13500123421"
	// 2048/3072/4096
	privateKey := GenerateKey(4096)
	start := time.Now().UnixMilli()
	for i := 0; i < 1000; i++ {
		err := PKCS1EncryptAndDecrypt(privateKey, testData)
		if err != nil {
			t.Fatal(err)
		}
	}
	tim := time.Now().UnixMilli() - start
	t.Logf("耗时:%d", tim)
}

func TestKeySizeEncryptAndDecrypt(t *testing.T) {
	testData := "13500123421"
	// 2048/3072/4096
	privateKey := GenerateKey(4096)
	start := time.Now().UnixMilli()
	for i := 0; i < 1000; i++ {
		err := OAEPEncryptAndDecrypt(privateKey, testData)
		if err != nil {
			t.Fatal(err)
		}
	}
	tim := time.Now().UnixMilli() - start
	t.Logf("耗时:%d", tim)
}
