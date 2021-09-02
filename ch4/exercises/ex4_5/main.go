// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 93.

// Exercise 4.5 from section 4.2 - Slices.
// Eliminates adjacent duplicates in a []string slice.
package main

import (
	"fmt"
)

func main() {
	s := []string{"bla", "ble", "bli", "bli", "blo", "blu"}
	fmt.Printf("%q -> ", s)
	s = remove_duplicates(s)
	fmt.Printf("%q\n", s)


	s = []string{"bla", "bla", "ble", "bli", "blo", "blo", "blu", "blu"}
	fmt.Printf("%q -> ", s)
	s = remove_duplicates(s)
	fmt.Printf("%q\n", s)
}

// removes adjacent duplicates from a slice
func remove_duplicates(s []string) []string {
	ret := s[:0]
	i := 0

	ret = append(ret, s[i])
	for j := 1; j < len(s); j++ {
		if s[j] != ret[i] {
			ret = append(ret, s[j])
			i++
		}
	}

	return ret
}
