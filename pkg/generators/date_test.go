package generators

import (
	"strconv"
	"strings"
	"testing"
)

func checkStringWithInt64Range(str string, low int64, high int64) bool {

	num, err := strconv.ParseInt(str, 10, 0)

	if err != nil {
		return false
	}

	return (num >= low && num <= high)
}

// TestGenerateTests will test GenerateTests to see if the generated dates
// are valid.
func TestGenerateDates(tester *testing.T) {

	var numberOfDataPoints uint64 = 5
	keyValues := make([]Any, numberOfDataPoints)

	GenerateDates(numberOfDataPoints, &keyValues)
	for i := uint64(0); i < numberOfDataPoints; i++ {

		var value string = keyValues[i].(string)
		var parts []string = strings.Split(value, "-")

		// Check if 3 parts exist
		if len(parts) != 3 {
			tester.Fatal("Unable to split string into 3 parts.")
		}
		// Check if year is valid
		if !checkStringWithInt64Range(parts[0], 1000, 9999) {
			tester.Fatalf("Year is invalid, %s.", parts[0])
		}

		// Check if month is valid
		if !checkStringWithInt64Range(parts[1], 1, 12) {
			tester.Fatalf("Month is invalid, %s.", parts[1])
		}

		// Check if date is valid
		if !checkStringWithInt64Range(parts[2], 1, 28) {
			tester.Fatalf("Year is invalid, %s.", parts[0])
		}
	}
}
