// Package cnt writes a countdown to stdout, then optionally sends a desktop notification when complete.
package main

import (
	"flag"
	"fmt"
	"github.com/gen2brain/beeep"
	"log"
	"os"
	"time"
)

var dArg int
var uArg time.Duration
var nArg bool
var qArg bool

func main() {
	flag.IntVar(&dArg,"duration", 5, "countdown duration")
	flag.DurationVar(&uArg, "unit", time.Second, "unit")
	flag.BoolVar(&nArg, "notify", false, "notify after cnt completes")
	flag.BoolVar(&qArg, "quiet", false, "don't output to console")
	flag.Parse()

	dur := time.Duration(dArg) * uArg
	cntdn(dur, uArg, nArg)
}

func cntdn(d time.Duration, u time.Duration, n bool) {
	if d > 0 {
		quietPrint(d)
		time.Sleep(u)
		cntdn(d-u, u, n)
	} else {
    fmt.Fprintln(os.Stdout, "cnt: complete")
		if n {
			notify()
		}
	}
}

func quietPrint(d time.Duration) {
	if !qArg {
		fmt.Printf("%s\n", d)
	}
}

func notify() {
	err := beeep.Notify("CNT", "Your coundown is complete", "./img/taco.png")
	if err != nil {
		log.Fatalf("cnt failed to notify: %v", err)
	}
}
