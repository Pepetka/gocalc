package calc

import (
	"fmt"
	"math"
)

type addOp struct {
	baseOperation
}

func newAdd() Operation {
	op := addOp{
		name:  "add",
		arity: 2,
	}
	return op
}

func (op addOp) Execute(operands []float64) (float64, error) {
	res := 0.0
	for _, operand := range operands {
		res += operand
	}
	return res, nil
}

type mulOp struct {
	baseOperation
}

func newMul() Operation {
	op := mulOp{
		name:  "mul",
		arity: 2,
	}
	return op
}

func (op mulOp) Execute(operands []float64) (float64, error) {
	res := 1.0
	for _, operand := range operands {
		res *= operand
	}
	return res, nil
}

type subOp struct {
	baseOperation
}

func newSub() Operation {
	op := subOp{
		name:  "sub",
		arity: 2,
	}
	return op
}

func (op subOp) Execute(operands []float64) (float64, error) {
	res := operands[0] - operands[1]
	return res, nil
}

type divOp struct {
	baseOperation
}

func newDiv() Operation {
	op := divOp{
		name:  "div",
		arity: 2,
	}
	return op
}

func (op divOp) Execute(operands []float64) (float64, error) {
	if operands[1] == 0 {
		return 0, fmt.Errorf("%w: %f / %f", ErrDivByZero, operands[0], operands[1])
	}
	res := operands[0] / operands[1]
	return res, nil
}

type powOp struct {
	baseOperation
}

func newPow() Operation {
	op := powOp{
		name:  "pow",
		arity: 2,
	}
	return op
}

func (op powOp) Execute(operands []float64) (float64, error) {
	res := math.Pow(operands[0], operands[1])
	return res, nil
}

type negOp struct {
	baseOperation
}

func newNeg() Operation {
	op := negOp{
		name:  "neg",
		arity: 1,
	}
	return op
}

func (op negOp) Execute(operands []float64) (float64, error) {
	res := -operands[0]
	return res, nil
}

type sqrtOp struct {
	baseOperation
}

func newSqrt() Operation {
	op := sqrtOp{
		name:  "sqrt",
		arity: 1,
	}
	return op
}

func (op sqrtOp) Execute(operands []float64) (float64, error) {
	if operands[0] < 0 {
		return 0, fmt.Errorf("%w: %f", ErrNegativeSqrt, operands[0])
	}
	res := math.Sqrt(operands[0])
	return res, nil
}
