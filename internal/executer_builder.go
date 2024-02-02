package internal

import (
	"bufio"
	"github.com/Ravgus/DevConverter/internal/converting"
	"github.com/Ravgus/DevConverter/internal/encoding"
	"github.com/Ravgus/DevConverter/internal/hashing"
	"github.com/Ravgus/DevConverter/internal/text"
)

const (
	Hashing    = "Hashing"
	Encoding   = "Encoding"
	Converting = "Converting"
	Text       = "Text"
)

type Executer interface {
	Execute(reader *bufio.Reader) string
}

func Build(purpose string) Executer {
	switch purpose {
	case Hashing:
		return hashing.Executer{}
	case Encoding:
		return encoding.Executer{}
	case Converting:
		return converting.Executer{}
	case Text:
		return text.Executer{}
	default:
		return nil
	}
}
