package encoding

const ENCODE = "Encode"
const DECODE = "Decode"
const BASE64 = "base64"

func Build(algorithm string) Encoder {
	switch algorithm {
	case BASE64:
		return Base64Encoder{}
	default:
		return nil
	}
}
