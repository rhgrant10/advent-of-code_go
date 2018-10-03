package trampolines

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func createMaze(filename string) []int {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	instructions := strings.Trim(string(data), "\n")
	lines := strings.Split(instructions, "\n")

	var maze []int
	for _, line := range lines {
		offset, err := strconv.Atoi(line)
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
		maze[index] += getAdjustment(maze[index])
		index = newIndex
		steps += 1
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

func Part1(args []string) interface{} {
	maze := createMaze(args[0])
	return countEscapeSteps(maze, simpleIncrementer)
}

func Part2(args []string) interface{} {
	maze := createMaze(args[0])
	return countEscapeSteps(maze, conditionalIncrementer)
}
