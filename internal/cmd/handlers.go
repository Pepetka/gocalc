// Package cmd is a package that contains handlers for different modes.
package cmd

import (
	"errors"
	"strconv"

	"github.com/pepetka/gocalc/calc"
	"github.com/pepetka/gocalc/rpn"
)

func RpnHandler(registry *calc.Registry, args []string) (float64, error) {
	if len(args) != 1 {
		return 0, errors.New("invalid number of arguments")
	}
	return rpn.Eval(registry, args[0])
}

func CalcHandler(registry *calc.Registry, args []string) (float64, error) {
	if len(args) < 2 {
		return 0, errors.New("invalid number of arguments")
	}
	op := args[0]
	args = args[1:]
	operands := make([]float64, 0, len(args))
	for _, arg := range args {
		operand, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			return 0, err
		}
		operands = append(operands, operand)
	}
	return registry.Execute(op, operands)
}

func ListHandler(registry *calc.Registry) []calc.BaseOperation {
	return registry.List()
}
