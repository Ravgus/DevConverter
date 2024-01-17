package hashing

import (
	"crypto/md5"
)

type MD5Hasher struct{}

func (MD5Hasher) Hash(data string) string {
	return computeHash(md5.New(), data)
}
