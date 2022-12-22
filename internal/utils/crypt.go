package utils

import (
	"crypto/sha256"
	"fmt"
)

func CryptString(s string, salt string) string {
	pwd := sha256.New()
	pwd.Write([]byte(s))
	pwd.Write([]byte(salt))
	return fmt.Sprintf("%x", pwd.Sum(nil))
}
