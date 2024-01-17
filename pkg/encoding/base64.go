package encoding

import (
	"encoding/base64"
	"fmt"
	"os"
)

type Base64Encoder struct{}

func (Base64Encoder) Encode(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

func (Base64Encoder) Decode(input string) string {
	data, err := base64.StdEncoding.DecodeString(input)

	if err != nil {
		fmt.Println("Error: ", err)

		os.Exit(-1)
	}

	return string(data)
}
