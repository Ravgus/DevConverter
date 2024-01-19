package main

import (
	"bufio"
	"fmt"
	"github.com/Ravgus/DevConverter/internal/converting"
	"github.com/Ravgus/DevConverter/internal/encoding"
	"github.com/Ravgus/DevConverter/internal/hashing"
	"log"
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

	purpose := getInput(reader)

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
	fmt.Println(hashing.Md5)
	fmt.Println(hashing.Sha1)
	fmt.Println(hashing.Sha156)
	fmt.Println(hashing.Sha512)
	fmt.Print("Enter your choice: ")

	algorithm := getInput(reader)

	hasher := hashing.Build(algorithm)

	if hasher == nil {
		return "Undefined hash algorithm"
	}

	fmt.Print("Enter source text: ")

	text := getInput(reader)

	return hasher.Hash(text)
}

func executeEncoding(reader *bufio.Reader) string {
	fmt.Println("Chose encoding algorithm: ")
	fmt.Println(encoding.Base64)
	fmt.Print("Enter your choice: ")

	algorithm := getInput(reader)

	encoder := encoding.Build(algorithm)

	if encoder == nil {
		return "Undefined encoding algorithm"
	}

	fmt.Print("Enter source text: ")

	text := getInput(reader)

	fmt.Println("Do you want to:")
	fmt.Println(encoding.Encode)
	fmt.Println(encoding.Decode)
	fmt.Print("Enter your choice: ")

	choice := getInput(reader)

	switch choice {
	case encoding.Encode:
		return encoder.Encode(text)
	case encoding.Decode:
		return encoder.Decode(text)
	default:
		return "Undefined choice"
	}
}

func executeConverting(reader *bufio.Reader) string {
	fmt.Println("Enter source numbering system: ")
	fmt.Println(converting.Decimal)
	fmt.Println(converting.Binary)
	fmt.Println(converting.Octal)
	fmt.Println(converting.Hexadecimal)
	fmt.Print("Enter your choice: ")

	initialNS := getInput(reader)

	converter := converting.Build(initialNS)

	if converter == nil {
		return "Undefined numbering system"
	}

	fmt.Println("Enter target numbering system: ")
	fmt.Println(converting.Decimal)
	fmt.Println(converting.Binary)
	fmt.Println(converting.Octal)
	fmt.Println(converting.Hexadecimal)
	fmt.Print("Enter your choice: ")

	finalNS := getInput(reader)

	fmt.Print("Enter source value: ")

	value := getInput(reader)

	switch finalNS {
	case converting.Decimal:
		return converter.ConvertToDecimal(value)
	case converting.Binary:
		return converter.ConvertToBinary(value)
	case converting.Octal:
		return converter.ConvertToOctal(value)
	case converting.Hexadecimal:
		return converter.ConvertToHexadecimal(value)
	default:
		return "Undefined target numbering system"
	}
}

func getInput(reader *bufio.Reader) string {
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	input = strings.TrimSpace(input)

	return input
}
