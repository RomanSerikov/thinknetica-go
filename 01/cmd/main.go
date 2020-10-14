package main

import (
	"01/pkg/fib"
	"flag"
	"fmt"
	"log"
)

const maxFibNumber = 20

func main() {
	number := flag.Int("n", maxFibNumber, "fibonacci number")
	flag.Parse()

	if *number > maxFibNumber {
		log.Fatal("exceeded max fibonacci number")
	}

	fmt.Printf("Fibonacci number #%d is %d\n", *number, fib.Num(*number))
}
