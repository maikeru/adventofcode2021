package day_01_2

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

func mapSlidingWindow(nums []int) []int {
	totals := make([]int, len(nums)-2)
	for i := 0; i < len(nums)-2; i++ {
		totals[i] = nums[i] + nums[i+1] + nums[i+2]
	}
	return totals
}

func getIncreasedCountOverSlidingWindow(depths []int) int {
	totals := mapSlidingWindow(depths)

	increasedCount := 0
	previousDepthTotal := 0
	for i := 0; i < len(totals); i++ {
		currentDepthTotal := totals[i]
		if i == 0 {
			previousDepthTotal = currentDepthTotal
			continue
		}

		if previousDepthTotal < currentDepthTotal {
			increasedCount++
		}

		previousDepthTotal = currentDepthTotal
	}
	return increasedCount
}

func Run() {
	fmt.Println(getIncreasedCountOverSlidingWindow(depthsTest[:]))

	depthBytes, err := ioutil.ReadFile("./input-01-1.txt")
	if err != nil {
		panic(err)
	}
	depthString := string(depthBytes[:])
	depthStrings := strings.Split(depthString, "\n")
	depths := utils.MapStringsToInts(depthStrings)

	fmt.Println(getIncreasedCountOverSlidingWindow(depths))
}
