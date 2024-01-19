package encoding

type Encoder interface {
	Encode(data string) string
	Decode(data string) string
}
