package generators

import (
	"math/big"
	"testing"
)

// TestGenerateRealValues will test the GenerateRealValues function for
// 500 data points to see if they fit within the provided range.
func TestGenerateRealValues(tester *testing.T) {

	var numberOfDataPoints uint64 = 500
	keyValues := make([]Any, numberOfDataPoints)
	var low float64 = -1000
	var high float64 = 2000000

	GenerateRealValues(numberOfDataPoints, low, high, &keyValues)

	for i := uint64(0); i < numberOfDataPoints; i++ {

		var value = keyValues[i].(float64)
		lowComp := big.NewFloat(value).Cmp(big.NewFloat(low))
		highComp := big.NewFloat(value).Cmp(big.NewFloat(high))

		if lowComp == -1 || highComp == 1 {
			tester.Fatalf("GenerateInts produced a value outside range = [%f, %f]. Value = %f.",
				low, high, value)
		}
	}
}
