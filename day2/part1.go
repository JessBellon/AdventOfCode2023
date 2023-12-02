package main

import (
	"bufio"
	"fmt"
	"strings"
)

func Part1(scanner *bufio.Scanner, debug bool) error {
	var sum int
	for scanner.Scan() {
		str := scanner.Text()
		i, err := ParseLine(str)
		if err != nil {
			return err
		}
		if debug {
			fmt.Println(i)
		}
		sum += i
	}
	fmt.Println("sum:", sum)
	return scanner.Err()
}

func ParseLine(st string) (int, error) {
	x := strings.Split(st, ":")

	for _, round := range strings.Split(x[1], ";") {
		for _, set := range strings.Split(round, ",") {
			s := strings.TrimSpace(set)
			var (
				count int
				color string
			)
			if _, err := fmt.Sscanf(s, "%d %s", &count, &color); err != nil {
				return 0, err
			}
			if !valid(color, count) {
				return 0, nil
			}
		}
	}
	return parseGameNumber(x[0])
}

func valid(color string, count int) bool {
	switch color {
	case blue:
		return count <= maxBlue
	case green:
		return count <= maxGreen
	case red:
		return count <= maxRed
	}
	return false
}

func parseGameNumber(str string) (int, error) {
	var i int
	_, err := fmt.Sscanf(str, "Game %d", &i)
	return i, err
}

const (
	blue     = "blue"
	red      = "red"
	green    = "green"
	maxBlue  = 14
	maxRed   = 12
	maxGreen = 13
)
