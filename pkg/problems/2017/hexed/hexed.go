package hexed

import (
	"strings"
)

// Point type
type Point [3]int

var cardinals = map[string]Point{
	"n":  [3]int{0, 1, -1},
	"s":  [3]int{0, -1, 1},
	"ne": [3]int{1, 0, -1},
	"sw": [3]int{-1, 0, 1},
	"nw": [3]int{-1, 1, 0},
	"se": [3]int{1, -1, 0},
}

var center = [3]int{0, 0, 0}

func readDirections(input string) []string {
	return strings.Split(input, ",")
}

func getLocation(directions []string, start Point) Point {
	var coords = follow(directions, start)
	return coords[len(coords)-1]
}

func follow(directions []string, start Point) []Point {
	coords := []Point{start}
	for i, direction := range directions {
		var coord = move(coords[i], cardinals[direction])
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
	for i := 0; i < 3; i++ {
		var difference = end[i] - start[i]
		if difference > distance {
			distance = difference
		}
	}
	return
}

// Part1 function
func Part1(input string) interface{} {
	var directions = readDirections(input)
	var location = getLocation(directions, center)
	return getDistance(location, center)
}

// Part2 function
func Part2(input string) interface{} {
	var directions = readDirections(input)
	var locations = follow(directions, center)
	var maxDistance int
	for _, location := range locations {
		var distance = getDistance(location, center)
		if distance > maxDistance {
			maxDistance = distance
		}
	}
	return maxDistance
}
