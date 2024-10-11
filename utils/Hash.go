package utils

import (
	"crypto/sha256"
	"fmt"
)

func CreateHash(str string) string {
	hash := sha256.New()
	hash.Write([]byte(str))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
