package main

import (
	"maps"
	"unicode"
)

func FindFirst(input string) int {
	for start := 0; start < len(input); start++ {
		r := rune(input[start])
		if unicode.IsDigit(r) {
			return runeToInt(r)
		}
		if !unicode.IsLetter(r) {
			continue
		}
		poss, ok := validNumbersStartingWithRune(r)
		if !ok {
			continue
		}
		for current := start + 1; current < len(input); current++ {
			substr := input[start:current]

			for n := range poss {
				if !n.hasSubstringForward(substr) {
					delete(poss, n)
					continue
				}
				if n.isMatch(substr) {
					return n.ToInt()
				}
			}
		}
	}
	return 0
}

func (n Number) hasSubstringForward(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(n) < len(s) {
		return false
	}
	if len(n) == len(s) {
		return n.isMatch(s)
	}

	return string(n)[0:len(s)] == s
}

func validNumbersStartingWithRune(r rune) (NumberSet, bool) {
	valid, ok := validNumberFirstRune[r]
	m := maps.Clone(valid)
	return m, ok
}

var validNumberFirstRune = map[rune]NumberSet{
	'e': makeNumberSet(eight),
	'f': makeNumberSet(five, four),
	'n': makeNumberSet(nine),
	'o': makeNumberSet(one),
	's': makeNumberSet(seven, six),
	't': makeNumberSet(three, two),
}
