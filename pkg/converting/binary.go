package converting

type BinaryConverter struct{}

func (BinaryConverter) ConvertToDecimal(data string) string {
	return convert(2, 10, data)
}

func (BinaryConverter) ConvertToBinary(data string) string {
	return data
}

func (BinaryConverter) ConvertToOctal(data string) string {
	return convert(2, 8, data)
}

func (BinaryConverter) ConvertToHexadecimal(data string) string {
	return convert(2, 16, data)
}
