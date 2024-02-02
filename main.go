package main

import (
	"bufio"
	"fmt"
	"github.com/Ravgus/DevConverter/internal"
	"github.com/Ravgus/DevConverter/internal/helpers"
	"os"
)

func main() {
	printChoseAction()

	reader := bufio.NewReader(os.Stdin)
	executer := internal.Build(helpers.GetInput(reader))

	if executer == nil {
		fmt.Println("Action is not implemented!")

		return
	}

	fmt.Println("Result:", executer.Execute(reader))
}

func printChoseAction() {
	fmt.Println("Chose your action:")
	fmt.Println(internal.Hashing)
	fmt.Println(internal.Encoding)
	fmt.Println(internal.Converting)
	fmt.Println(internal.Text)
	fmt.Print("Enter your choice: ")
}
