package day_03_1

import (
	"fmt"
	"math"
	"strings"

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

func getMostCommon(counts []int, midpoint int) []int {
	result := make([]int, len(counts))
	for i := 0; i < len(counts); i++ {
		if counts[i] > midpoint {
			result[i] = 1
		}
	}
	return result
}

func GetGammaAndEpsilon(binaryReport []int) (int, int) {
	gamma := 0
	epsilon := 0

	for i := 0; i < len(binaryReport); i++ {
		position := len(binaryReport) - 1 - i // "highest" position is earliest so invert
		increment := int(math.Pow(2, float64(position)))
		if binaryReport[i] == 1 {
			gamma += increment
		} else {
			epsilon += increment
		}
	}
	return gamma, epsilon
}

func getPower(report []string) int {
	reportLineSize := len(report[0])
	counts := make([]int, reportLineSize)
	for i := 0; i < len(report); i++ {
		current := report[i]
		bits := utils.MapStringsToInts(strings.Split(current, ""))
		for j := 0; j < reportLineSize; j++ {
			counts[j] += bits[j]
		}
	}
	midpoint := len(report) / 2
	mostCommon := getMostCommon(counts, midpoint)
	gamma, epsilon := GetGammaAndEpsilon(mostCommon)
	return gamma * epsilon
}

func Run() {
	report := utils.LoadFile("./day_03_1/input.txt")
	fmt.Println(getPower(report))
}
