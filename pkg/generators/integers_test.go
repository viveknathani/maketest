package generators

import (
	"testing"
)

// TestGenerateInts will test the GenerateInts function for
// 500 data points to see if they fit within the provided range.
func TestGenerateInts(tester *testing.T) {

	var numberOfDataPoints uint64 = 500
	keyValues := make([]Any, numberOfDataPoints)
	var low int64 = -1000
	var high int64 = 2000000

	GenerateInts(numberOfDataPoints, low, high, &keyValues)

	for i := uint64(0); i < numberOfDataPoints; i++ {

		var value = keyValues[i].(int64)
		if value < low || value > high {
			tester.Fatalf("GenerateInts produced a value outside range = [%d, %d]. Value = %d.",
				low, high, value)
		}
	}
}
