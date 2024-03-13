package main

import "math"

func main() {
	a := 3
	b := 5

	sum := a + b
	subtract := a - b
	multiplic := a * b
	div := a / b
	remOfDiv := a % b

	println(sum)
	println(subtract)
	println(multiplic)
	println(div)
	println(remOfDiv)

	c := 4.4

	squareRoot := math.Sqrt(c)
	round := math.Round(c)
	ceil := math.Ceil(c)
	floor := math.Floor(c)

	println(squareRoot)
	println(round)
	println(ceil)
	println(floor)
}
