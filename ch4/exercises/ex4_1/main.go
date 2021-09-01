// Copyright © 2021 Paulo Vital
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 84.

// Exercise 4.1 from section 4.1 - Arrays.
// Function that counts the number of bits that are different in two SHA256
// hashes.

package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"math"
	"os"

	"gopl.io/ch2/popcount"
)

func gen_sha256(s string) [sha256.Size]byte {
	return sha256.Sum256([]byte(s)) 
}

func main() {
	// Read from std input
	reader := bufio.NewReader(os.Stdin)	
	fmt.Printf("Enter the 1st text to encrypt: ")
	text1, _ := reader.ReadString('\n')
	fmt.Printf("Enter the 2nd text to encrypt: ")
	text2, _ := reader.ReadString('\n')

	// Remove the delimiter from the ReadString()
	text1 = text1[:len(text1)-1]
	text2 = text2[:len(text2)-1]

	// Calculate the hashes for the input texts
	hash1 := gen_sha256(text1)
	hash2 := gen_sha256(text2)

	// Count the number of bits that are different
	if hash1 != hash2 {
		fmt.Printf("Both SHA256 hashes are different.\n")

		sum1 := 0
		for i := 0; i < len(hash1); i++ {
			sum1 += popcount.PopCount(uint64(hash1[i]))
		}

		sum2 := 0
		for i := 0; i < len(hash2); i++ {
			sum2 += popcount.PopCount(uint64(hash2[i]))
		}

		diff := int(math.Abs(float64(sum2 - sum1)))
		fmt.Printf("The number of bits different between the SHA526 hashes are %d\n", diff)
	} else {
		fmt.Printf("Both SHA256 hashes are equal.\n")
		fmt.Printf("SHA256 of \"%s\": %x\n", text1, hash1)
		fmt.Printf("SHA256 of \"%s\": %x\n", text2, hash2)
	}
}