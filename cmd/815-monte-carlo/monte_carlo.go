// This problem was asked by Google.
// The area of a circle is defined as πr^2. Estimate π to 3 decimal places using a Monte Carlo method.
// Hint: The basic equation of a circle is x^2 + y^2 = r^2.


package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	PiTo6Decimals = 3.141593
	Billion = 1000000000
)

func PI(tries int) float64 {
	var inside int

	for i := 0; i < tries; i++ {
		x := rand.Float64()
		y := rand.Float64()
		if (x*x + y*y) < 1 {
			inside++
		}
	}

	return 4.0 * float64(inside) / float64(tries)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	tries := 1000
	for tries < Billion {
		pi := PI(tries)
		percentageDiff := 100.0 * math.Abs(PiTo6Decimals-pi)/PiTo6Decimals
		fmt.Printf("Pi with %9d samples:\t%0.6f\tDifference:\t%0.3f%%\n", tries, pi, percentageDiff)
		tries *= 10
	}
}
