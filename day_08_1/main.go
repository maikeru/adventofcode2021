package day_08_1

import (
	"fmt"
	"strings"

	"github.com/maikeru/adventofcode2021/utils"
)

func isEasyNumber(entry string) bool {
	size := len(entry)
	switch size {
	case 2, 4, 3, 7: // 1, 4, 7, 8
		return true
	}
	return false
}

func CountEasyNumbersInOutput(outputs [][]string) int {
	total := 0
	for i := range outputs {
		outputLine := outputs[i]
		for j := range outputLine {
			output := outputLine[j]
			if isEasyNumber(output) {
				total++
			}
		}
	}
	return total
}

func GetOutputs(data []string) [][]string {
	outputs := make([][]string, 0)
	for i := range data {
		line := data[i]
		inputsAndOutput := strings.Split(line, " | ")
		output := strings.Fields(inputsAndOutput[1])
		outputs = append(outputs, output)
	}
	return outputs
}

func Run() {
	// data := utils.LoadFile("./day_08_1/testInput.txt")
	data := utils.LoadFile("./day_08_1/input.txt")
	outputs := GetOutputs(data)
	count := CountEasyNumbersInOutput(outputs)
	fmt.Println("PART 1: count of easy numbers", count)
}
