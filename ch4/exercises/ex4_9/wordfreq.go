// Copyright © 2021 Paulo Vital
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 99.

// Exercise 4.9 from section 4.3 - Maps.
// wordfreq reports the frequency of each word in an input text file.

package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

// get_path_separator returns the correct separator string for the current OS
func get_path_separator() string {
	var sep string
	
	switch os := runtime.GOOS; os {
	case "windows":
		sep = "\\"
	default:
		sep = "/"
	}

	return sep
}

// usage prints the usage or help message.
func usage() {
	prog := strings.Split(os.Args[0], get_path_separator())
	fmt.Printf("Usage: %s [-h] [-i file]\n", prog[len(prog)-1])
	fmt.Printf("\t -i file: input file name (default: input.txt)\n")
	fmt.Printf("\t -h:      print this message\n")
}

func main() {
	var input_file string

	// Read arguments
	if len(os.Args) <=1 {
		input_file = "input.txt"
	} else {
		switch arg := os.Args[1]; arg {
		case "-i":
			input_file = os.Args[2]
		case "-h":
			usage()
			os.Exit(0)
		default:
			fmt.Printf("ERROR - invalid argument: %s\n", arg)
			usage()
			os.Exit(1)
		}		
	}

	// Open the file
	file, err := os.Open(input_file)
    if err != nil {
        fmt.Printf("ERROR - %q\n", err)
		os.Exit(1)
    }
	defer file.Close()

	// Read the file
	input := bufio.NewScanner(file)
	input.Split(bufio.ScanWords)

	count := make(map[string]int)
	for input.Scan() {
		count[input.Text()]++
	}

	// Print results
	fmt.Printf("The file %s has the following frequence words:\n", input_file)
	for word, w_count := range count {
		fmt.Printf("\t%d\t%s\n", w_count, word)
	}

}
