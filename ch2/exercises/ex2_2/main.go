// Copyright © 2021 Paulo Vital
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 44.

// Exercise 2.2 from section 2.6.1 - imports.
// General purpose unit-conversation reading numbers and convert them to temperature, lenght, weight

package main

import (
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/exercises/ex2_2/converter"
	"gopl.io/ch2/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)

		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("++++++++++++++++++++++++++++++++++++++\n")

		c := tempconv.Celsius(t)
		f := tempconv.Fahrenheit(t)

		fmt.Printf("+ Temperature converter:\n")
		fmt.Printf("+ %s = %s\n", c, tempconv.CToF(c))
		fmt.Printf("+ %s = %s\n", f, tempconv.FToC(f))

		feet := converter.Foot(t)
		meters := converter.Meter(t)

		fmt.Printf("+ ====================================\n")
		fmt.Printf("+ Lenght converter:\n")
		fmt.Printf("+ %s = %s\n", meters, converter.M2F(meters))
		fmt.Printf("+ %s = %s\n", feet, converter.F2M(feet))

		pounds := converter.Pound(t)
		grams := converter.Gram(t)

		fmt.Printf("+ ====================================\n")
		fmt.Printf("+ Weight converter:\n")
		fmt.Printf("+ %s = %s\n", grams, converter.G2P(grams))
		fmt.Printf("+ %s = %s\n", pounds, converter.P2G(pounds))

	}
	fmt.Printf("++++++++++++++++++++++++++++++++++++++\n")
}
