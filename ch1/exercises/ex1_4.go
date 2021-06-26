// Copyright © 2021 Paulo Vital
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 13.

// Exercises for section 1.3 - findind duplicate lines.

// How to run:
// echo "ha\nha\nhe\nhe" | go run ex1_4.go
// go run ex1_4.go ../exercises/ex1_4a.txt ../exercises/ex1_4b.txt

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := make(map[string] []string)
	input_files := os.Args[1:]
	if len(input_files) == 0 {
		countLines(os.Stdin, counts, files, "STDIN")
	} else {
		for _, arg := range input_files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, files, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, files[line])
		}
	}
}

func countLines(
	f *os.File, 
	counts map[string]int, 
	files map[string] []string, 
	input_name string,
) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		files[input.Text()] = append(files[input.Text()], input_name)
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
