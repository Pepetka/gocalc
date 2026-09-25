# gocalc

A minimal command-line calculator written in Go. Executes named arithmetic
operations either directly or as reverse Polish notation (RPN) expressions.
Can be used as a command-line tool or as a library.

## Features

- Seven builtin operations: `add`, `mul`, `sub`, `div`, `pow`, `neg`, `sqrt`
- Direct mode: run a single operation with its operands
- RPN mode: evaluate a whole expression in one call
- Extensible: register your own operations via the `calc` library API
- Generic data structures library (`containers`): stack, queue, set, map

## Installation

```sh
go install github.com/pepetka/gocalc/cmd/gocalc@latest
```

Or build from source:

```sh
git clone https://github.com/pepetka/gocalc.git
cd gocalc
go build -o gocalc ./cmd/gocalc
```

## Usage

```sh
gocalc <command> [arguments]
```

### calc

Execute a single operation (default mode — the command word is the operation
name). Every operation takes exactly as many operands as its arity:

```sh
gocalc add 1 2
gocalc div 10 2
gocalc sqrt 9
```

### rpn

Evaluate a whitespace-separated reverse Polish notation expression:

```sh
gocalc rpn "2 3 add"
gocalc rpn "10 2 div 3 mul"
```

### ops

List all registered operations with their arity:

```sh
gocalc ops
```

## Architecture

```text
cmd/gocalc      CLI entry point: argument parsing and mode dispatch
internal/cmd    command handlers for the calc, rpn and ops modes
calc            operation registry and builtin operations (the public API)
rpn             reverse Polish notation evaluator
containers      generic data structures: stack, queue, set, map
```

Dependency graph: `cmd/gocalc` → `internal/cmd` → `calc`, `rpn`;
`rpn` → `calc`, `containers`. The `containers` package has no internal
dependencies and can be used on its own.

## Design decision: fixed arity

All operations have a fixed arity, exposed via `Operation.Arity()`, and the
operand count is validated against it before execution — by the RPN evaluator
when popping from the stack and by `Registry.Execute` for direct calls
(`ErrInvalidOperandsNum` on mismatch). This is deliberate: in RPN a variadic
operation would make evaluation ambiguous — when the evaluator pops operands
from the stack, it cannot know how many of them a variadic `add` should
consume.

Variadic behavior is instead emulated by chaining binary operations:

```sh
# instead of "add 1 2 3":
gocalc rpn "1 2 add 3 add"
```

## Escape analysis

Inspect with:

```sh
go build -gcflags="-m" ./...
```

Notable results:

- Operation values escape to the heap at registration time: each builtin
  constructor boxes its concrete struct into the `Operation` interface
  (`op escapes to heap` in `calc/builtin.go`).
- The operands slice allocated per operation in `rpn.Eval` escapes
  (`make([]float64, arity) escapes to heap`), as does the stack's backing
  array grown by `append` in `Stack.Push`.
- Result values printed by the CLI escape due to `float64` → `any` boxing
  in the `fmt.Println` calls in `cmd/gocalc/main.go`.

These allocations are inherent to the interface-based design and are
negligible for a CLI calculator; no hot loop depends on them.

## Adding an operation

Operations implement the `calc.Operation` interface: `Name`, `Arity` and
`Execute`. To add a builtin-style operation, e.g. `mod`:

1. In `calc/builtin.go`, define the operation type embedding `baseOperation`:

   ```go
   type modOp struct {
   	baseOperation
   }

   func newMod() Operation {
   	op := modOp{
   		name:  "mod",
   		arity: 2,
   	}
   	return op
   }

   func (op modOp) Execute(operands []float64) (float64, error) {
   	return math.Mod(operands[0], operands[1]), nil
   }
   ```

   Operand count is validated by `Registry.Execute` against `Arity()`, so
   `Execute` only needs to check operation-specific constraints (if any —
   e.g. `div` rejects a zero divisor, while `mod` accepts it and returns
   NaN via `math.Mod`).

2. Register it in `registerBuiltin()` in `calc/registry.go`:

   ```go
   _ = r.Register(newMod())
   ```

3. Use it from the CLI:

   ```sh
   gocalc mod 10 3
   gocalc rpn "10 3 mod"
   ```

Alternatively, from library code, implement `calc.Operation` on your own
type and pass it to `registry.Register` — no changes to the `calc` package
are needed.

## Library usage

```sh
go get github.com/pepetka/gocalc
```

```go
package main

import (
	"fmt"
	"log"

	"github.com/pepetka/gocalc/calc"
	"github.com/pepetka/gocalc/rpn"
)

func main() {
	registry := calc.New()

	res, err := registry.Execute("mul", []float64{4, 5})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res) // 20

	res, err = rpn.Eval(registry, "2 3 add")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res) // 5
}
```

## Known limitations

- All operations are fixed-arity; variadic calls must be chained (see above).
- Only `float64` operands; no integers, complex numbers or arbitrary precision.
- RPN expressions must be whitespace-separated; `2 3 add` works, `2,3,add` does not.
- Exactly one expression per `rpn` invocation.
- No infix notation, parentheses, variables or operator precedence.

## License

[MIT](LICENSE)
