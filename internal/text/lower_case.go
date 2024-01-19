package text

import "strings"

type LowerCaseProcessor struct{}

func (LowerCaseProcessor) Process(data string) string {
	return strings.ToLower(data)
}
