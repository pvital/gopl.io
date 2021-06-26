// Copyright © 2021 Paulo Vital
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Exercises for section 1.2 - command-line arguments.

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)


func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}


func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}


func echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}


func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Please, provide some extra arguments")
	} else {
		// Measure the diff in running time among echo1, echo2 and echo3
		// programs.
		start := time.Now()
		echo1()
		fmt.Printf("Running time for echo1(): %fs\n", time.Since(start).Seconds())

		start = time.Now()
		echo2()
		fmt.Printf("Running time for echo2(): %fs\n", time.Since(start).Seconds())

		start = time.Now()
		echo3()
		fmt.Printf("Running time for echo3(): %fs\n", time.Since(start).Seconds())
	}
}