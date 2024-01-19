package text

type Processor interface {
	Process(data string) string
}
