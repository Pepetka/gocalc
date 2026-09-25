// Package rpn evaluates expressions written in reverse Polish notation.
package rpn

import (
	"strconv"
	"strings"

	"github.com/pepetka/gocalc/calc"
	"github.com/pepetka/gocalc/containers"
)

// Eval evaluates a whitespace-separated reverse Polish notation expression
// using the operations registered in registry and returns the result.
// Each token must either parse as a number or name a registered operation,
// otherwise an error is returned. An empty expression yields
// calc.ErrEmptyExpression.
func Eval(registry *calc.Registry, s string) (float64, error) {
	a := strings.Fields(s)
	if len(a) == 0 {
		return 0, calc.ErrEmptyExpression
	}
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
		arity := op.Arity()

		operands := make([]float64, arity)
		for i := range arity {
			operand, ok := stack.Pop()
			if !ok {
				return 0, calc.ErrInvalidOperandsNum
			}
			operands[arity-i-1] = operand
		}

		res, err := op.Execute(operands)
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
