package main

import (
	"bufio"
	"fmt"
	"strings"
)

func Part2(scanner *bufio.Scanner, debug bool) error {
	var sum int
	for scanner.Scan() {
		str := scanner.Text()
		i, err := ParseLinePart2(str)
		if err != nil {
			return err
		}
		if debug {
			fmt.Println("power:", i)
		}
		sum += i
	}
	fmt.Println("sum:", sum)
	return scanner.Err()
}

func ParseLinePart2(st string) (int, error) {
	x := strings.Split(st, ":")
	var (
		maxRed   = 0
		maxGreen = 0
		maxBlue  = 0
	)
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
			switch color {
			case blue:
				if count > maxBlue {
					maxBlue = count
				}
			case green:
				if count > maxGreen {
					maxGreen = count
				}
			case red:
				if count > maxRed {
					maxRed = count
				}
			}
		}
	}
	return maxBlue * maxRed * maxGreen, nil
}
