package day_04_1_test

import (
	"reflect"
	"testing"

	"github.com/maikeru/adventofcode2021/day_04_1"
	"github.com/maikeru/adventofcode2021/utils"
)

func TestLoadFile(t *testing.T) {
	data := utils.LoadFile("./testData.txt")

	bingoData := day_04_1.ParseBingoData(data)

	expectedDraws := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}
	if !reflect.DeepEqual(bingoData.Draws, expectedDraws) {
		t.Errorf("incorrect draws. Expected %v actual %v", expectedDraws, bingoData.Draws)
	}

	if len(bingoData.Cards) != 3 {
		t.Errorf("wrong number of cards! Expected 3 but got %d", len(bingoData.Cards))
	}
}
