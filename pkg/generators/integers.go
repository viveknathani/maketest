package generators

import "math/rand"

// GenerateInts will fill up the provided array with integers in the given range.
func GenerateInts(numberOfDataPoints uint64, low int64, high int64, arr *[]Any) {

	for i := uint64(0); i < numberOfDataPoints; i++ {
		(*arr)[i] = rand.Int63n(high-low) + low
	}
}
