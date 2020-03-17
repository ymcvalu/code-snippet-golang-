package zip

import (
	"archive/zip"
	"bytes"
	"errors"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

var ErrNotZipFile = errors.New("not a zip file")

func Unzip(source, target string) error {
	isZip := IsZip(source)
	if !isZip {
		return ErrNotZipFile
	}

	reader, err := zip.OpenReader(source)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(target, 0666); err != nil {
		return err
	}

	for _, file := range reader.File {
		var filename = file.Name
		if file.Flags == 0 {
			i := bytes.NewReader([]byte(file.Name))
			decoder := transform.NewReader(i, simplifiedchinese.GB18030.NewDecoder())
			byts, _ := ioutil.ReadAll(decoder)
			filename = string(byts)
		}

		path := filepath.Join(target, filename)

		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		if err := func() error {
			fileReader, err := file.Open()
			if err != nil {
				return err
			}
			defer fileReader.Close()

			targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
			if err != nil {
				return err
			}
			defer targetFile.Close()

			if _, err := io.Copy(targetFile, fileReader); err != nil {
				return err
			}

			return nil
		}(); err != nil {
			return err
		}
	}
	return nil
}

func IsZip(path string) bool {
	f, err := os.Open(path)
	if err != nil {
		return false
	}

	defer f.Close()
	buf := make([]byte, 4)
	if n, err := f.Read(buf); err != nil || n < 4 {
		return false
	}
	return bytes.Equal(buf, []byte("PK\x03\x04"))
}
