package text

type FlipProcessor struct{}

func (FlipProcessor) Process(data string) string {
	runes := []rune(data)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
