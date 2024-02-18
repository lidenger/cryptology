package testdata

import (
	"crypto/sha256"
	_ "embed"
	"encoding/hex"
	"github.com/brianvoe/gofakeit/v6"
)

//go:embed go1.21.5.src.tar.gz
var TestFile []byte

//go:embed go1.21.6.darwin-arm64.pkg
var TestGo121File []byte

//go:embed github.com.cert
var TestCertFile []byte

// GetTestFile 获取测试文件，25.7MB
func GetTestFile() []byte {
	return TestFile
}

func GetTestGo112File() []byte {
	return TestGo121File
}

func GetTestCertFile() []byte {
	return TestCertFile
}

func GetDigestHex(content []byte) string {
	digest := sha256.Sum256(content)
	return hex.EncodeToString(digest[:])
}

// GetTestPhoneArr 获取一组手机号,10个
func GetTestPhoneArr() []string {
	arr := make([]string, 0)
	for i := 0; i < 10; i++ {
		arr = append(arr, gofakeit.Phone())
	}
	return arr
}

// GetTestBankCardNumArr 获取一组银行卡号,10个
func GetTestBankCardNumArr() []string {
	arr := make([]string, 0)
	for i := 0; i < 10; i++ {
		arr = append(arr, gofakeit.CreditCardNumber(nil))
	}
	return arr
}
