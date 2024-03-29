package generators

import (
	"math/rand"
	"time"
)

// GenerateRealValues will fill up the input array with pseudo random real
// values in the given range.
func GenerateRealValues(numberOfDataPoints uint64, low float64, high float64, arr *[]Any) {

	rand.Seed(time.Now().Unix())
	for i := uint64(0); i < numberOfDataPoints; i++ {
		(*arr)[i] = low + rand.Float64()*(high-low)
	}
}
