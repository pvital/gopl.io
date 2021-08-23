// Copyright © 2021 Paulo Vital
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Weight converter

package converter

import "fmt"

type Pound	float64
type Gram	float64


//
// Pound converters
//

func (p Pound) String() string { 
	return fmt.Sprintf("%g lb", p)
}

// P2G converts a Pound length to Gram.
func P2G(p Pound) Gram {
	return Gram(p / 0.0022046)
}

//
// Gram converters
//

func (g Gram) String() string {
	return fmt.Sprintf("%g g", g)
}

// G2P converts a Gram length to Pound.
func G2P(g Gram) Pound {
	return Pound(g * 0.0022046)
}
