package hashing

import (
	"bufio"
	"fmt"
	"github.com/Ravgus/DevConverter/internal/helpers"
)

type Executer struct{}

func (Executer) Execute(reader *bufio.Reader) string {
	fmt.Println("Chose hash algorithm:")
	fmt.Println(Md5)
	fmt.Println(Sha1)
	fmt.Println(Sha156)
	fmt.Println(Sha512)
	fmt.Print("Enter your choice: ")

	algorithm := helpers.GetInput(reader)

	hasher := Build(algorithm)

	if hasher == nil {
		return "Undefined hash algorithm"
	}

	fmt.Print("Enter source text: ")

	text := helpers.GetInput(reader)

	return hasher.Hash(text)
}
