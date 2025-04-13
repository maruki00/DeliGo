package pkgUtils

import (
	"crypto/md5"
	"crypto/sha512"
	"fmt"
)

func Md5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return fmt.Sprintf("%x", hash)
}

func Sha512(text string) string {
	hash := sha512.Sum512([]byte(text))
	return fmt.Sprintf("%x", hash)
}
