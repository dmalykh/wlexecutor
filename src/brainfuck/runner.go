package brainfuck

import (
	"bufio"
	"context"
	"fmt"
	"github.com/dmalykh/wlinterpreter"
	"github.com/dmalykh/wlinterpreter/dialect/brainfuck"
	"github.com/dmalykh/wlinterpreter/stack"
	"github.com/dmalykh/wlinterpreter/stack/slice"
	"log"
	"os"
	"sync"
)

func Run[S wlinterpreter.CellSize](ctx context.Context, size int, storage wlinterpreter.Storage) {
	// Input and output channels
	var input = make(chan S)
	var output = make(chan S)
	defer func() {
		close(input)
		close(output)
	}()

	var st = slice.NewStack[S](size)
	bf, err := New[S](st, storage, input, output)
	if err != nil {
		panic(err)
	}

	// Create waitgroup for reading output
	var wg = new(sync.WaitGroup)
	defer wg.Wait()
	go func() {
		for symbol := range output {
			wg.Add(1)
			func(s S) {
				defer wg.Done()
				fmt.Printf("%s", string(rune(s)))
			}(symbol)
		}
	}()

	// Overwrite operator. Use ask channel for fork between program input and user input
	var ask = make(chan struct{})
	if err := bf.interpreter.Interpreter().RegisterOperator(',', func(i wlinterpreter.Interpreter) error {
		if !i.InternalStorage().Empty(brainfuck.IGNORE) {
			return nil
		}
		go func() {
			ask <- struct{}{}
		}()
		var val = <-input
		if err := stack.SetValue[S](st, val); err != nil {
			return fmt.Errorf(`cann't set value on %d: %w`, stack.GetPosition(st), err)
		}
		return nil
	}); err != nil {
		log.Fatalln(err)
	}

	// If input chan waiting for data, and user inputs only one char, put it to the input chan. Otherwise, run as program.
	var received = make(chan []byte, 1)
	defer close(received)
	go func() {
		for b := range received {
			select {
			case <-ask:
				input <- S(b[0])
			default:
				go func(b []byte) {
					if err := bf.Run(ctx, b); err != nil {
						log.Fatalln(err.Error())
					}
				}(b)
			}
		}
	}()

	// Scan input line by line. If user input received, send it to program.
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return
		case received <- scanner.Bytes():
			break
		}
	}

}
