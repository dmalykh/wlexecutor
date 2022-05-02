package brainfuck

import (
	"context"
	"github.com/dmalykh/wlinterpreter"
	"github.com/dmalykh/wlinterpreter/dialect/brainfuck"
	"github.com/dmalykh/wlinterpreter/interpreter"
)

type Brainfuck[S wlinterpreter.CellSize] struct {
	interpreter *brainfuck.Brainfuck[S]
}

func (b *Brainfuck[S]) Run(ctx context.Context, bytes []byte) error {
	return b.interpreter.Run(bytes...)
}

func New[S wlinterpreter.CellSize](st wlinterpreter.Stack[S], store wlinterpreter.Storage, input, output chan S) (*Brainfuck[S], error) {
	// Create interpreter and stack
	var wli = interpreter.NewInterpreter(store)

	// Create brainfuck instance
	bfi, err := brainfuck.New[S](st, wli, input, output)
	if err != nil {
		return nil, err
	}

	return &Brainfuck[S]{
		interpreter: bfi,
	}, err
}
