package zip

import (
	"bytes"
	"testing"
)

func TestCompressor(t *testing.T) {
	c, err := NewCompressor("./test.zip")
	if err != nil {
		t.Errorf("failed to create compressor: %s", err.Error())
		return
	}

	defer c.Close()

	buf := bytes.NewBuffer([]byte("this is a new file"))

	// the path test/test.txt in zip archive
	err = c.AddStream("test/", BaseFileInfo("text.txt", int64(buf.Len())), buf)
	if err != nil {
		t.Errorf("failed to add file into zip archive: %s", err.Error())
		return
	}
}
