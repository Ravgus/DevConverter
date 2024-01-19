package encoding

const (
	Encode = "Encode"
	Decode = "Decode"
	Base64 = "base64"
)

func Build(algorithm string) Encoder {
	switch algorithm {
	case Base64:
		return Base64Encoder{}
	default:
		return nil
	}
}
