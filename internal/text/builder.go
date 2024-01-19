package text

const (
	Flip      = "Flip"
	UpperCase = "UpperCase"
	LowerCase = "LowerCase"
)

func Build(algorithm string) Processor {
	switch algorithm {
	case Flip:
		return FlipProcessor{}
	case UpperCase:
		return UpperCaseProcessor{}
	case LowerCase:
		return LowerCaseProcessor{}
	default:
		return nil
	}
}
