package day_06_1

import (
	"fmt"
	"strings"

	"github.com/maikeru/adventofcode2021/utils"
)

func SimulateDay(fishCollection *[]int) *[]int {
	newFish := make([]int, 0)
	for i := range *fishCollection {
		fish := (*fishCollection)[i]
		if fish-1 < 0 {
			(*fishCollection)[i] = 6
			newFish = append(newFish, 8)
		} else {
			(*fishCollection)[i] = fish - 1
		}
	}

	// New fish should appear at the end
	result := append(*fishCollection, newFish...)
	return &result
}

func GetFishCount(initialFish []int, days int) int {
	currentFish := &initialFish
	// fmt.Println("initial fish:", initialFish)
	for i := 0; i < days; i++ {
		currentFish = SimulateDay(currentFish)
		// fmt.Printf("Day %d: %v\n", i+1, *currentFish)
	}
	return len(*currentFish)
}

func Run() {
	days := 80
	data := utils.LoadFile("./day_06_1/testInput.txt")
	// data := utils.LoadFile("./day_06_1/input.txt")
	fishStrings := strings.Split(data[0], ",")
	initialFish := utils.MapStringsToInts(fishStrings)
	fishCollection := make([]int, 0, len(initialFish)*len(initialFish))
	fishCollection = append(fishCollection, initialFish...)
	fishCount := GetFishCount(fishCollection, days)
	fmt.Printf("Total fish after %d days: %d\n", days, fishCount)
}
