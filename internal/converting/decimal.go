package converting

type DecimalConverter struct{}

func (DecimalConverter) ConvertToDecimal(data string) string {
	return data
}

func (DecimalConverter) ConvertToBinary(data string) string {
	return convert(10, 2, data)
}

func (DecimalConverter) ConvertToOctal(data string) string {
	return convert(10, 8, data)
}

func (DecimalConverter) ConvertToHexadecimal(data string) string {
	return convert(10, 16, data)
}
