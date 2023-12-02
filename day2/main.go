package main

import (
	"bufio"
	"embed"
	"log"
)

//go:embed *.txt
var fs embed.FS

func main() {
	f, err := fs.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	// if err := Part1(scanner, false); err != nil {
	// 	log.Fatal(err)
	// }
	if err := Part2(scanner, false); err != nil {
		log.Fatal(err)
	}
}
