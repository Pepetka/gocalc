// Package cmd implements the command handlers for the gocalc CLI modes.
package cmd

import (
	"errors"
	"strconv"

	"github.com/pepetka/gocalc/calc"
	"github.com/pepetka/gocalc/rpn"
)

// RpnHandler evaluates args[0] as a reverse Polish notation expression.
// It returns an error if args does not contain exactly one expression.
func RpnHandler(registry *calc.Registry, args []string) (float64, error) {
	if len(args) != 1 {
		return 0, errors.New("invalid number of arguments")
	}
	return rpn.Eval(registry, args[0])
}

// CalcHandler executes the operation named by args[0] with the remaining
// arguments parsed as its float64 operands.
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

// ListHandler returns all operations registered in registry.
func ListHandler(registry *calc.Registry) []calc.Operation {
	return registry.List()
}
