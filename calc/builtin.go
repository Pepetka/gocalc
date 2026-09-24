package calc

import (
	"math"
)

type addOp struct {
	BaseOperation
}

func NewAdd() Operation {
	op := addOp{
		name:  "add",
		arity: 2,
	}
	return op
}

func (op addOp) Execute(operands []float64) (float64, error) {
	if err := op.validate(operands); err != nil {
		return 0, err
	}
	res := 0.0
	for _, operand := range operands {
		res += operand
	}
	return res, nil
}

type mulOp struct {
	BaseOperation
}

func NewMul() Operation {
	op := mulOp{
		name:  "mul",
		arity: 2,
	}
	return op
}

func (op mulOp) Execute(operands []float64) (float64, error) {
	if err := op.validate(operands); err != nil {
		return 0, err
	}
	res := 1.0
	for _, operand := range operands {
		res *= operand
	}
	return res, nil
}

type subOp struct {
	BaseOperation
}

func NewSub() Operation {
	op := subOp{
		name:  "sub",
		arity: 2,
	}
	return op
}

func (op subOp) Execute(operands []float64) (float64, error) {
	if err := op.validate(operands); err != nil {
		return 0, err
	}
	res := operands[0] - operands[1]
	return res, nil
}

type divOp struct {
	BaseOperation
}

func NewDiv() Operation {
	op := divOp{
		name:  "div",
		arity: 2,
	}
	return op
}

func (op divOp) Execute(operands []float64) (float64, error) {
	if err := op.validate(operands); err != nil {
		return 0, err
	}
	res := operands[0] / operands[1]
	return res, nil
}

func (op divOp) validate(operands []float64) error {
	if err := op.BaseOperation.validate(operands); err != nil {
		return err
	}
	if operands[1] == 0 {
		return ErrInvalidOperands
	}
	return nil
}

type powOp struct {
	BaseOperation
}

func NewPow() Operation {
	op := powOp{
		name:  "pow",
		arity: 2,
	}
	return op
}

func (op powOp) Execute(operands []float64) (float64, error) {
	if err := op.validate(operands); err != nil {
		return 0, err
	}
	res := math.Pow(operands[0], operands[1])
	return res, nil
}

type negOp struct {
	BaseOperation
}

func NewNeg() Operation {
	op := negOp{
		name:  "neg",
		arity: 1,
	}
	return op
}

func (op negOp) Execute(operands []float64) (float64, error) {
	if err := op.validate(operands); err != nil {
		return 0, err
	}
	res := -1 * operands[0]
	return res, nil
}

type sqrtOp struct {
	BaseOperation
}

func NewSqrt() Operation {
	op := sqrtOp{
		name:  "sqrt",
		arity: 1,
	}
	return op
}

func (op sqrtOp) Execute(operands []float64) (float64, error) {
	if err := op.validate(operands); err != nil {
		return 0, err
	}
	res := math.Sqrt(operands[0])
	return res, nil
}

func (op sqrtOp) validate(operands []float64) error {
	if err := op.BaseOperation.validate(operands); err != nil {
		return err
	}
	if operands[0] < 0 {
		return ErrInvalidOperands
	}
	return nil
}
