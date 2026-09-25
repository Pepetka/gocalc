package calc

import "errors"

// Operation is a named arithmetic operation that accepts a fixed number of operands.
type Operation interface {
	// Name returns the operation's unique name, used for registration and lookup.
	Name() string
	// Arity returns the exact number of operands the operation accepts.
	Arity() int
	// Execute applies the operation to operands and returns the result.
	// It expects exactly Arity() operands; the arity check is the caller's
	// responsibility (Registry.Execute performs it). It returns an error
	// if the operands are otherwise invalid for the operation.
	Execute(operands []float64) (float64, error)
}

// baseOperation implements the metadata common to all builtin operations.
// Concrete operations embed it and provide their own Execute.
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

var (
	// ErrOpRegistered is returned when an operation with the same name is already registered.
	ErrOpRegistered = errors.New("calc: operation already registered")
	// ErrOpNotFound is returned when no operation is registered under the requested name.
	ErrOpNotFound = errors.New("calc: operation not found")
	// ErrInvalidOperandsNum is returned when the number of operands does not match the operation's arity.
	ErrInvalidOperandsNum = errors.New("calc: invalid number of operands")
	// ErrEmptyExpression is returned when the expression is empty.
	ErrEmptyExpression = errors.New("calc: empty expression")
	// ErrDivByZero is returned by the div operation when the divisor is zero.
	ErrDivByZero = errors.New("calc: division by zero")
	// ErrNegativeSqrt is returned by the sqrt operation when the operand is negative.
	ErrNegativeSqrt = errors.New("calc: negative square root")
	// ErrNilOp is returned when a nil operation is passed to Register.
	ErrNilOp = errors.New("calc: nil operation")
)
