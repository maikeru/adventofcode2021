package day_01_1

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/maikeru/adventofcode2021/utils"
)

var depthsTest = [10]int{199,
	200,
	208,
	210,
	200,
	207,
	240,
	269,
	260,
	263}

func getIncreasedCount(depths []int) int {
	increasedCount := 0
	previousDepth := 0
	for i := 0; i < len(depths); i++ {
		currentDepth := depths[i]
		if i == 0 {
			fmt.Println("Skipping first index")
			previousDepth = currentDepth
			continue
		}

		fmt.Println("previousDepth", previousDepth, "< currentDepth", currentDepth, previousDepth < currentDepth)
		if previousDepth < currentDepth {
			increasedCount++
		}

		previousDepth = currentDepth
	}
	return increasedCount
}

func Run() {
	fmt.Println(getIncreasedCount(depthsTest[:]))

	depthBytes, err := ioutil.ReadFile("./input-01-1.txt")
	if err != nil {
		panic(err)
	}
	depthString := string(depthBytes[:])
	depthStrings := strings.Split(depthString, "\n")
	depths := utils.MapStringsToInts(depthStrings)

	fmt.Println(getIncreasedCount(depths))
}
