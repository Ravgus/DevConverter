package text

import "strings"

type UpperCaseProcessor struct{}

func (UpperCaseProcessor) Process(data string) string {
	return strings.ToUpper(data)
}
