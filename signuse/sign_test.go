package signuse

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"testing"
)

func TestRsaSign(t *testing.T) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	// 源数据
	data := []byte("123abc")
	digested := sha256.Sum256(data)

	// 使用私钥 sha256 + PKCS1v15 算法签名
	sign, err := rsa.SignPKCS1v15(nil, privateKey, crypto.SHA256, digested[:])
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("sign:%s", hex.EncodeToString(sign))

	// 使用公钥验签
	err = rsa.VerifyPKCS1v15(&privateKey.PublicKey, crypto.SHA256, digested[:], sign)
	if err != nil {
		t.Logf("验签失败:%+v", err)
	} else {
		t.Log("验签通过")
	}
}

func TestSign2(t *testing.T) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	data := []byte("123abc")
	digested := sha256.Sum256(data)
	d, err := rsa.DecryptPKCS1v15(nil, privateKey, digested[:])
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("d:%s", hex.EncodeToString(d))
}

func TestEcdsaSign(t *testing.T) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	data := []byte("123abc")
	digested := sha256.Sum256(data)
	// esdsa签名
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, digested[:])
	sign := r.Bytes()
	sign = append(sign, s.Bytes()...)
	t.Logf("sign:%s", hex.EncodeToString(sign))
	// 验签
	isVerify := ecdsa.Verify(&privateKey.PublicKey, digested[:], r, s)
	if isVerify {
		t.Log("验签通过")
	} else {
		t.Log("验签失败")
	}
}
