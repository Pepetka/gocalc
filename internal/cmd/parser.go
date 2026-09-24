package cmd

type Mode string

const (
	RpnMode  Mode = "rpn"
	ListMode Mode = "ops"
	CalcMode Mode = "calc"
)

func ParseMode(mode string) Mode {
	switch parsed := Mode(mode); parsed {
	case RpnMode, ListMode:
		return parsed
	default:
		return CalcMode
	}
}
