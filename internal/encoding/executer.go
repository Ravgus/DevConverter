package encoding

import (
	"bufio"
	"fmt"
	"github.com/Ravgus/DevConverter/internal/helpers"
)

type Executer struct{}

func (Executer) Execute(reader *bufio.Reader) string {
	fmt.Println("Chose encoding algorithm: ")
	fmt.Println(Base64)
	fmt.Print("Enter your choice: ")

	algorithm := helpers.GetInput(reader)

	encoder := Build(algorithm)

	if encoder == nil {
		return "Undefined encoding algorithm"
	}

	fmt.Print("Enter source text: ")

	text := helpers.GetInput(reader)

	fmt.Println("Do you want to:")
	fmt.Println(Encode)
	fmt.Println(Decode)
	fmt.Print("Enter your choice: ")

	choice := helpers.GetInput(reader)

	switch choice {
	case Encode:
		return encoder.Encode(text)
	case Decode:
		return encoder.Decode(text)
	default:
		return "Undefined choice"
	}
}
