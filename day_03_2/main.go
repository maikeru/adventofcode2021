package day_03_2

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/maikeru/adventofcode2021/utils"
)

var reportTest = []string{
	"00100",
	"11110",
	"10110",
	"10111",
	"10101",
	"01111",
	"00111",
	"11100",
	"10000",
	"11001",
	"00010",
	"01010",
}

func compareByPosition(report []string, position int) int {
	zeros := 0
	ones := 0
	for i := 0; i < len(report); i++ {
		entry := report[i]
		bit := byteToInt(entry[position])
		if bit == 0 {
			zeros++
		} else {
			ones++
		}
	}

	if zeros == ones {
		return 0
	}
	if zeros > ones {
		return -1
	}

	return 1
}

func byteToInt(value byte) int {
	asString := string(value) // Because the byte is a string representation of a number
	asInt, err := strconv.Atoi(asString)
	if err != nil {
		panic(err)
	}
	return asInt
}

func filterByValueInPosition(report []string, position int, mostCommonValue int) []string {
	filteredReport := make([]string, 0)
	for i := 0; i < len(report); i++ {
		entry := report[i]
		fmt.Println("entry", entry)
		valueAtPosition := byteToInt(entry[position])
		fmt.Println("valueAtPosition: ", valueAtPosition, "i", i, "position", position)
		if valueAtPosition == mostCommonValue {
			filteredReport = append(filteredReport, entry)
		}
	}
	return filteredReport
}

func fromBinary(binaryValue string) int {
	total := 0
	for i := len(binaryValue) - 1; i >= 0; i-- {
		bitAtPosition := byteToInt(binaryValue[i])
		if bitAtPosition == 1 {
			power := len(binaryValue) - 1 - i
			total += int(math.Pow(2, float64(power)))
		}
	}
	return total
}

const oxygenGeneratorDefaultBit = 1

func GetOxygenGeneratorRating(report []string) int {
	filteredReport := report[:]
	for i := 0; i < len(report[0]); i++ {
		fmt.Println("filteredReport: ", filteredReport)
		comparison := compareByPosition(filteredReport, i)
		fmt.Println("comparison", comparison, "position", i)
		mostCommonValue := oxygenGeneratorDefaultBit
		if comparison == -1 {
			mostCommonValue = 0
		} else if comparison == 1 {
			mostCommonValue = 1
		}
		fmt.Println("most common value: ", mostCommonValue)
		filteredReport = filterByValueInPosition(filteredReport, i, mostCommonValue)
		if len(filteredReport) == 1 {
			break
		}
	}

	if len(filteredReport) != 1 {
		panic(errors.New(fmt.Sprintf("GetOxygenGeneratorRating: incorrect filtering. filteredReport: %v", filteredReport[:])))
	}

	finalValue := filteredReport[0]
	fmt.Println("finalValue", finalValue)

	return fromBinary(finalValue)
}

const co2ScrubberDefaultBit = 0

func GetCO2ScrubberRating(report []string) int {
	filteredReport := report[:]
	for i := 0; i < len(report[0]); i++ {
		fmt.Println("filteredReport: ", filteredReport)
		comparison := compareByPosition(filteredReport, i)
		fmt.Println("comparison", comparison, "position", i)
		leastCommonValue := co2ScrubberDefaultBit
		if comparison == -1 {
			leastCommonValue = 1
		} else if comparison == 1 {
			leastCommonValue = 0
		}
		fmt.Println("least common value: ", leastCommonValue)
		filteredReport = filterByValueInPosition(filteredReport, i, leastCommonValue)
		if len(filteredReport) == 1 {
			break
		}
	}

	if len(filteredReport) != 1 {
		panic(errors.New(fmt.Sprintf("GetCO2ScrubberRating: incorrect filtering. filteredReport: %v", filteredReport[:])))
	}

	finalValue := filteredReport[0]
	fmt.Println("finalValue", finalValue)

	return fromBinary(finalValue)
}

func getLifeSupportRating(report []string) int {
	return GetOxygenGeneratorRating(report) * GetCO2ScrubberRating(report)
}

func Run() {
	report := utils.LoadFile("./day_03_1/input.txt")
	fmt.Println(getLifeSupportRating(report))
}
