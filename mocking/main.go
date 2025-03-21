package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

var (
	finalWord     = "Go!"
	coundownStart = 3
)

func Countdown(out io.Writer) {
	for i := coundownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
