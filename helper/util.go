package helper

import (
	"crypto/md5"
	"encoding/hex"
)

func CreateHash(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	rs := hex.EncodeToString(h.Sum(nil))
	return rs
}

func HashPw(s string) string {
	return CreateHash(s)
}
