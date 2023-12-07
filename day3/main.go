package main

import (
	"bufio"
	"embed"
	"fmt"
	"log"
)

//go:embed *.txt
var fs embed.FS

func main() {
	f, err := fs.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	grid, err := ParseGrid(bufio.NewScanner(f))
	if err != nil {
		log.Fatal(err)
	}
	sum := Process(grid)
	fmt.Println(grid.String())
	fmt.Println(sum)
}
