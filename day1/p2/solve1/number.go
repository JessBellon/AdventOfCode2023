package main

type Number string

const (
	one   = "one"
	two   = "two"
	three = "three"
	four  = "four"
	five  = "five"
	six   = "six"
	seven = "seven"
	eight = "eight"
	nine  = "nine"
)

func (n Number) ToInt() int {
	return numberToInt[n]
}

func (n Number) isMatch(s string) bool {
	return string(n) == s
}

type NumberSet map[Number]struct{}

func makeNumberSet(numbs ...Number) NumberSet {
	ns := make(NumberSet)
	for _, n := range numbs {
		ns[n] = struct{}{}
	}
	return ns
}

var numberToInt = map[Number]int{
	one:   1,
	two:   2,
	three: 3,
	four:  4,
	five:  5,
	six:   6,
	seven: 7,
	eight: 8,
	nine:  9,
}
