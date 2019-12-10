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
