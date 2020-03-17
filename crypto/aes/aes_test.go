package aes

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestAes(t *testing.T) {
	var key = []byte{'K', 'i', 'd', 'k', '3', 'J', '7', 'l', 's', '0', 'v', 'A', 'x', 'e', 'G', '5'}
	origin := []byte("hello world")
	result, err := AesEncrypt(origin, key)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(base64.StdEncoding.EncodeToString(result))
	}
	origData, err := AesDecrypt(result, key)
	if err != nil {
		t.Error(err)
	} else if string(origin) != string(origData) {
		t.Error("失败")
	}
	fmt.Println(string(origData))
}
