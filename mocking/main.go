package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

var (
	finalWord     = "Go!"
	coundownStart = 3
)

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := coundownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
