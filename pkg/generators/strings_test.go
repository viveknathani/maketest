package generators

import "testing"

// TestGenerateStrings will test the GenerateStrings function
// based on the requirement string: 10 [A-Z] [a-z] none
func TestGenerateStrings(tester *testing.T) {

	var requirement string = "10 [A-Z] [a-z] none"
	var numberOfDataPoints uint64 = 5
	keyValues := make([]Any, numberOfDataPoints)

	GenerateStrings(numberOfDataPoints, requirement, &keyValues)

	for i := uint64(0); i < numberOfDataPoints; i++ {

		var value string = keyValues[i].(string)

		// Check length
		if len(value) != 10 {
			tester.Fatalf("Incorrect length. Needed = %d, Produced = %d", 10, len(value))
		}

		// Check that no number should exist
		for j := 0; j < len(value); j++ {

			var asciiValue int = int(value[i])
			if asciiValue >= 48 && asciiValue <= 57 {
				tester.Fatalf("A number was produced in the string. String: %s", value)
			}
		}
	}
}
