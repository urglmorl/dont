package utils

import "crypto/rand"

func TokenGenerator() (b []byte) {
	b = make([]byte, 32)
	_, _ = rand.Read(b)
	return b
}
