package generators

import (
	"math/rand"
	"time"
)

// Utility function to generate a random number within a range
func int64RandomRange(low int64, high int64) int64 {
	return rand.Int63n(high-low) + low
}

// GenerateInts will fill up the provided array with integers in the given range.
func GenerateInts(numberOfDataPoints uint64, low int64, high int64, arr *[]Any) {

	rand.Seed(time.Now().Unix())
	for i := uint64(0); i < numberOfDataPoints; i++ {
		(*arr)[i] = int64RandomRange(low, high)
	}
}
