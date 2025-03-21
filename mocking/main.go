package main

import (
	"fmt"
	"io"
	"os"
)

var (
	finalWord     = "Go!"
	coundownStart = 3
)

func Countdown(out io.Writer) {
	for i := coundownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
