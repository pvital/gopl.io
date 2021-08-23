module github.com/pvital/gopl.io/ch2/exercises/ex2_2

go 1.17

replace github.com/pvital/gopl.io/ch2/exercises/ex2_2/converter => ./converter

replace github.com/pvital/gopl.io/ch2/tempconv => ../../tempconv

require (
	github.com/pvital/gopl.io/ch2/exercises/ex2_2/converter v0.0.0-00010101000000-000000000000
	github.com/pvital/gopl.io/ch2/tempconv v0.0.0-00010101000000-000000000000
)
