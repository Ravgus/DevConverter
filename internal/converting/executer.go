package converting

import (
	"bufio"
	"fmt"
	"github.com/Ravgus/DevConverter/internal/helpers"
)

type Executer struct{}

func (Executer) Execute(reader *bufio.Reader) string {
	fmt.Println("Enter source numbering system: ")
	fmt.Println(Decimal)
	fmt.Println(Binary)
	fmt.Println(Octal)
	fmt.Println(Hexadecimal)
	fmt.Print("Enter your choice: ")

	initialNS := helpers.GetInput(reader)

	converter := Build(initialNS)

	if converter == nil {
		return "Undefined numbering system"
	}

	fmt.Println("Enter target numbering system: ")
	fmt.Println(Decimal)
	fmt.Println(Binary)
	fmt.Println(Octal)
	fmt.Println(Hexadecimal)
	fmt.Print("Enter your choice: ")

	finalNS := helpers.GetInput(reader)

	fmt.Print("Enter source value: ")

	value := helpers.GetInput(reader)

	switch finalNS {
	case Decimal:
		return converter.ConvertToDecimal(value)
	case Binary:
		return converter.ConvertToBinary(value)
	case Octal:
		return converter.ConvertToOctal(value)
	case Hexadecimal:
		return converter.ConvertToHexadecimal(value)
	default:
		return "Undefined target numbering system"
	}
}
