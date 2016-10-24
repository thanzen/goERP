package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func PasswordMD5(passwd, salt string) string {
	h := md5.New()
	h.Write([]byte(passwd + salt))
	cipherStr := h.Sum(nil)
	result := hex.EncodeToString(cipherStr)
	return result
}
