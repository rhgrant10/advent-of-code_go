package hexed

import (
	"io/ioutil"
	"strings"
)

type Point [3]int

var DIRECTIONS = map[string]Point{
	"n":  [3]int{0, 1, -1},
	"s":  [3]int{0, -1, 1},
	"ne": [3]int{1, 0, -1},
	"sw": [3]int{-1, 0, 1},
	"nw": [3]int{-1, 1, 0},
	"se": [3]int{1, -1, 0},
}

var CENTER = [3]int{0, 0, 0}

func readDirections(filename string) []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var contents = strings.TrimSpace(string(data))
	return strings.Split(contents, ",")
}

func getLocation(directions []string, start Point) Point {
	var coords = follow(directions, start)
	return coords[len(coords)-1]
}

func follow(directions []string, start Point) []Point {
	coords := []Point{start}
	for i, direction := range directions {
		var coord = move(coords[i], DIRECTIONS[direction])
		coords = append(coords, coord)
	}
	return coords
}

func move(point Point, offset Point) (result Point) {
	result[0] = point[0] + offset[0]
	result[1] = point[1] + offset[1]
	result[2] = point[2] + offset[2]
	return
}

func getDistance(end Point, start Point) (distance int) {
	for i := 0; i < 3; i += 1 {
		var difference = end[i] - start[i]
		if difference > distance {
			distance = difference
		}
	}
	return
}

func Part1(args []string) interface{} {
	var directions = readDirections(args[0])
	var location = getLocation(directions, CENTER)
	return getDistance(location, CENTER)
}

func Part2(args []string) interface{} {
	var directions = readDirections(args[0])
	var locations = follow(directions, CENTER)
	var maxDistance int
	for _, location := range locations {
		var distance = getDistance(location, CENTER)
		if distance > maxDistance {
			maxDistance = distance
		}
	}
	return maxDistance
}
