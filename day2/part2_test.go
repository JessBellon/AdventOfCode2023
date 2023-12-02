package main

import (
	"bufio"
	"log"
)

func ExamplePart2() {
	f, err := fs.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	if err := Part2(bufio.NewScanner(f), true); err != nil {
		panic(err)
	}
	// Output:
	// power: 48
	// power: 12
	// power: 1560
	// power: 630
	// power: 36
	// sum: 2286
}
