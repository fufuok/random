package main

import (
	"fmt"

	"github.com/fufuok/random"
)

func main() {
	// random.Rand is a goroutine-safe rand.New() instance
	// FastIntn similar to rand.Intn or random.Rand.Intn, but faster.
	for i := 1; i <= 1000000; i *= 10 {
		fmt.Print(random.Rand.Intn(i), random.FastIntn(i), " ")
	}

	fmt.Println()

	// Fast random string generator
	fmt.Println(random.RandString(20), random.RandString(20))

	// Instantiate *rand.Rand with the current time by default
	rd := random.NewRand()
	fmt.Println(rd.Int63())

	// Returns a new pseudo-random generator seeded with the given value.
	rd = random.NewRand(1)
	fmt.Println(rd.Int63())
}
