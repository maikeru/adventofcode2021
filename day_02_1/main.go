package day_02_1

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/maikeru/adventofcode2021/utils"
)

var routeTest = [6]string{
	"forward 5",
	"down 5",
	"forward 8",
	"up 3",
	"down 8",
	"forward 2",
}

func getCommandAndValue(input string) (string, int) {
	tokens := strings.Split(input, " ")
	value, err := strconv.Atoi(tokens[1])
	if err != nil {
		panic(err)
	}
	return tokens[0], value
}

func applyCommand(horizontal int, depth int, command string, value int) (int, int) {
	switch command {
	case "forward":
		return horizontal + value, depth
	case "down":
		return horizontal, depth + value
	case "up":
		return horizontal, depth - value
	default:
		panic(errors.New("unrecognised command: " + command))
	}
}

func getFinalPosition(route []string) (int, int) {
	horizontal := 0
	depth := 0
	for i := 0; i < len(route); i++ {
		commandStr := route[i]
		command, value := getCommandAndValue(commandStr)
		fmt.Println(command, "value:", value)
		horizontal, depth = applyCommand(horizontal, depth, command, value)
	}
	return horizontal, depth
}

func Run() {
	route := utils.LoadFile("./day_02_1/input_02.txt")
	horizontal, depth := getFinalPosition(route[:])
	fmt.Println("horizontal:", horizontal, "depth:", depth)
	fmt.Println("product:", horizontal*depth)
}
