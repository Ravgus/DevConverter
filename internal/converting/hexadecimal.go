package converting

type HexadecimalConverter struct{}

func (HexadecimalConverter) ConvertToDecimal(data string) string {
	return convert(16, 10, data)
}

func (HexadecimalConverter) ConvertToBinary(data string) string {
	return convert(16, 2, data)
}

func (HexadecimalConverter) ConvertToOctal(data string) string {
	return convert(16, 8, data)
}

func (HexadecimalConverter) ConvertToHexadecimal(data string) string {
	return data
}
