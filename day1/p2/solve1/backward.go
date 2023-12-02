package main

import (
	"maps"
	"unicode"
)

var validNumberLastRune = map[rune]NumberSet{
	'e': makeNumberSet(one, three, five, nine),
	'o': makeNumberSet(two),
	'r': makeNumberSet(four),
	'x': makeNumberSet(six),
	'n': makeNumberSet(seven),
	't': makeNumberSet(eight),
}

func validNumbersEndingWithRune(r rune) (NumberSet, bool) {
	valid, ok := validNumberLastRune[r]
	m := maps.Clone(valid)
	return m, ok
}

func (n Number) hasSubstringBackward(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(n) < len(s) {
		return false
	}
	if len(n) == len(s) {
		return n.isMatch(s)
	}
	return string(n)[len(n)-len(s):] == s
}

func FindLast(input string) int {
	for end := len(input) - 1; end >= 0; end-- {
		r := rune(input[end])
		if unicode.IsDigit(r) {
			return runeToInt(r)
		}
		if !unicode.IsLetter(r) {
			continue
		}
		poss, ok := validNumbersEndingWithRune(r)
		if !ok {
			continue
		}
		for current := end - 1; current >= 0; current-- {

			substr := input[current : end+1]
			for n := range poss {
				if !n.hasSubstringBackward(substr) {
					delete(poss, n)
					if len(poss) == 0 {
						break
					}
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
