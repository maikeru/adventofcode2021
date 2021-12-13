package day_02_2

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/maikeru/adventofcode2021/utils"
)

type position struct {
	horizontal int
	depth      int
	aim        int
}

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

func applyCommand(pos position, command string, value int) position {
	switch command {
	case "forward":
		return position{
			horizontal: pos.horizontal + value,
			depth:      pos.depth + (value * pos.aim),
			aim:        pos.aim,
		}
	case "down":
		return position{
			horizontal: pos.horizontal,
			depth:      pos.depth,
			aim:        pos.aim + value,
		}
	case "up":
		return position{
			horizontal: pos.horizontal,
			depth:      pos.depth,
			aim:        pos.aim - value,
		}
	default:
		panic(errors.New("unrecognised command: " + command))
	}
}

func getFinalPosition(route []string) position {
	pos := position{horizontal: 0,
		depth: 0,
		aim:   0}
	for i := 0; i < len(route); i++ {
		commandStr := route[i]
		command, value := getCommandAndValue(commandStr)
		fmt.Println(command, "value:", value)
		pos = applyCommand(pos, command, value)
	}
	return pos
}

func Run() {
	route := utils.LoadFile("./day_02_1/input_02.txt")
	pos := getFinalPosition(route[:])
	fmt.Println("horizontal:", pos.horizontal, "depth:", pos.depth)
	fmt.Println("product:", pos.horizontal*pos.depth)
}
