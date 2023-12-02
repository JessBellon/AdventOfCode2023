package main

import (
	"bufio"
	"fmt"
	"unicode"
)

type Grid [][]rune

func (g Grid) LenRow() int {
	return len(g[0])
}

func (g Grid) LenCol() int {
	return len(g)
}

func (g Grid) Valid(r, c int) bool {
	return (r >= 0 && r <= g.LenRow()) && (c >= 0 && c <= g.LenCol())
}

func (g Grid) isNumber(r, c int) bool {
	return unicode.IsDigit(g[r][c])
}

func (g Grid) getNumber(r, c int) int {
	this := toDigit(g[r][c])
	fmt.Println("this:", this)
	// check right
	for i := 1; i < g.Len(); i++ {
		rd := r + i
		fmt.Println(rd, c, "**", string(g[rd][c]))
		if !g.Valid(rd, c) || !g.isNumber(rd, c) {
			break
		}

		// r := toDigit(g[xd][y])
		// this = (this * 10) + r
		// fmt.Println("right:", this)
	}

	// check left
	// for i := 1; i >= 0 && g.isNumber(x-i, y); i++ {
	// 	l := toDigit(g[x-i][y]) * 10
	// 	this = l + this
	// 	fmt.Println("left:", this)
	// }
	return this
}

func toDigit(x rune) int {
	return int(x - '0')
}

func ParseGrid(s *bufio.Scanner) (Grid, error) {
	var grid Grid
	for s.Scan() {
		var line []rune
		for _, x := range s.Text() {
			line = append(line, x)
		}
		fmt.Println(string(line))
		grid = append(grid, line)
	}
	return grid, s.Err()
}

// func Process(grid Grid) (int, error) {
// 	for y, line := range grid {
// 		for x, v := range line {
// 			if v == '.' {
// 				continue
// 			}
// 			// a b c
// 			// d e f
// 			// g h i
// 			// xleft := x - 1
// 			// xright := x + 1
// 			// yup := y - 1
// 			// ydown := y + 1
// 			//
// 			// if xleft >= 0 {
// 			// }
// 			if !unicode.IsDigit(v) {
// 				if grid.Valid(x-1, y) {
// 				}
// 				// check x
// 				// up left
// 				// up
// 				// up right
// 				// left
// 				// right
// 				// down left
// 				// down
// 				// down right
// 			}
// 		}
// 	}
// }
