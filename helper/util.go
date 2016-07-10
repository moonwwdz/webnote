package helper

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/russross/blackfriday"
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

func Markdown(s string) string {
	output := blackfriday.MarkdownCommon([]byte(s))
	return string(output)
}
