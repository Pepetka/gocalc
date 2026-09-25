// Package calc provides a registry of named arithmetic operations over float64 operands.
package calc

import (
	"slices"
	"strings"
)

// Registry stores operations by name and executes them on demand.
// Use New to create one pre-populated with the builtin operations.
type Registry struct {
	ops map[string]Operation
}

// New returns a Registry with all builtin operations registered:
// add, mul, sub, div, pow, neg and sqrt.
func New() *Registry {
	r := &Registry{
		ops: make(map[string]Operation),
	}
	r.registerBuiltin()
	return r
}

func (r *Registry) registerBuiltin() {
	_ = r.Register(newAdd())
	_ = r.Register(newMul())
	_ = r.Register(newSub())
	_ = r.Register(newDiv())
	_ = r.Register(newPow())
	_ = r.Register(newNeg())
	_ = r.Register(newSqrt())
}

// Register adds op to the registry under op.Name().
// It returns ErrNilOp if op is nil, or ErrOpRegistered if the name
// is already taken.
func (r *Registry) Register(op Operation) error {
	if op == nil {
		return ErrNilOp
	}
	n := op.Name()
	if _, ok := r.ops[n]; ok {
		return ErrOpRegistered
	}
	r.ops[n] = op
	return nil
}

// Get returns the operation registered under name, or ErrOpNotFound
// if the name is unknown.
func (r *Registry) Get(name string) (Operation, error) {
	op, ok := r.ops[name]
	if !ok {
		return nil, ErrOpNotFound
	}
	return op, nil
}

// Execute runs the operation registered under name with the given operands.
// It returns ErrOpNotFound if the name is unknown, or the operation's own
// error if the operands are invalid.
func (r *Registry) Execute(name string, operands []float64) (float64, error) {
	op, ok := r.ops[name]
	if !ok {
		return 0, ErrOpNotFound
	}
	return op.Execute(operands)
}

// List returns all registered operations sorted by name.
func (r *Registry) List() []Operation {
	l := make([]Operation, 0, len(r.ops))
	for _, op := range r.ops {
		l = append(l, op)
	}
	slices.SortFunc(l, func(a, b Operation) int {
		return strings.Compare(a.Name(), b.Name())
	})
	return l
}
