// Copyright © 2021 Paulo Vital
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 93.

// Exercise 4.3 from section 4.2 - Slices.
// Rewrite reverse() to use an array pointer instead of a slice.

package main

import (
	"fmt"
)

func main() {
	a := [6]int{0, 1, 2, 3, 4, 5}
	fmt.Println(a)
	fmt.Println("reversing...")
	reverse(&a)
	fmt.Println(a)
}

// reverses a slice of ints in place.
func reverse(s *[6]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
