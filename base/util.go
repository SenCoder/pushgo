package base

import (
	"io"
	"time"
	"encoding/hex"
	"crypto/rand"
	"encoding/base64"
	"crypto/sha256"
)

func Hash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	io.WriteString(h, time.Now().String())
	return hex.EncodeToString(h.Sum(nil))
}

func RandomString() string {
	b := make([]byte, 64)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}