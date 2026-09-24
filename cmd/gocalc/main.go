// Command gocalc is a CLI calculator that executes arithmetic operations
// either directly or as reverse Polish notation expressions.
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/pepetka/gocalc/calc"
	"github.com/pepetka/gocalc/internal/cmd"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		return errors.New("not enough arguments")
	}
	command := args[0]
	mode := cmd.ParseMode(command)

	registry := calc.New()

	switch mode {
	case cmd.RpnMode:
		args = args[1:]
		res, err := cmd.RpnHandler(registry, args)
		if err != nil {
			return err
		}
		fmt.Println(res)
	case cmd.ListMode:
		list := cmd.ListHandler(registry)
		for _, op := range list {
			n := op.Name()
			s := op.Arity()
			fmt.Printf("Name: %s, Arity: %d\n", n, s)
		}
	case cmd.CalcMode:
		res, err := cmd.CalcHandler(registry, args)
		if err != nil {
			return err
		}
		fmt.Println(res)
	}
	return nil
}
