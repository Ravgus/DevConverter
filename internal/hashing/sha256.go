package hashing

import (
	"crypto/sha256"
)

type SHA256Hasher struct{}

func (SHA256Hasher) Hash(data string) string {
	return computeHash(sha256.New(), data)
}
