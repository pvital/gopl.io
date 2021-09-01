// Copyright © 2021 Paulo Vital
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 84.

// Exercise 4.2 from section 4.1 - Arrays.
// Prints the SHA256 hash of its std input by default, but supports cmd-line
// flag to print the SHA384 or SHA512 hashes.

package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func gen_sha256(s string) [sha256.Size]byte {
	return sha256.Sum256([]byte(s)) 
}

func gen_sha384(s string) [sha512.Size384]byte {
	return sha512.Sum384([]byte(s))
}

func gen_sha512(s string) [sha512.Size]byte {
	return sha512.Sum512([]byte(s))
}

func main() {
	// Read from std input
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter the text to encrypt: ")
	text, _ := reader.ReadString('\n')

	for i := 0; i < len(os.Args); i++ {
		switch arg := os.Args[i]; arg {
		case "--sha384":
			fmt.Printf("SHA384 = %x\n", gen_sha384(text))
		case "--sha512":
			fmt.Printf("SHA512 = %x\n", gen_sha512(text))
		default:
			fmt.Printf("SHA256 = %x\n", gen_sha256(text))
		}
	}
}