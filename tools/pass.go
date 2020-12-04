package tools

import (
	"crypto/sha256"
	"encoding/hex"
)

// 密码加密
func GeneratePass(pass string) string {
	sum256 := sha256.Sum256([]byte(pass))
	return hex.EncodeToString(sum256[:])
}
