package text

import (
	"bufio"
	"fmt"
	"github.com/Ravgus/DevConverter/internal/helpers"
)

type Executer struct{}

func (Executer) Execute(reader *bufio.Reader) string {
	fmt.Println("Chose what do you want to do: ")
	fmt.Println(Flip)
	fmt.Println(UpperCase)
	fmt.Println(LowerCase)
	fmt.Print("Enter your choice: ")

	action := helpers.GetInput(reader)

	processor := Build(action)

	if processor == nil {
		return "Undefined action"
	}

	fmt.Print("Enter source text: ")

	text := helpers.GetInput(reader)

	return processor.Process(text)
}
