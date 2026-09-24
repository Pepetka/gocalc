// Package rpn is a reverse polish notation calculator.
package rpn

import (
	"strconv"
	"strings"

	"github.com/pepetka/gocalc/calc"
	"github.com/pepetka/gocalc/containers"
)

func Eval(registry *calc.Registry, s string) (float64, error) {
	a := strings.Split(s, " ")
	stack := containers.NewStack[float64](len(a))
	for _, v := range a {
		operand, err := strconv.ParseFloat(v, 64)
		if err == nil {
			stack.Push(operand)
			continue
		}
		op, err := registry.Get(v)
		if err != nil {
			return 0, err
		}
		name := op.Name()
		arity := op.Arity()

		operands := make([]float64, arity)
		for i := range arity {
			operand, ok := stack.Pop()
			if !ok {
				return 0, calc.ErrInvalidOperandsNum
			}
			operands[arity-i-1] = operand
		}

		res, err := registry.Execute(name, operands)
		if err != nil {
			return 0, err
		}
		stack.Push(res)
	}
	if stack.Len() != 1 {
		return 0, calc.ErrInvalidOperandsNum
	}
	res, _ := stack.Pop()
	return res, nil
}
