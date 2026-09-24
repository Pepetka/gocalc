// Package calc is a calculator package that allows you to register and execute operations.
package calc

type Registry struct {
	ops map[string]Operation
}

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

func (r *Registry) Register(op Operation) error {
	n := op.Name()
	if _, ok := r.ops[n]; ok {
		return ErrOpRegistered
	}
	r.ops[n] = op
	return nil
}

func (r *Registry) Get(name string) (Operation, error) {
	op, ok := r.ops[name]
	if !ok {
		return nil, ErrOpNotFound
	}
	return op, nil
}

func (r *Registry) Execute(name string, operands []float64) (float64, error) {
	op, ok := r.ops[name]
	if !ok {
		return 0, ErrOpNotFound
	}
	return op.Execute(operands)
}

func (r *Registry) List() []Operation {
	l := make([]Operation, 0, len(r.ops))
	for _, op := range r.ops {
		l = append(l, op)
	}
	return l
}
