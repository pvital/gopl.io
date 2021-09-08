// Copyright © 2021 Paulo Vital
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 93.

// Exercise 4.4 from section 4.2 - Slices.
// Rotates elements of a slice in a single pass.
package main

import (
	"fmt"
)

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(s)

	pos := 2
	fmt.Printf("rotating by %d positions...\n", pos)

	// Rotate s left by pos positions.
	rotate(s, pos)
	fmt.Println(s)
}

// rotates a slice of ints in place by its given position.
func rotate(s []int, pos int) {
	for i, j := 0, pos; i < j && j < len(s); i, j = i+1, j+1 {
		s[i], s[j] = s[j], s[i]
	}
}
