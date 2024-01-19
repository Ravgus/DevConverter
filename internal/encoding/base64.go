package encoding

import (
	"encoding/base64"
	"log"
)

type Base64Encoder struct{}

func (Base64Encoder) Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func (Base64Encoder) Decode(data string) string {
	decodeResult, err := base64.StdEncoding.DecodeString(data)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return string(decodeResult)
}
