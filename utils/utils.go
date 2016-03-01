package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func MD5(src string) string {
	hash := md5.New()
	hash.Write([]byte(src))
	return string(hex.EncodeToString(hash.Sum(nil)))
}

func GetFileFormat(filename string) string {
	if strings.ContainsRune(filename, '.') {
		filearr := strings.Split(filename, ".")
		return strings.ToLower(filearr[len(filearr)-1])
	}
	return filename
}
