package zip

import (
	"archive/zip"
	"errors"
	"io"
	"os"
	"path/filepath"
	"time"
)

type Compressor struct {
	path string
	fd   *os.File
	zw   *zip.Writer
}

// @path: where save the zip archive
func NewCompressor(path string) (*Compressor, error) {
	fd, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	return &Compressor{
		path: path,
		fd:   fd,
		zw:   zip.NewWriter(fd),
	}, nil
}

func (c *Compressor) Close() error {
	err := c.zw.Close()
	if err == nil {
		err = c.fd.Close()
	}
	return err
}

// @base: the base path in zip archive
// @fi: the file info
// @reader: stream for writing into zip
func (c *Compressor) AddStream(base string, fi os.FileInfo, reader io.Reader) error {
	header, err := zip.FileInfoHeader(fi)
	if err != nil {
		return err
	}
	path := filepath.Join(base, fi.Name())
	header.Name = path
	header.Method = zip.Deflate // 设置压缩方法，默认归档不压缩
	writer, err := c.zw.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, reader)
	return err
}

var ErrIsDirectory = errors.New("the file is a directory")
var ErrNotDirectory = errors.New("the file isn't a directory")

func (c *Compressor) AddFile(base string, path string) error {
	fd, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fd.Close()
	stat, err := fd.Stat()
	if err != nil {
		return err
	}

	if stat.IsDir() {
		return ErrIsDirectory
	}

	return c.AddStream(base, stat, fd)
}

func (c *Compressor) AddDir(base string, path string, skip func(path string, fi os.FileInfo) bool) error {
	stat, err := os.Stat(path)
	if err != nil {
		return err
	}
	if !stat.IsDir() {
		return ErrNotDirectory
	}

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if skip != nil && skip(path, info) {
			return nil
		}

		return c.AddFile(base, path)
	})
	return nil
}

var _ os.FileInfo = new(MemoryFileInfo)

type MemoryFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func BaseFileInfo(name string, size int64) os.FileInfo {
	return &MemoryFileInfo{
		name:    name,
		size:    size,
		mode:    0666,
		modTime: time.Now(),
	}
}

func NewMemoryFileInfo(name string, size int64, mode os.FileMode, tm time.Time) os.FileInfo {
	return &MemoryFileInfo{
		name:    name,
		size:    size,
		mode:    mode,
		modTime: tm,
	}
}

func (fi *MemoryFileInfo) Name() string {
	return fi.name
}

func (fi *MemoryFileInfo) Size() int64 {
	return fi.size
}

func (fi *MemoryFileInfo) Mode() os.FileMode {
	return fi.mode
}

func (fi *MemoryFileInfo) ModTime() time.Time {
	return fi.modTime
}

func (fi *MemoryFileInfo) IsDir() bool {
	return false
}

func (fi *MemoryFileInfo) Sys() interface{} {
	return fi
}
