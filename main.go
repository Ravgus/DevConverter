package main

import (
	"DevConverter/pkg/encoding"
	"DevConverter/pkg/hashing"
	"bufio"
	"fmt"
	"os"
	"strings"
)

const Hashing = "Hashing"
const Encoding = "Encoding"
const Converting = "Converting"

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Chose your purpose:")
	fmt.Println(Hashing)
	fmt.Println(Encoding)
	fmt.Println(Converting)
	fmt.Print("Enter your choice: ")

	purpose, _ := reader.ReadString('\n')
	purpose = strings.TrimSpace(purpose)

	var result string

	switch purpose {
	case Hashing:
		result = executeHashing(reader)
	case Encoding:
		result = executeEncoding(reader)
	case Converting:
		result = ""
	default:
		fmt.Println("Undefined purpose")

		return
	}

	fmt.Println("Result:", result)
}

func executeHashing(reader *bufio.Reader) string {
	fmt.Println("Enter hash algorithm:")
	fmt.Println(hashing.MD5)
	fmt.Println(hashing.SHA1)
	fmt.Println(hashing.SHA256)
	fmt.Println(hashing.SHA512)
	fmt.Print("Enter your choice: ")

	algorithm, _ := reader.ReadString('\n')
	algorithm = strings.TrimSpace(algorithm)

	hasher := hashing.Build(algorithm)

	if hasher == nil {
		return "Undefined hash algorithm"
	}

	fmt.Print("Enter source text: ")

	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	return hasher.Hash(text)
}

func executeEncoding(reader *bufio.Reader) string {
	fmt.Println("Enter encoding algorithm: ")
	fmt.Println(encoding.BASE64)
	fmt.Print("Enter your choice: ")

	algorithm, _ := reader.ReadString('\n')
	algorithm = strings.TrimSpace(algorithm)

	encoder := encoding.Build(algorithm)

	if encoder == nil {
		return "Undefined encoding algorithm"
	}

	fmt.Print("Enter source text: ")

	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	fmt.Println("Do you want to:")
	fmt.Println(encoding.ENCODE)
	fmt.Println(encoding.DECODE)
	fmt.Print("Enter your choice: ")

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case encoding.ENCODE:
		return encoder.Encode(text)
	case encoding.DECODE:
		return encoder.Decode(text)
	default:
		return "Undefined choice"
	}
}
