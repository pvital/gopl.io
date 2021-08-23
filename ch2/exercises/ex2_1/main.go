// Copyright © 2021 Paulo Vital
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 42.

// Exercise 2.1 from section 2.6 - packages and files.
// Adapt tempconv to process temperatures in the Kelvin scale.

package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/pvital/gopl.io/ch2/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)

		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
			os.Exit(1)
		}

		c := tempconv.Celsius(t)
		f := tempconv.Fahrenheit(t)
		k := tempconv.Kelvin(t)

		fmt.Printf("%s = %s = %s \n", c, tempconv.CToF(c), tempconv.CToK(c))
		fmt.Printf("%s = %s = %s \n", f, tempconv.FToC(f), tempconv.FToK(f))
		fmt.Printf("%s = %s = %s \n", k, tempconv.KToC(k), tempconv.KToF(k))
	}
}
