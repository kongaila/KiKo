package tools

import (
	"github.com/satori/go.uuid"
	"strings"
)

// 统一使用此方式生成32位uuid
func GenerateUUID() string {
	return strings.ReplaceAll(uuid.NewV4().String(), "-", "")
}

// 统一使用此方式生成32位uuid
func GenerateUserNick() string {
	return strings.ReplaceAll(uuid.NewV4().String(), "-", "")[0:5]
}
