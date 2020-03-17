package text

import (
	"testing"
)

func TestTransform(t *testing.T) {
	utf8 := "你好"
	gbk := Utf8ToGbk(utf8)
	t.Log("gbk: ", gbk)
	if GbkToUtf8(gbk) != utf8 {
		t.Errorf("failed to transform gbk to utf8")
	}
}
