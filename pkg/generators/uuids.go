package generators

import "github.com/google/uuid"

// GenerateUUIDs will fill the list with random UUIDs
func GenerateUUIDs(numberOfDataPoints uint64, arr *[]Any) {

	for i := uint64(0); i < numberOfDataPoints; i++ {
		(*arr)[i] = uuid.New().String()
	}
}
