package day_07_1

import (
	"fmt"
	"strings"

	"github.com/maikeru/adventofcode2021/utils"
)

func getOptimumPosition(positionToFuel map[int]int) int {
	var currentPos int
	var currentFuel int
	firstPos := true
	fmt.Println("position to fuel:", positionToFuel)
	for position, fuel := range positionToFuel {
		if firstPos {
			firstPos = false
			currentPos = position
			currentFuel = fuel
		} else {
			if fuel < currentFuel {
				currentPos = position
				currentFuel = fuel
			}
		}
	}
	return currentPos
}

func calculateFuelForPosition(startPositions []int, targetPosition int) int {
	total := 0
	for i := range startPositions {
		startPosition := startPositions[i]
		total += utils.Abs(targetPosition - startPosition)
	}
	return total
}

func GetFuelUsedForOptimumPositionPart1(startPositions []int) int {
	positionToFuel := make(map[int]int)
	for i := range startPositions {
		position := startPositions[i]
		fuel := calculateFuelForPosition(startPositions, position)
		positionToFuel[position] = fuel
	}

	optimumPosition := getOptimumPosition(positionToFuel)

	return positionToFuel[optimumPosition]
}

func Run() {
	// data := utils.LoadFile("./day_07_1/testInput.txt")
	data := utils.LoadFile("./day_07_1/input.txt")
	startPositions := utils.MapStringsToInts(strings.Split(data[0], ","))
	fuel := GetFuelUsedForOptimumPositionPart1(startPositions)
	fmt.Println("Optimum fuel (part1):", fuel)
}
