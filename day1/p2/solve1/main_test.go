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
	// 29
	// 83
	// 13
	// 24
	// 42
	// 14
	// 76
	// sum: 281
}
