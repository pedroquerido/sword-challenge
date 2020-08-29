package rand

import (
	"math/rand"
	"time"
)

const defaultCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// StringWithCharSet ...
func StringWithCharSet(length int, charset string) string {

	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(b)
}

// String ...
func String(length int) string {
	return StringWithCharSet(length, defaultCharset)
}
