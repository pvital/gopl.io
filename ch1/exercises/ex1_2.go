// Copyright © 2021 Paulo Vital
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Exercise 1.2 from section 1.2 - command-line arguments.

package main

import (
	"fmt"
	"os"
)


func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Please, provide some extra arguments")
	} else {
		// Modified echo1 program to print the index and value of each 
		// argument, one per line.
		for i:= 1; i < len(os.Args); i++ {
			fmt.Println(i, os.Args[i])
		}
	}
}