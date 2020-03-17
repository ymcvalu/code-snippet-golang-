package md5

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
)

func GetMD5String(s string) string {
	m := md5.New()
	m.Write([]byte(s))

	return hex.EncodeToString(m.Sum(nil))
}

func GetByteMD5String(data []byte) string {
	m := md5.New()
	m.Write(data)
	return hex.EncodeToString(m.Sum(nil))
}

func GetFileMD5String(file string) (string, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}

	m := md5.New()
	m.Write(data)
	return hex.EncodeToString(m.Sum(nil)), nil
}