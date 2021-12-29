package day_06_1

import (
	"fmt"
	"strings"

	"github.com/maikeru/adventofcode2021/utils"
)

// Credit to u/Evargalo for the idea of using a map like this
// https://www.reddit.com/r/adventofcode/comments/r9z49j/comment/hq75dbm/?utm_source=share&utm_medium=web2x&context=3
func SimulateDay(fishMap map[int]int) map[int]int {
	motherFishCount := fishMap[0]
	// Decrease time left for all fish not giving birth
	for i := 1; i < 9; i++ {
		fishMap[i-1] = fishMap[i]
	}
	// roll birthing fish around to 6 days left
	fishMap[6] += motherFishCount

	// birth new fish equivalent to number of mother fish
	fishMap[8] = motherFishCount
	return fishMap
}

func createFishMap(initialFish []int) map[int]int {
	fishMap := make(map[int]int, 8)
	for i := range initialFish {
		currentFish := initialFish[i]
		fishMap[currentFish] = fishMap[currentFish] + 1
	}
	return fishMap
}

func GetFishCount(initialFish []int, days int) int {
	fishMap := createFishMap(initialFish)

	for i := 0; i < days; i++ {
		fishMap = SimulateDay(fishMap)
		// fmt.Printf("Day %d: %v\n", i+1, *currentFish)
	}

	var total int
	for i := 0; i < 9; i++ {
		total += fishMap[i]
	}

	return total
}

func Run() {
	days := 256
	// data := utils.LoadFile("./day_06_1/testInput.txt")
	data := utils.LoadFile("./day_06_1/input.txt")
	fishStrings := strings.Split(data[0], ",")
	initialFish := utils.MapStringsToInts(fishStrings)
	fishCollection := make([]int, 0, len(initialFish)*len(initialFish))
	fishCollection = append(fishCollection, initialFish...)
	fishCount := GetFishCount(fishCollection, days)
	fmt.Printf("Total fish after %d days: %d\n", days, fishCount)
}
