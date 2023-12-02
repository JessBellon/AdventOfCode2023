package main

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"os"
)

//go:embed *.txt
var f embed.FS

func main() {
	r, err := f.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(r)
	// var sum int
	for scanner.Scan() {
		i := ProcessLine(scanner.Text())
		fmt.Println(i)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading file:", err)
	}
	// fmt.Println(sum)
}

const runeRepresentingASCIIZero = '0'

func runeToInt(v rune) int {
	return int(v - runeRepresentingASCIIZero)
}

// eight
// five
// four
// nine
// one
// seven
// six
// three
// two

func ProcessLine(line string) int {
}

// d1, d2 := 0, 0
// for _, v := range s {
// 	if unicode.IsDigit(v) {
// 		i := runeToInt(v)
// 		if d1 == 0 {
// 			d1 = i
// 		}
// 		d2 = runeToInt(v)
// 	}
// }
// sum += (10 * d1) + d2
