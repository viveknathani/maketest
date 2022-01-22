package generators

import (
	"math/rand"
	"strconv"
	"time"
)

// GenerateDates will fill up the input list with random dates.
// A date is of the format: yyyy-mm-dd
func GenerateDates(numberOfDataPoints uint64, arr *[]Any) {

	rand.Seed(time.Now().Unix())
	for i := uint64(0); i < numberOfDataPoints; i++ {

		year := int64RandomRange(1000, 9999)
		month := int64RandomRange(1, 12)
		date := int64RandomRange(1, 28)

		var dateString string
		dateString += strconv.FormatInt(year, 10)
		dateString += "-"

		if month < 10 {
			dateString += "0"
		}

		dateString += strconv.FormatInt(month, 10)
		dateString += "-"

		if date < 10 {
			dateString += "0"
		}

		dateString += strconv.FormatInt(date, 10)

		(*arr)[i] = dateString
	}
}
