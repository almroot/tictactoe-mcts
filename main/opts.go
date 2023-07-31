package main

import (
	"fmt"
	"io"
	"time"

	"github.com/jessevdk/go-flags"
)

type Options struct {
	Seed            int64         `short:"s" long:"seed" description:"The RNG seed used by MCTS"`
	Timeout         time.Duration `short:"t" long:"timeout" description:"The maximum amount of time the MCTS algorithm may take per action"`
	Parallelization int           `short:"p" long:"parallelization" description:"The amount of parallel goroutines to execute for the MCTS algorithm"`
}

func NewOptions() *Options {
	return &Options{
		Seed:            time.Now().UnixMilli(),
		Timeout:         time.Second,
		Parallelization: 4,
	}
}

func (o *Options) Parse(args []string, stderr io.Writer) (int, bool) {
	const errorFormat = "error: %v\n"
	var parser = flags.NewParser(o, flags.Default)
	_, err := parser.ParseArgs(args)
	if flags.WroteHelp(err) {
		return 0, true
	} else if err != nil {
		_, _ = fmt.Fprintf(stderr, errorFormat, err)
		return 1, true
	}
	return 0, false
}
