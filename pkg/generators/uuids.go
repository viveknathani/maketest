package generators

import "github.com/google/uuid"

func GenerateUUIDs(numberOfDataPoints uint64, arr *[]Any) {

	for i := uint64(0); i < numberOfDataPoints; i++ {
		(*arr)[i] = uuid.New().String()
	}
}
