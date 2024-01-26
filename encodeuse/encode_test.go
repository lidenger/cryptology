package encodeuse

import (
	"encoding/base64"
	"encoding/hex"
	"testing"
)

func TestHex(t *testing.T) {
	data := []byte("123abc中文")
	// 编码为16进制数据
	code := hex.EncodeToString(data)
	t.Logf("dataLen:%d,code:%s,codeLen:%d", len(data), code, len(code))
	// 解密为源数据
	origin, err := hex.DecodeString(code)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("origin data:%s", string(origin))
}

func TestRawStdEncoding(t *testing.T) {
	data := []byte("123abc中文测试数据123ccddEncodeToString xxc")
	encode := base64.RawStdEncoding.EncodeToString(data)
	t.Logf("encode:%s", encode)
	origin, err := base64.RawStdEncoding.DecodeString(encode)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("decode:%s", string(origin))
}

func TestRawURLEncoding(t *testing.T) {
	data := []byte("123abc中文测试数据123ccddEncodeToString xxc")
	encode := base64.RawURLEncoding.EncodeToString(data)
	t.Logf("encode:%s", encode)
	origin, err := base64.RawURLEncoding.DecodeString(encode)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("decode:%s", string(origin))
}

func TestStdEncoding(t *testing.T) {
	data := []byte("123abc中文测试数据123ccddEncodeToString xxc")
	encode := base64.StdEncoding.EncodeToString(data)
	t.Logf("encode:%s", encode)
	origin, err := base64.StdEncoding.DecodeString(encode)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("decode:%s", string(origin))
}

func TestURLEncoding(t *testing.T) {
	data := []byte("123abc中文测试数据123ccddEncodeToString xxc")
	encode := base64.URLEncoding.EncodeToString(data)
	t.Logf("encode:%s", encode)
	origin, err := base64.URLEncoding.DecodeString(encode)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("decode:%s", string(origin))
}
