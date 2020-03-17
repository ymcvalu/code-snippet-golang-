package crypto

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"github.com/tjfoc/gmsm/sm3"
	"io"
)

func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func GetSM3String(s string) string {
	hw := sm3.New()
	hw.Write([]byte(s))
	return hex.EncodeToString(hw.Sum(nil))
}

func GetByteMD5String(data []byte) string {
	m := md5.New()
	m.Write(data)
	return hex.EncodeToString(m.Sum(nil))
}

func GetRandString(len int) string {
	b := make([]byte, len)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}

func MaskPassword(pass string) string {
	sum := sha256.Sum256([]byte(pass))
	hash := md5.Sum(sum[:])
	return hex.EncodeToString(hash[:])
}
