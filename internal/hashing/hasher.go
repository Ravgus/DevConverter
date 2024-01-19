package hashing

import (
	"encoding/hex"
	"hash"
)

type Hasher interface {
	Hash(data string) string
}

func computeHash(h hash.Hash, data string) string {
	h.Write([]byte(data))

	return hex.EncodeToString(h.Sum(nil))
}
