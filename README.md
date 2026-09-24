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

Execute a single operation (default mode — the command word is the operation name):

```sh
gocalc add 1 2 3
gocalc div 10 2
gocalc sqrt 9
```

### rpn

Evaluate a reverse Polish notation expression:

```sh
gocalc rpn "2 3 add"
gocalc rpn "10 2 div 3 mul"
```

### ops

List all registered operations with their arity:

```sh
gocalc ops
```

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

Custom operations implement the `calc.Operation` interface (`Name`, `Arity`,
`Execute`) and are added with `registry.Register`.

## License

[MIT](LICENSE)
