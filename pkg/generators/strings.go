package generators

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

// GenerateStrings will fill up the given array with generated strings.
// The strings would be generated as per the requirement parameter.
// The requirement string should be of 4 words. It defines how a string
// should be.
// The first word should be of length 1, describing the length of string.
// The second word should be: [X-Y] where X and Y are starting and ending
// uppercase characters between which all characters would be included. (inclusive)
// The third word should be: [x-y] where x and y are starting and ending
// lowercase characters between which all characters would be included. (inclusive)
// The fourth word should be: [a-b] where a and b define the range of numeric values
// you need. Conditions: {0 <= a, b <= 9 and a <= b}
// If you want omit any of the above three character sets, use "none"
// Example: 5 [A-Z] [a-z] [0-9]
// Example: 6 [A-Z] none [0-4]
func GenerateStrings(numberOfDataPoints uint64, requirement string, arr *[]Any) {

	var splitSet []string = strings.Split(requirement, " ")

	possibleValues := make([]int, 0)

	for i := 1; i < len(splitSet); i++ {

		if splitSet[i] != "none" {
			for j := int(splitSet[i][1]); j <= int(splitSet[i][3]); j++ {
				possibleValues = append(possibleValues, j)
			}
		}
	}

	strLength, err := strconv.Atoi(splitSet[0])
	if err != nil {
		fmt.Println("something went wrong")
	}

	for i := uint64(0); i < numberOfDataPoints; i++ {

		var str string
		for j := 0; j < strLength; j++ {
			str += string(rune(possibleValues[rand.Int()%len(possibleValues)]))
		}

		(*arr)[i] = str
	}
}
