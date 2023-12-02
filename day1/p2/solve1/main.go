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
	if err := Run(bufio.NewScanner(f), WithDebug()); err != nil {
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
		str := scanner.Text()
		i := ProcessLine(str)
		if cfg.WithDebug {
			fmt.Println(str, i)
		}
		sum += i
	}
	fmt.Println("sum:", sum)
	return scanner.Err()
}

func ProcessLine(line string) int {
	tensDigit := FindFirst(line)
	onesDigit := FindLast(line)
	return (10 * tensDigit) + onesDigit
}

const runeRepresentingASCIIZero = '0'

func runeToInt(v rune) int {
	return int(v - runeRepresentingASCIIZero)
}
