package main

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"os"
	"unicode"
)

//go:embed *.txt
var f embed.FS

func main() {
	r, err := f.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(r)
	var sum int
	for scanner.Scan() {
		s := scanner.Text()
		d1, d2 := 0, 0
		for _, v := range s {
			if unicode.IsDigit(v) {
				i := runeToInt(v)
				if d1 == 0 {
					d1 = i
				}
				d2 = runeToInt(v)
			}
		}
		sum += (10 * d1) + d2
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading file:", err)
	}
	fmt.Println(sum)
}

func runeToInt(v rune) int {
	return int(v - '0')
}
