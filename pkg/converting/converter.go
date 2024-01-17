package converting

import (
	"log"
	"strconv"
)

type Converter interface {
	ConvertToDecimal(data string) string
	ConvertToBinary(data string) string
	ConvertToOctal(data string) string
	ConvertToHexadecimal(data string) string
}

func convert(source int, target int, data string) string {
	result, err := strconv.ParseInt(data, source, 64)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return strconv.FormatInt(result, target)
}
