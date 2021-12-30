package main

import (
	"fmt"
	"strings"
)

type passwd string

func (p passwd) hasStraight() bool {

	for _, straight := range []string{"abc", "bcd", "cde", "def", "efg", "fgh", "ghi", "hij", "ijk", "jkl", "klm", "lmn", "mno", "nop", "opq", "pqr", "qrs", "rst", "stu", "tuv", "uvw", "vwx", "wxy", "xyz"} {
		if strings.Contains(string(p), straight) {
			return true
		}
	}

	return false
}

func (p passwd) hasDoubleChars() bool {
	for _, pair := range []string{"aa", "bb", "cc", "dd", "ee", "ff", "gg", "hh", "ii", "jj", "kk", "ll", "mm", "nn", "oo", "pp", "qq", "rr", "ss", "tt", "uu", "vv", "ww", "xx", "yy", "zz"} {
		for _, other := range []string{"aa", "bb", "cc", "dd", "ee", "ff", "gg", "hh", "ii", "jj", "kk", "ll", "mm", "nn", "oo", "pp", "qq", "rr", "ss", "tt", "uu", "vv", "ww", "xx", "yy", "zz"} {
			if pair != other {
				if strings.Contains(string(p), pair) && strings.Contains(string(p), other) {
					return true
				}
			}
		}
	}

	return false
}

func nextLetter(c byte, carry bool) (byte, bool) {
	a := c

	continueCarry := false

	switch a {
	case 'a':
		a = 'b'
	case 'b':
		a = 'c'
	case 'c':
		a = 'd'
	case 'd':
		a = 'e'
	case 'e':
		a = 'f'
	case 'f':
		a = 'g'
	case 'g':
		a = 'h'
	case 'h':
		a = 'i'
	case 'i':
		a = 'j'
	case 'j':
		a = 'k'
	case 'k':
		a = 'l'
	case 'l':
		a = 'm'
	case 'm':
		a = 'n'
	case 'n':
		a = 'o'
	case 'o':
		a = 'p'
	case 'p':
		a = 'q'
	case 'q':
		a = 'r'
	case 'r':
		a = 's'
	case 's':
		a = 't'
	case 't':
		a = 'u'
	case 'u':
		a = 'v'
	case 'v':
		a = 'w'
	case 'w':
		a = 'x'
	case 'x':
		a = 'y'
	case 'y':
		a = 'z'
	case 'z':
		a = 'a'
		continueCarry = true
	}

	return a, continueCarry
}

func (p passwd) isValid() bool {
	if !p.hasStraight() {
		return false
	}

	if !p.hasDoubleChars() {
		return false
	}

	if strings.ContainsAny(string(p), "iol") {
		return false
	}

	return true
}

func (p passwd) next() passwd {
	np := passwd("")
	carry := false
	var n int
	for n = len(p) - 1; n >= 0; n-- {
		if carry == false && n < len(p)-1 {
			break
		}
		c, continueCarry := nextLetter(p[n], carry)
		carry = continueCarry
		np = passwd(c) + np
	}
	for n >= 0 {
		np = passwd(p[n]) + np
		n--
	}

	for !np.isValid() {
		np = np.next()
	}

	return np
}

func day11(puzzle_data []string) {
	p := passwd(puzzle_data[0])
	np := p.next()
	fmt.Printf("Day 10 (part 1): np = %s\n", np)
	np = np.next()
	fmt.Printf("Day 10 (part 2): np = %s\n", np)
}
