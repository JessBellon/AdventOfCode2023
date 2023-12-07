package main

import (
	"bufio"
	"fmt"
	"testing"
)

func TestParseGrid(t *testing.T) {
	f, err := fs.Open("example.txt")
	if err != nil {
		t.Fatal(err)
	}
	grid, err := ParseGrid(bufio.NewScanner(f))
	if err != nil {
		t.Fatal(err)
	}

	// for y := range grid {
	// 	for x := range grid[y] {
	// 		// fmt.Println(string(grid[x][y]))
	// 	}
	// }

	fmt.Println(grid.String())
	fmt.Println()
	fmt.Println()
	fmt.Println()
	sum := Process(grid)

	fmt.Println()
	fmt.Println()

	fmt.Println(grid.String())
	fmt.Println(sum)

	// num := grid.getNumber(0, 2)
	// fmt.Println(num)
	// fmt.Println(grid.String())
	// num = grid.getNumber(0, 1)
	// fmt.Println()
	// fmt.Println()
	// fmt.Println()
	// fmt.Println()
	// fmt.Println(num)
	// fmt.Println(grid.String())
	// _ = grid.getNumber(0, 2)
	// _ = grid.getNumber(0, 3)
	// _ = grid.getNumber(0, 4)
	// _ = grid.getNumber(0, 5)
}

// func TestParseLine(t *testing.T) {
// 	type testcases struct {
// 		in  string
// 		num []Number
// 		sym []Symbol
// 	}
//
// 	tc := []testcases{
// 		{
// 			in: "467..114..",
// 			num: []Number{
// 				{Data: "467", YPos: []int{0, 1, 2}},
// 				{Data: "114", YPos: []int{5, 6, 7}},
// 			},
// 		},
// 		{
// 			in:  ".../..%..=.",
// 			sym: []Symbol{3, 6, 9},
// 		},
// 		{
// 			in:  "...*......",
// 			sym: []Symbol{3},
// 		},
// 		{
// 			in: "..35..633.",
// 			num: []Number{{
// 				Data: "35", YPos: []int{2, 3},
// 			}, {
// 				Data: "633", YPos: []int{6, 7, 8},
// 			}},
// 		},
//
// 		{
// 			in:  "......#...",
// 			sym: []Symbol{6},
// 		},
// 		{
// 			in:  "617*......",
// 			sym: []Symbol{3},
// 			num: []Number{{Data: "617", YPos: []int{0, 1, 2}}},
// 		},
// 		{
// 			in:  ".....+.58.",
// 			sym: []Symbol{5},
// 			num: []Number{{Data: "58", YPos: []int{7, 8}}},
// 		},
// 		{
// 			in:  "..592.....",
// 			num: []Number{{Data: "592", YPos: []int{2, 3, 4}}},
// 		},
// 		{
// 			in:  "......755.",
// 			num: []Number{{Data: "755", YPos: []int{6, 7, 8}}},
// 		},
// 		{
// 			in:  "...$.*....",
// 			sym: []Symbol{3, 5},
// 		},
// 		{
// 			in: ".664.598..",
// 			num: []Number{
// 				{Data: "664", YPos: []int{1, 2, 3}},
// 				{Data: "598", YPos: []int{5, 6, 7}},
// 			},
// 		},
// 	}
//
// 	for _, tt := range tc {
// 		t.Run(string(tt.in), func(t *testing.T) {
// 			gotNum, gotSym := ParseLine(tt.in)
// 			if diff := cmp.Diff(tt.num, gotNum); diff != "" {
// 				t.Error(diff)
// 			}
// 			if diff := cmp.Diff(tt.sym, gotSym); diff != "" {
// 				t.Error(diff)
// 			}
// 		})
// 	}
// }
