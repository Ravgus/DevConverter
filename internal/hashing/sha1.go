package hashing

import (
	"crypto/sha1"
)

type SHA1Hasher struct{}

func (SHA1Hasher) Hash(data string) string {
	return computeHash(sha1.New(), data)
}
