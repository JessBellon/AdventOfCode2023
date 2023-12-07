package main

import (
	"bufio"
	"unicode"
)

type Grid struct {
	g [][]rune
}

func (g Grid) LenRow() int {
	return len(g.g)
}

func (g Grid) LenCol() int {
	return len(g.g[0])
}

func (g Grid) Valid(c, r int) bool {
	return (r >= 0 && r < g.LenRow()) && (c >= 0 && c < g.LenCol())
}

func (g Grid) isNumber(r, c int) bool {
	return unicode.IsDigit(g.g[r][c])
}

func (g *Grid) markUsed(r, c int) {
	g.g[r][c] = 'x'
}

func toDigit(x rune) int {
	return int(x - '0')
}

func (g *Grid) Digit(r, c int) int {
	d := toDigit(g.g[r][c])
	g.markUsed(r, c)
	return d
}

func (g *Grid) getNumber(r, c int) int {
	if !g.Valid(r, c) || !unicode.IsNumber(g.g[r][c]) {
		return 0
	}
	var ret int
	stack := NewStack()
	for i := 0; i < g.LenCol(); i++ {
		if v := c - i; v >= 0 {
			// fmt.Println("i:", i)
			// fmt.Println("v:", v)
			if g.Valid(r, v) && g.isNumber(r, v) {
				d := g.Digit(r, v)
				// fmt.Println("d:", d)
				stack.Push(d)
				// fmt.Println("pushed:", d)
				// num := g.Digit(r, v)
				// fmt.Println("-", ret, num)
				// ret = (num * 10) + ret
				// fmt.Println("->", ret)
			} else {
				break
			}
		}
	}
	for {
		v, ok := stack.Pop()
		if !ok {
			break
		}
		// fmt.Println("popped:", v)
		ret = (ret * 10) + v
	}

	// lookRight
	for i := 1; i < g.LenCol(); i++ {
		if v := c + i; v <= g.LenCol() {
			if g.Valid(r, v) && g.isNumber(r, v) {
				num := g.Digit(r, v)
				ret = (ret * 10) + num
				// fmt.Println(ret)
			} else {
				break
			}
		}
	}

	return ret
}

func ParseGrid(s *bufio.Scanner) (*Grid, error) {
	var grid Grid
	for s.Scan() {
		var line []rune
		for _, x := range s.Text() {
			line = append(line, x)
		}
		grid.g = append(grid.g, line)
	}
	return &grid, s.Err()
}

func (g Grid) String() string {
	var str string
	for _, v := range g.g {
		str += (string(v) + "\n")
	}
	return str
}

func isSymbol(r rune) bool {
	return !unicode.IsDigit(r) && !unicode.IsLetter(r)
}

func Process(grid *Grid) int {
	var sum int
	for r, line := range grid.g {
		for c, v := range line {
			if v == '.' {
				continue
			}
			var localNumbers []int
			if isSymbol(v) {
				// fmt.Println(string(v), r, c)
				// up 1, left 1
				if n := grid.getNumber(r-1, c-1); n != 0 {
					// fmt.Println(n, "up 1, left 1")
					// localSum += n
					localNumbers = append(localNumbers, n)
					// runningCount++
				}
				// up 1
				if n := grid.getNumber(r-1, c); n != 0 {
					// fmt.Println(n, "up 1")
					localNumbers = append(localNumbers, n)
					// localSum += n
					// runningCount++
				}
				if n := grid.getNumber(r-1, c+1); n != 0 {
					// fmt.Println(n, "up 1, right 1")
					localNumbers = append(localNumbers, n)
					// localSum += n
					// runningCount++
				}
				if n := grid.getNumber(r, c-1); n != 0 {
					// fmt.Println(n, "left 1")
					localNumbers = append(localNumbers, n)
					// localSum += n
					// runningCount++
				}
				if n := grid.getNumber(r, c+1); n != 0 {
					// fmt.Println(n, "right 1")
					localNumbers = append(localNumbers, n)
					// localSum += n
					// runningCount++
				}
				if n := grid.getNumber(r+1, c-1); n != 0 {
					// fmt.Println(n, "down 1, left 1")
					// localSum += n
					// runningCount++
					localNumbers = append(localNumbers, n)
				}
				if n := grid.getNumber(r+1, c); n != 0 {
					// fmt.Println(n, "down 1")
					localNumbers = append(localNumbers, n)
					// localSum += n
					// runningCount++
				}
				if n := grid.getNumber(r+1, c+1); n != 0 {
					// fmt.Println(n, "down 1, right 1")
					// localSum += n
					// runningCount++
					localNumbers = append(localNumbers, n)
				}
				if len(localNumbers) == 2 {
					sum += localNumbers[0] * localNumbers[1]
				}
			}
		}
	}
	return sum
}
