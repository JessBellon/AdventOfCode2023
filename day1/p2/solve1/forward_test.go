package main

import "testing"

func TestFindFirst(t *testing.T) {
	tc := []struct {
		input string
		first int
	}{
		{input: "eightwothree", first: 8},
		{input: "abcone2threexyz", first: 1},
		{input: "4nineeightseven2", first: 4},
		{input: "xtwone3four", first: 2},
		{input: "zoneight234", first: 1},
		{input: "7pqrstsixteen", first: 7},
		{input: "!7pqrstsixteen", first: 7},
		{input: "five8pbcsllrbvg787", first: 5},
	}
	for _, tt := range tc {
		t.Run(tt.input, func(t *testing.T) {
			got := FindFirst(tt.input)
			if got != tt.first {
				t.Error(got, tt.first)
			}
		})
	}
}

func TestHasSubstringForward(t *testing.T) {
	tc := []struct {
		n     Number
		input string
		exp   bool
	}{
		{n: one, input: "o", exp: true},
		{n: one, input: "on", exp: true},
		{n: one, input: "one", exp: true},
		{n: one, input: "ones", exp: false},
		{n: one, input: "", exp: true},
		{n: one, input: "x", exp: false},
	}

	for _, tt := range tc {
		t.Run(tt.input, func(t *testing.T) {
			got := tt.n.hasSubstringForward(tt.input)
			if got != tt.exp {
				t.Errorf("[%s - %s] expected %t got %t", tt.n, tt.input, tt.exp, got)
			}
		})
	}
}
