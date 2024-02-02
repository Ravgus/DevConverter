package helpers

import (
	"bufio"
	"log"
	"strings"
)

func GetInput(reader *bufio.Reader) string {
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	input = strings.TrimSpace(input)

	return input
}
