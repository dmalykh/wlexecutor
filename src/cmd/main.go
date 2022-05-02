package main

import (
	"context"
	"flag"
	"github.com/dmalykh/wlexecutor/brainfuck"
	"github.com/dmalykh/wlinterpreter/storage/list"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	var cellsize, stacksize int
	var help bool
	flag.IntVar(&cellsize, `cell`, 0, `Size of cell: 8, 32 or 64`)
	flag.IntVar(&stacksize, `stack`, 30000, `Size of stack`)
	flag.BoolVar(&help, `help`, false, `Displays this help`)
	flag.Parse()
	if help || flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	// Graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan os.Signal)
	go func() {
		<-done
		cancel()
	}()
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	//Recover
	defer func() {
		if r := recover(); r != nil {
			cancel()
			os.Exit(1)
		}
	}()

	// Get store for internal interpreter
	var store = list.New()

	switch cellsize {
	case 8:
		brainfuck.Run[int8](ctx, cellsize, store)
	case 32:
		brainfuck.Run[int32](ctx, cellsize, store)
	case 64:
		brainfuck.Run[int64](ctx, cellsize, store)
	default:
		log.Fatalf(`Not allowed cellsize, %d`, cellsize)
	}

}
