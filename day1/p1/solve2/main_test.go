package main

import (
	"bufio"
	"log"
)

func ExampleRun() {
	f, err := fs.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	if err := Run(bufio.NewScanner(f), WithDebug()); err != nil {
		panic(err)
	}
	// Output:
	// 12
	// 38
	// 15
	// 77
	// sum: 142
}
