package main

import (
	"bufio"
	"log"
)

func ExamplePart1() {
	f, err := fs.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	if err := Part1(bufio.NewScanner(f), true); err != nil {
		panic(err)
	}
	// Output:
	// 1
	// 2
	// 0
	// 0
	// 5
	// sum: 8
}
