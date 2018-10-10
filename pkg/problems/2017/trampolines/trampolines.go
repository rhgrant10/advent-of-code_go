package trampolines

import (
	"strconv"
	"strings"
)

func createMaze(input string) []int {
	var maze []int
	for _, instruction := range strings.Split(input, "\n") {
		offset, err := strconv.Atoi(instruction)
		if err != nil {
			panic(err)
		}
		maze = append(maze, offset)
	}

	return maze
}

func countEscapeSteps(maze []int, getAdjustment func(int) int) (steps int) {
	var index int
	for 0 <= index && index < len(maze) {
		newIndex := index + maze[index]
		maze[index] = getAdjustment(maze[index])
		index = newIndex
		steps++
	}
	return
}

func simpleIncrementer(value int) int {
	return value + 1
}

func conditionalIncrementer(value int) int {
	if value < 3 {
		return value + 1
	} else {
		return value - 1
	}
}

// Part1 function
func Part1(input string) interface{} {
	maze := createMaze(input)
	return countEscapeSteps(maze, simpleIncrementer)
}

// Part2 function
func Part2(input string) interface{} {
	maze := createMaze(input)
	return countEscapeSteps(maze, conditionalIncrementer)
}
