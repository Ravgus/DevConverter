package converting

const (
	Decimal     = "Decimal"
	Binary      = "Binary"
	Octal       = "Octal"
	Hexadecimal = "Hexadecimal"
)

func Build(numberingSystem string) Converter {
	switch numberingSystem {
	case Decimal:
		return DecimalConverter{}
	case Binary:
		return BinaryConverter{}
	case Octal:
		return OctalConverter{}
	case Hexadecimal:
		return HexadecimalConverter{}
	default:
		return nil
	}
}
