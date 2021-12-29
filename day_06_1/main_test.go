package day_06_1_test

import (
	"reflect"
	"testing"

	"github.com/maikeru/adventofcode2021/day_06_1"
)

func TestGetFishCount(t *testing.T) {
	initialFish := []int{3, 4, 3, 1, 2}
	days := 18
	expectedFishCount := 26
	finalFishCount := day_06_1.GetFishCount(initialFish, days)

	if finalFishCount != expectedFishCount {
		t.Errorf("Incorrect fish count after %d days: got %d but expected %d", days, finalFishCount, expectedFishCount)
	}
}

func TestSimulateDayWithoutBirths(t *testing.T) {
	initialFish := map[int]int{0: 1, 1: 1, 2: 1, 3: 2, 4: 1, 5: 0, 6: 0, 7: 0, 8: 0}
	expectedFish := map[int]int{0: 1, 1: 1, 2: 2, 3: 1, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}
	actualFish := day_06_1.SimulateDay(initialFish)

	if !reflect.DeepEqual(expectedFish, actualFish) {
		t.Errorf("Incorrect fish! Expected %v but was %v", expectedFish, actualFish)
	}
}

func TestSimulateDayWithBirths(t *testing.T) {
	initialFish := map[int]int{0: 1, 1: 1, 2: 2, 3: 1, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}
	expectedFish := map[int]int{0: 1, 1: 2, 2: 1, 3: 0, 4: 0, 5: 0, 6: 1, 7: 0, 8: 1}
	actualFish := day_06_1.SimulateDay(initialFish)

	if !reflect.DeepEqual(expectedFish, actualFish) {
		t.Errorf("Incorrect fish! Expected %v but was %v", expectedFish, actualFish)
	}
}
