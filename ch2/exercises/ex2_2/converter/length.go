// Copyright © 2021 Paulo Vital
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Length converter

package converter

import "fmt"

type Foot 	float64
type Meter	float64


//
// Foot converters
//

func (f Foot) String() string { 
	return fmt.Sprintf("%g ft", f)
}

// F2M converts a Foot length to Meter.
func F2M(f Foot) Meter {
	return Meter(f / 3.2808)
}

//
// Meter converters
//

func (m Meter) String() string {
	return fmt.Sprintf("%g m", m)
}

// M2F converts a Meter length to Foot.
func M2F(m Meter) Foot {
	return Foot(m * 3.2808)
}
