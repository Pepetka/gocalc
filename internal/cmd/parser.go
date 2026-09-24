package cmd

// Mode identifies a CLI operating mode.
type Mode string

const (
	// RpnMode evaluates a single reverse Polish notation expression.
	RpnMode Mode = "rpn"
	// ListMode prints all registered operations.
	ListMode Mode = "ops"
	// CalcMode executes a single operation with the given operands.
	CalcMode Mode = "calc"
)

// ParseMode maps the first CLI argument to a Mode.
// Any value other than RpnMode or ListMode selects CalcMode.
func ParseMode(mode string) Mode {
	switch parsed := Mode(mode); parsed {
	case RpnMode, ListMode:
		return parsed
	default:
		return CalcMode
	}
}
