package converting

type OctalConverter struct{}

func (OctalConverter) ConvertToDecimal(data string) string {
	return convert(8, 10, data)
}

func (OctalConverter) ConvertToBinary(data string) string {
	return convert(8, 2, data)
}

func (OctalConverter) ConvertToOctal(data string) string {
	return data
}

func (OctalConverter) ConvertToHexadecimal(data string) string {
	return convert(8, 16, data)
}
