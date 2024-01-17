package hashing

import (
	"crypto/sha512"
)

type SHA512Hasher struct{}

func (SHA512Hasher) Hash(data string) string {
	return computeHash(sha512.New(), data)
}
