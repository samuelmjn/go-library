package utils

import (
	"math/rand"
	"os"
	"strings"
	"time"
)

// GenerateID based on current time
func GenerateID() int64 {
	return int64(time.Now().UnixNano()) + int64(rand.Intn(10000))
}

// UseEnvIfExists :nodoc:
func UseEnvIfExists(keyword string) string {
	if strings.Contains(keyword, "$") {
		return os.Getenv(keyword[1:])
	}
	return keyword
}

// Offset :nodoc:
func Offset(page, limit int64) (offset int64) {
	offset = (page - 1) * limit
	if offset < 0 {
		return 0
	}

	return
}
