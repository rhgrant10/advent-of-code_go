package spiral

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
)

var ZERO = big.NewInt(0)
var ONE = big.NewInt(1)
var TWO = big.NewInt(2)

var LEFT = [2]int{-1, 0}
var RIGHT = [2]int{1, 0}
var UP = [2]int{0, 1}
var DOWN = [2]int{0, -1}

var UP_LEFT = [2]int{-1, 1}
var UP_RIGHT = [2]int{1, 1}
var DOWN_LEFT = [2]int{-1, -1}
var DOWN_RIGHT = [2]int{1, -1}

var CARDINALS = [4][2]int{RIGHT, UP, LEFT, DOWN}
var DIAGNOALS = [4][2]int{UP_LEFT, UP_RIGHT, DOWN_LEFT, DOWN_RIGHT}

var ADJACENTS = [8][2]int{RIGHT, UP, LEFT, DOWN, UP_LEFT, UP_RIGHT, DOWN_LEFT, DOWN_RIGHT}

func calculateSpiralManhattan(square uint64) uint64 {
	if square == 1 {
		return 0
	}

	// Find the first axis.
	squareFloor := uint64(math.Sqrt(float64(square - 1)))
	ringMaxRoot := squareFloor + (squareFloor % 2) + 1
	ringMax := uint64(math.Pow(float64(ringMaxRoot), 2))
	ringID := (ringMaxRoot+1)/2 - 1

	// Find the perpendicular axis.
	ringSize := ringMax - uint64(math.Pow(math.Sqrt(float64(ringMax))-2, 2))
	edgeIndex := ringSize - (ringMax - square)
	perpendicular := uint64(math.Abs(float64(edgeIndex%(ringID*2) - ringID)))

	// Manhattan distance is sum of axis movements
	distance := ringID + perpendicular
	return distance
}

// use big.Int
func calculateBigSpiralManhattan(square *big.Int) *big.Int {

	if square.Cmp(ONE) == 0 {
		return ZERO
	}

	// Make serious use of a temp variable
	t := new(big.Int)

	// squareFloor := int(math.Sqrt(float64(square - 1)))
	t.Sub(square, ONE)
	squareFloor := new(big.Int).Sqrt(t)

	// ringMaxRoot := squareFloor + (squareFloor % 2) + 1
	t.Mod(squareFloor, TWO)
	t.Add(t, ONE)
	ringMaxRoot := new(big.Int).Add(t, squareFloor)

	// ringMax := int(math.Pow(float64(ringMaxRoot), 2))
	ringMax := new(big.Int).Exp(ringMaxRoot, TWO, nil)

	// ringID := (ringMaxRoot + 1) / 2 - 1
	t.Add(ringMaxRoot, ONE)
	t.Div(t, TWO)
	ringID := new(big.Int).Sub(t, ONE)

	// ringSize := ringMax - int(math.Pow(math.Sqrt(float64(ringMax)) - 2, 2))
	t.Sqrt(ringMax)
	t.Sub(t, TWO)
	t.Exp(t, TWO, nil)
	ringSize := new(big.Int).Sub(ringMax, t)

	// edgeIndex := ringSize - (ringMax - square)
	t.Sub(ringMax, square)
	edgeIndex := new(big.Int).Sub(ringSize, t)

	// perpendicular := int(math.Abs(float64(edgeIndex % (ringID * 2) - ringID)))
	t.Mul(ringID, TWO)
	t.Mod(edgeIndex, t)
	t.Sub(t, ringID)
	perpendicular := new(big.Int).Abs(t)

	// distance := ringID + perpendicular
	distance := t.Add(ringID, perpendicular)
	return distance
}

func readBigSquare(input string) *big.Int {
	var square = new(big.Int)
	_, ok := square.SetString(input, 10)
	if !ok {
		panic(fmt.Errorf("could not parse integer"))
	}
	return square
}

func move(point [2]int, offset [2]int) [2]int {
	return [2]int{point[0] + offset[0], point[1] + offset[1]}
}

func getNeighborSum(point [2]int, values map[[2]int]int) int {
	sum := 0
	for _, offset := range ADJACENTS {
		sum += values[move(point, offset)]
	}
	return sum
}

func getFirstSquareGreaterThan(target int) int {
	length := 1
	point := [2]int{0, 0}
	adjustment := false

	squares := make(map[[2]int]int)
	squares[[2]int{0, 0}] = 1

	for {
		for _, direction := range CARDINALS {
			for i := 0; i < length; i++ {
				point = move(point, direction)
				squares[point] = getNeighborSum(point, squares)
				if squares[point] > target {
					return squares[point]
				}
			}

			if adjustment {
				length += 1
			}
			adjustment = !adjustment
		}
	}
}

// Part1 function
func Part1(input string) interface{} {
	var square = readBigSquare(input)
	if square.IsUint64() {
		return calculateSpiralManhattan(square.Uint64())
	} else {
		return calculateBigSpiralManhattan(square)
	}
}

// Part2 function
func Part2(input string) interface{} {
	target, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return getFirstSquareGreaterThan(target)
}
