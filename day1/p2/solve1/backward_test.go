package main

import "testing"

func TestFindLast(t *testing.T) {
	tc := []struct {
		input string
		last  int
	}{
		{input: "eightwothree", last: 3},
		{input: "abcone2threexyz", last: 3},
		{input: "4nineeightseven2", last: 2},
		{input: "xtwone3four", last: 4},
		{input: "zoneight234", last: 4},
		{input: "7pqrstsixteen", last: 6},
		{input: "7pqrstsixtthreeen", last: 3},
		{input: "abcone2threexyz9", last: 9},
		{input: "abcone2threex8yz", last: 8},
		{input: "abcone2threexyz!", last: 3},
		{input: "", last: 0},
	}
	for _, tt := range tc {
		t.Run(tt.input, func(t *testing.T) {
			got := FindLast(tt.input)
			if got != tt.last {
				t.Error(got, tt.last)
			}
		})
	}
}

func TestHasSubstringBackword(t *testing.T) {
	tc := []struct {
		n     Number
		input string
		exp   bool
	}{
		{n: three, input: "txhree", exp: false},
		{n: three, input: "three", exp: true},
		{n: three, input: "hree", exp: true},
		{n: three, input: "ree", exp: true},
		{n: three, input: "ee", exp: true},
		{n: three, input: "e", exp: true},
		{n: three, input: "", exp: true},
		{n: three, input: "r", exp: false},
	}

	for _, tt := range tc {
		t.Run(tt.input, func(t *testing.T) {
			got := tt.n.hasSubstringBackward(tt.input)
			if got != tt.exp {
				t.Errorf("[%s - %s] expected %t got %t", tt.n, tt.input, tt.exp, got)
			}
		})
	}
}
