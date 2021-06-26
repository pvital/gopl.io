// Copyright © 2021 Paulo Vital
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Exercise 1.1 from section 1.2 - command-line arguments.

package main

import (
	"fmt"
	"os"
	"strings"
)


func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Please, provide some extra arguments")
	} else {
		// Modified echo3 program to also print the name of the command that
		// invoked it.
		fmt.Println(strings.Join(os.Args, " "))
	}
}