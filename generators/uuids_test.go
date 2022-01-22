package generators

import (
	"testing"

	"github.com/google/uuid"
)

func isValidUUID(str string) bool {
	_, err := uuid.Parse(str)
	return err == nil
}

// TestGenerateUUIDs will test GenerateUUIDs to see if the generated values
// are valid UUIDs
func TestGenerateUUIDs(tester *testing.T) {

	var numberOfDataPoints uint64 = 5
	keyValues := make([]Any, numberOfDataPoints)

	GenerateUUIDs(numberOfDataPoints, &keyValues)

	for i := uint64(0); i < numberOfDataPoints; i++ {

		var value string = keyValues[i].(string)
		if !isValidUUID(value) {
			tester.Fatal("Generated value is not a valid UUID")
		}
	}
}
