package pkcs7

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"testing"
)

func TestPadAndUnPad(t *testing.T) {
	var data [3][]byte
	data[0] = []byte("123abc")
	data[1] = []byte("01234567890123456789012345678912")
	data[2] = []byte("0123456789012345678901234567891201234567890123456789012345678912")
	for _, d := range data {
		padData := Pad(d, aes.BlockSize)
		if len(padData)%aes.BlockSize != 0 {
			t.Fatalf("长度不正确，应该为%d的倍数，当前为:%d", aes.BlockSize, len(padData))
		}
		fmt.Printf("pad base64url:%s\n", base64.URLEncoding.EncodeToString(padData))
		origin, err := Unpad(padData)
		if err != nil {
			t.Fatal(err)
		}
		if string(d) != string(origin) {
			t.Fatalf("unPad和预期不一致,源数据:%s,unPad后:%s", string(d), string(origin))
		}
		fmt.Printf("unPad:%s\n", string(origin))
	}
}
