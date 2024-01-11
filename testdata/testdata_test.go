package testdata

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestGetTestFile(t *testing.T) {
	fmt.Println(GetDigestHex(GetTestFile()))
	fmt.Println(hex.EncodeToString(GetTestFile()))
	fmt.Println(len(GetTestFile()))
}

func TestGetTestPhoneArr(t *testing.T) {
	arr := GetTestPhoneArr()
	fmt.Println(arr)
}

func TestGetTestBankCardNumArr(t *testing.T) {
	arr := GetTestBankCardNumArr()
	fmt.Println(arr)
}
