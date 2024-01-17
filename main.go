package main

import (
	"DevConverter/pkg/converting"
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

	fmt.Println("Chose your action:")
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
		result = executeConverting(reader)
	default:
		fmt.Println("Undefined choice")

		return
	}

	fmt.Println("Result:", result)
}

func executeHashing(reader *bufio.Reader) string {
	fmt.Println("Chose hash algorithm:")
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
	fmt.Println("Chose encoding algorithm: ")
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

func executeConverting(reader *bufio.Reader) string {
	fmt.Println("Enter source numbering system: ")
	fmt.Println(converting.DECIMAL)
	fmt.Println(converting.BINARY)
	fmt.Println(converting.OCTAL)
	fmt.Println(converting.HEXADECIMAL)
	fmt.Print("Enter your choice: ")

	initialNS, _ := reader.ReadString('\n')
	initialNS = strings.TrimSpace(initialNS)

	converter := converting.Build(initialNS)

	if converter == nil {
		return "Undefined numbering system"
	}

	fmt.Println("Enter target numbering system: ")
	fmt.Println(converting.DECIMAL)
	fmt.Println(converting.BINARY)
	fmt.Println(converting.OCTAL)
	fmt.Println(converting.HEXADECIMAL)
	fmt.Print("Enter your choice: ")

	finalNS, _ := reader.ReadString('\n')
	finalNS = strings.TrimSpace(finalNS)

	fmt.Print("Enter source value: ")

	value, _ := reader.ReadString('\n')
	value = strings.TrimSpace(value)

	switch finalNS {
	case converting.DECIMAL:
		return converter.ConvertToDecimal(value)
	case converting.BINARY:
		return converter.ConvertToBinary(value)
	case converting.OCTAL:
		return converter.ConvertToOctal(value)
	case converting.HEXADECIMAL:
		return converter.ConvertToHexadecimal(value)
	default:
		return "Undefined target numbering system"
	}
}
