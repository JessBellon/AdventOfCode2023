package main

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"unicode"
)

//go:embed *.txt
var fs embed.FS

func main() {
	f, err := fs.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	if err := Run(bufio.NewScanner(f)); err != nil {
		log.Fatal(err)
	}
}

type Option func(*config)

func WithDebug() Option {
	return func(cfg *config) {
		cfg.WithDebug = true
	}
}

type config struct {
	WithDebug bool
}

func Run(scanner *bufio.Scanner, opts ...Option) error {
	var cfg config
	for _, opt := range opts {
		opt(&cfg)
	}
	var sum int
	for scanner.Scan() {
		i := ProcessLine(scanner.Text())
		if cfg.WithDebug {
			fmt.Println(i)
		}
		sum += i
	}
	fmt.Println("sum:", sum)
	return scanner.Err()
}

func ProcessLine(line string) int {
	r := []rune(line)
	tensDigit := findFirstInt(r)
	onesDigit := findLastInt(r)
	return (10 * tensDigit) + onesDigit
}

const runeRepresentingASCIIZero = '0'

func runeToInt(v rune) int {
	return int(v - runeRepresentingASCIIZero)
}

func findFirstInt(str []rune) int {
	for i := 0; i < len(str); i++ {
		v := str[i]
		if unicode.IsDigit(v) {
			return runeToInt(v)
		}
	}
	return 0
}

func findLastInt(str []rune) int {
	for i := len(str) - 1; i >= 0; i-- {
		v := str[i]
		if unicode.IsDigit(v) {
			return runeToInt(v)
		}
	}
	return 0
}
