package main

import (
	"embed"
)

//go:embed *.txt
var fs embed.FS

func main() {
	// f, err := fs.Open("input.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if err := ParseGrid(bufio.NewScanner(f)); err != nil {
	// 	log.Fatal(err)
	// }
}
