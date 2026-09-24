package calc

import "errors"

type Operation interface {
	Name() string
	Arity() int
	Execute(operands []float64) (float64, error)
}

type baseOperation struct {
	name  string
	arity int
}

func (op baseOperation) Name() string {
	return op.name
}

func (op baseOperation) Arity() int {
	return op.arity
}

func (op baseOperation) validate(operands []float64) error {
	s := op.arity
	l := len(operands)
	if l != s {
		return ErrInvalidOperandsNum
	}
	return nil
}

var (
	ErrOpRegistered       = errors.New("calc: operation already registered")
	ErrOpNotFound         = errors.New("calc: operation not found")
	ErrInvalidOperandsNum = errors.New("calc: invalid number of operands")
	ErrDivByZero          = errors.New("calc: division by zero")
	ErrNegativeSqrt       = errors.New("calc: negative square root")
)
