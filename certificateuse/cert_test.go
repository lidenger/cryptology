package certificateuse

import (
	"crypto/x509"
	"encoding/pem"
	"github.com/lidenger/cryptology/testdata"
	"testing"
)

func TestCert(t *testing.T) {
	certBytes := testdata.GetTestCertFile()
	block, _ := pem.Decode(certBytes)
	if block == nil {
		t.Fatal("解析证书数据失败")
	}
	// 获取证书内容
	certificate, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		t.Fatal(err)
	}
	// 获取证书中的公钥
	publicKey := certificate.PublicKey
	t.Logf("public:%+v", publicKey)
}
