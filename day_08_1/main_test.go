package day_08_1_test

import (
	"testing"

	"github.com/maikeru/adventofcode2021/day_08_1"
	"github.com/maikeru/adventofcode2021/utils"
)

func TestCountEasyNumbersInOutput(t *testing.T) {
	data := utils.LoadFile("./testInput.txt")
	outputs := day_08_1.GetOutputs(data)
	count := day_08_1.CountEasyNumbersInOutput(outputs)
	expectedCount := 26

	if count != expectedCount {
		t.Errorf("Incorrect count! Expected %d but got %d", expectedCount, count)
	}

}
