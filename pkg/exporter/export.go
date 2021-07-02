package exporter

import (
	"fmt"
	"log"
	"os"

	"github.com/viveknathani/maketest/pkg/generators"
)

// ExportJSON will export the input values to a JSON file.
func ExportJSON(numberOfDataPoints uint64, numberOfKeys uint64,
	keyNames []string, keyValues [][]generators.Any) {

	file, err := os.Create("release.json")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(file, "[")
	for j := uint64(0); j < numberOfDataPoints; j++ {

		fmt.Fprintln(file, "  {")
		for i := uint64(0); i < numberOfKeys; i++ {
			fmt.Fprint(file, "    ")
			fmt.Fprint(file, "\"")
			fmt.Fprint(file, keyNames[i]+"\": ")
			fmt.Fprint(file, "\"")
			fmt.Fprint(file, keyValues[i][j])
			fmt.Fprint(file, "\"")

			if i < (numberOfKeys - 1) {
				fmt.Fprint(file, ",")
			}
			fmt.Fprintln(file)
		}
		fmt.Fprint(file, "  }")
		if j < (numberOfDataPoints - 1) {
			fmt.Fprint(file, ",")
		}
		fmt.Fprintln(file)
	}
	fmt.Fprintln(file, "]")
}
