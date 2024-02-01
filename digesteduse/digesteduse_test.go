package digesteduse

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"github.com/lidenger/cryptology/testdata"
	"testing"
	"time"
)

func TestSha256_10(t *testing.T) {
	data := []byte("abc123")
	s := sha256.New()
	s.Write(data)
	digest := s.Sum(nil)
	t.Logf("digest:%s", hex.EncodeToString(digest))

	digest2 := sha256.Sum256(data)
	t.Logf("digest2:%s", hex.EncodeToString(digest2[:]))
}

func TestSha(t *testing.T) {
	data := []byte("abc123")
	//sha256.Sum224(data)

	sha256.Sum256(data)
	//sha512.Sum384(data)
	//sha512.Sum512(data)
}

func TestSha256(t *testing.T) {
	data := testdata.GetTestGo112File()
	digested := sha256.Sum256(data)
	hexStr := hex.EncodeToString(digested[:])
	t.Logf("sha256 hex:%s", hexStr)
}

func TestSha256_2(t *testing.T) {
	//digested := sha256.Sum256([]byte("bV%5yX9Z123abcdddddcxxas5yX9Z123abcdddddcxxas5yX9Z123abcdddddcxxas5yX9Z123abcdddddcxxas5yX9Z123abcdddddcxxas5yX9Z123abcdddddcxxas"))
	digested := sha256.Sum256([]byte("bV%5yX9Z123abcdddddcxxas5yX9Z123abcdddddcxxas5yX9Z123abcddd123xxas5yX9Z123abcdddddcxxas5yX9Z123abcdddddcxxas5yX9Z123abcdddddcxxas"))
	t.Log(hex.EncodeToString(digested[:]))
}

func TestMd5(t *testing.T) {
	data := []byte("123abc")
	digested := md5.Sum(data)
	t.Logf("digested:%s", hex.EncodeToString(digested[:]))
}

func TestShaCompare(t *testing.T) {
	testDataArr := testdata.GetTestBankCardNumArr()
	testDataArrBytes := make([][]byte, 0)
	for _, d := range testDataArr {
		b := []byte(d)
		testDataArrBytes = append(testDataArrBytes, b)
	}
	start := time.Now().UnixMilli()
	for i := 0; i < 50000; i++ {
		for _, b := range testDataArrBytes {
			//sha256.Sum224(b)
			//sha256.Sum256(b)
			//sha512.Sum384(b)
			sha512.Sum512(b)
		}
	}
	tim := time.Now().UnixMilli() - start
	t.Logf("耗时:%d", tim)
}

func TestFileShaCompare(t *testing.T) {
	start := time.Now().UnixMilli()
	for i := 0; i < 100; i++ {
		data := testdata.GetTestGo112File()
		//sha256.Sum224(data)
		sha256.Sum256(data)
		//sha512.Sum384(data)
		//sha512.Sum512(data)
	}
	tim := time.Now().UnixMilli() - start
	t.Logf("耗时:%d", tim)
}

func TestMac(t *testing.T) {
	key := []byte("mySecretKey")
	data := []byte("123abc")
	h := hmac.New(sha256.New, key)
	h.Write(data)
	digested := h.Sum(nil)
	t.Logf("digested:%s", hex.EncodeToString(digested))
}
