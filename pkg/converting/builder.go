package converting

const DECIMAL = "Decimal"
const BINARY = "Binary"
const OCTAL = "Octal"
const HEXADECIMAL = "Hexadecimal"

func Build(numberingSystem string) Converter {
	switch numberingSystem {
	case DECIMAL:
		return DecimalConverter{}
	case BINARY:
		return BinaryConverter{}
	case OCTAL:
		return OctalConverter{}
	case HEXADECIMAL:
		return HexadecimalConverter{}
	default:
		return nil
	}
}
