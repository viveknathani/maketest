package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/viveknathani/maketest/pkg/exporter"
	"github.com/viveknathani/maketest/pkg/generators"
	"github.com/viveknathani/maketest/pkg/utils"
)

const (
	asciiLogo = ` __   __  _______  ___   _  _______  _______  _______  _______  _______ 
	|  |_|  ||   _   ||   | | ||       ||       ||       ||       ||       |
	|       ||  |_|  ||   |_| ||    ___||_     _||    ___||  _____||_     _|
	|       ||       ||      _||   |___   |   |  |   |___ | |_____   |   |  
	|       ||       ||     |_ |    ___|  |   |  |    ___||_____  |  |   |  
	| ||_|| ||   _   ||    _  ||   |___   |   |  |   |___  _____| |  |   |  
	|_|   |_||__| |__||___| |_||_______|  |___|  |_______||_______|  |___|  `
	askDataPointCount = "How many data points do you need? "
	askDataType       = "What datatype do you need?: "
	askHigherValue    = "Enter higher value of your range: "
	askKeyCount       = "Enter number of keys in each data point: "
	askKeyName        = "Enter name of key %d (no spaces): "
	askLowerValue     = "Enter lower value of your range: "
	askRequirement    = "Enter your requirement: "
	colorReset        = "\033[0m"
	colorRed          = "\033[31m"
	colorGreen        = "\033[32m"
	colorYellow       = "\033[33m"
	colorBlue         = "\033[34m"
	colorPurple       = "\033[35m"
	colorCyan         = "\033[36m"
	colorWhite        = "\033[37m"
	colorPink         = "\033[38;5;13m"
	introText         = "\tWelcome to maketest, a command-line tool to generate test data."
	requirementSpecs  = "You need to enter your requirements.\n" +
		"Specify how you want your strings to be.\n" +
		"The first word should be of length 1, describing the length of string.\n" +
		"The second word should be: [X-Y] where X and Y are starting and ending\n" +
		"uppercase characters between which all characters would be included. (inclusive)\n" +
		"The third word should be: [x-y] where x and y are starting and ending\n" +
		"lowercase characters between which all characters would be included. (inclusive)\n" +
		"The fourth word should be: [a-b] where a and b define the range of numeric values\n" +
		"you need. Conditions: {0 <= a, b <= 9 and a <= b}\n" +
		"If you want omit any of the above three character sets, use \"none\"\n" +
		"Example: 5 [A-Z] [a-z] [0-9]\n" +
		"Example: 6 [A-Z] none [0-4]\n"
	closingText = "Exported to JSON!\n" +
		"Check the release.json file under the root of the current directory.\n" +
		"Thank you for using maketest!\n" +
		"- Vivek (github.com/viveknathani)\n"
)

var (
	numberOfDataPoints uint64
	numberOfKeys       uint64
	keyNames           []string
	keyValues          [][]generators.Any
)

func main() {

	fmt.Print(colorCyan)
	fmt.Println(asciiLogo)
	fmt.Println(introText)
	fmt.Println(colorReset)

	fmt.Print(colorYellow)
	fmt.Print(askDataPointCount)
	fmt.Print(colorReset)
	fmt.Scan(&numberOfDataPoints)

	fmt.Print(colorYellow)
	fmt.Print(askKeyCount)
	fmt.Print(colorReset)
	fmt.Scan(&numberOfKeys)

	keyNames = make([]string, numberOfKeys)
	keyValues = make([][]generators.Any, numberOfKeys)

	for i := uint64(0); i < numberOfKeys; i++ {
		keyValues[i] = make([]generators.Any, numberOfDataPoints)
	}

	for i := uint64(0); i < numberOfKeys; i++ {

		fmt.Print(colorYellow)
		fmt.Printf(askKeyName, i)
		fmt.Print(colorReset)
		fmt.Scan(&keyNames[i])

		fmt.Print(colorPink)
		utils.PrintMenu()
		fmt.Print(colorReset)

		fmt.Print(colorYellow)
		fmt.Print(askDataType)
		fmt.Print(colorReset)
		var dataType uint8
		fmt.Scan(&dataType)

		switch dataType {

		case utils.TYPE_INTEGER:

			var low int64
			var high int64

			fmt.Print(colorYellow)
			fmt.Print(askLowerValue)
			fmt.Print(colorReset)
			fmt.Scan(&low)

			fmt.Print(colorYellow)
			fmt.Print(askHigherValue)
			fmt.Print(colorReset)
			fmt.Scan(&high)

			generators.GenerateInts(numberOfDataPoints, low, high, &keyValues[i])

		case utils.TYPE_REAL:

			var low float64
			var high float64

			fmt.Print(colorYellow)
			fmt.Print(askLowerValue)
			fmt.Print(colorReset)
			fmt.Scan(&low)

			fmt.Print(colorYellow)
			fmt.Print(askHigherValue)
			fmt.Print(colorReset)
			fmt.Scan(&high)

			generators.GenerateRealValues(numberOfDataPoints, low, high, &keyValues[i])

		case utils.TYPE_STRING:

			fmt.Print(colorBlue)
			fmt.Print(requirementSpecs)
			fmt.Print(colorReset)

			fmt.Print(colorYellow)
			fmt.Print(askRequirement)
			fmt.Print(colorReset)
			var requirement string
			reader := bufio.NewScanner(os.Stdin)
			if reader.Scan() {
				requirement = reader.Text()
			}
			generators.GenerateStrings(numberOfDataPoints, requirement, &keyValues[i])

		case utils.TYPE_DATE:
			generators.GenerateDates(numberOfDataPoints, &keyValues[i])

		case utils.TYPE_UUID:
			generators.GenerateUUIDs(numberOfDataPoints, &keyValues[i])
		}
	}

	exporter.ExportJSON(numberOfDataPoints, numberOfKeys, keyNames, keyValues)
	fmt.Println(colorGreen)
	fmt.Println(closingText)
	fmt.Println(colorReset)
}
