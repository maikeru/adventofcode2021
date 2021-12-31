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

func getFurthestPosition(positions []int) int {
	currentPosition := 0
	for i := range positions {
		position := positions[i]
		if position > currentPosition {
			currentPosition = position
		}
	}
	return currentPosition
}

type fuelCalculator func([]int, int) int

func calculateFuelForPosition(startPositions []int, targetPosition int) int {
	total := 0
	for i := range startPositions {
		startPosition := startPositions[i]
		total += utils.Abs(targetPosition - startPosition)
	}
	return total
}

func calculateFuelForSteps(steps int) int {
	return steps * (steps + 1) / 2
}

func calculateFuelForPositionPart2(startPositions []int, targetPosition int) int {
	total := 0
	for i := range startPositions {
		startPosition := startPositions[i]
		distance := utils.Abs(targetPosition - startPosition)
		fuel := calculateFuelForSteps(distance)
		total += fuel
	}
	return total
}

func GetFuelUsedForOptimumPosition(startPositions []int, calculator fuelCalculator) int {
	positionToFuel := make(map[int]int)

	furthestPosition := getFurthestPosition(startPositions)
	for position := 0; position <= furthestPosition; position++ {
		fuel := calculator(startPositions, position)
		positionToFuel[position] = fuel
	}

	fmt.Println("position to fuel:", positionToFuel)
	optimumPosition := getOptimumPosition(positionToFuel)

	return positionToFuel[optimumPosition]
}

func GetFuelUsedForOptimumPositionPart1(startPositions []int) int {
	return GetFuelUsedForOptimumPosition(startPositions, calculateFuelForPosition)
}

func GetFuelUsedForOptimumPositionPart2(startPositions []int) int {
	return GetFuelUsedForOptimumPosition(startPositions, calculateFuelForPositionPart2)
}

func Run() {
	// data := utils.LoadFile("./day_07_1/testInput.txt")
	data := utils.LoadFile("./day_07_1/input.txt")
	startPositions := utils.MapStringsToInts(strings.Split(data[0], ","))
	fuel := GetFuelUsedForOptimumPositionPart1(startPositions)
	fmt.Println("Optimum fuel (part1):", fuel)
	part2Fuel := GetFuelUsedForOptimumPositionPart2(startPositions)
	fmt.Println("Optimum fuel (part2):", part2Fuel)
}
