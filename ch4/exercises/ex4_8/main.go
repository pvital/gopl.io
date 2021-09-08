// Copyright © 2021 Paulo Vital
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 99.

// Exercise 4.8 from section 4.3 - Maps.
// Adaptation of charcount app to count letters, digits, and so on in their Unicode categories.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

// category L - Letter, Title, Lower and Upper
var cat_L = []*unicode.RangeTable {	
	unicode.Letter, unicode.L,
	unicode.Lm,
	unicode.Lo,
	unicode.Lower, unicode.Ll,
	unicode.Title, unicode.Lt,
	unicode.Upper, unicode.Lu,
}

// category M - Mark
var cat_M = []*unicode.RangeTable {
	unicode.Mark, unicode.M,
	unicode.Mc,
	unicode.Me,
	unicode.Mn,
}

// category N - Number, Digit
var cat_N = []*unicode.RangeTable {
	unicode.Number, unicode.N,
	unicode.Nl,
	unicode.No,
	unicode.Digit, unicode.Nd,
}

// category P - Punctuation
var cat_P = []*unicode.RangeTable {
	unicode.Punct, unicode.P,
	unicode.Pc,
	unicode.Pd,
	unicode.Pe,
	unicode.Pf,
	unicode.Pi,
	unicode.Po,
	unicode.Ps,
}

// category S - Symbol
var cat_S = []*unicode.RangeTable {
	unicode.Symbol, unicode.S,
	unicode.Sc,
	unicode.Sk,
	unicode.Sm,
	unicode.So,
}

// category Z - Separator, Space
var cat_Z = []*unicode.RangeTable {
	unicode.Space, unicode.Z,
	unicode.Zl,
	unicode.Zp,
	unicode.Zs,
}

func main() {

	l_counts := make(map[rune]int)	// counts of cat_L Unicode characters
	m_counts := make(map[rune]int)	// counts of cat_M Unicode characters
	n_counts := make(map[rune]int)	// counts of cat_N Unicode characters
	p_counts := make(map[rune]int)	// counts of cat_P Unicode characters
	s_counts := make(map[rune]int)	// counts of cat_S Unicode characters
	z_counts := make(map[rune]int)	// counts of cat_Z Unicode characters
	c_counts := make(map[rune]int)	// counts of cat_C (Other) Unicode characters

	invalid := 0					// count of invalid UTF-8 characters


	fmt.Println("Type as many Unicode characters you want. To finish, type ENTER and CTRL+D.")
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		if unicode.In(r, cat_L...) {
			l_counts[r]++
		} else if unicode.In(r, cat_M...) {
			m_counts[r]++
		} else if unicode.In(r, cat_N...) {
			n_counts[r]++
		} else if unicode.In(r, cat_P...) {
			p_counts[r]++
		} else if unicode.In(r, cat_S...) {
			s_counts[r]++
		} else if unicode.In(r, cat_Z...) {
			z_counts[r]++
		} else {
			c_counts[r]++
		}
	}

	// Print results
	fmt.Printf("Category\tcount\trunes\n")

	var runes []rune
	total := 0

	// cat_L
	for c, n := range l_counts {
		runes = append(runes, c)
		total += n
	}
	fmt.Printf("Letter\t\t%d\t%q\n", total, runes)

	// cat_M
	runes = runes[:0]
	total = 0
	for c, n := range m_counts {
		runes = append(runes, c)
		total += n
	}
	fmt.Printf("Mark\t\t%d\t%q\n", total, runes)

	// cat_N
	runes = runes[:0]
	total = 0
	for c, n := range n_counts {
		runes = append(runes, c)
		total += n
	}
	fmt.Printf("Number\t\t%d\t%q\n", total, runes)

	// cat_P
	runes = runes[:0]
	total = 0
	for c, n := range p_counts {
		runes = append(runes, c)
		total += n
	}
	fmt.Printf("Punctuation\t%d\t%q\n", total, runes)

	// cat_S
	runes = runes[:0]
	total = 0
	for c, n := range s_counts {
		runes = append(runes, c)
		total += n
	}
	fmt.Printf("Symbol\t\t%d\t%q\n", total, runes)

	// cat_Z
	runes = runes[:0]
	total = 0
	for c, n := range z_counts {
		runes = append(runes, c)
		total += n
	}
	fmt.Printf("Separator\t%d\t%q\n", total, runes)

	// cat_C
	runes = runes[:0]
	total = 0
	for c, n := range c_counts {
		runes = append(runes, c)
		total += n
	}
	fmt.Printf("Other\t\t%d\t%q\n", total, runes)

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
