package text

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
)

func GbkToUtf8(s string) string {
	reader := transform.NewReader(bytes.NewBufferString(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return s
	}
	return string(d)
}

func Utf8ToGbk(s string) (string) {
	reader := transform.NewReader(bytes.NewBufferString(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return s
	}
	return string(d)
}
