package defragmentation

import "github.com/rhgrant10/advent-of-code_go/pkg/problems/2017/hash"
import "fmt"

const NUM_KNOTS = 256
const NUM_KEYS = 128
const NUM_ROUNDS = 64

var SUFFIX = []byte{17, 31, 73, 47, 23}

var UP = Coord{0, -1}
var DOWN = Coord{0, 1}
var LEFT = Coord{-1, 0}
var RIGHT = Coord{1, 0}
var NEIGHBORS = [4]Coord{UP, DOWN, LEFT, RIGHT}

type SparseHash = [NUM_KNOTS]byte
type Coord [2]int

type ByteNode struct {
	Value byte
	Next  *ByteNode
}

type ByteStack struct {
	Top *ByteNode
}

func (s *ByteStack) Push(value byte) {
	node := &ByteNode{Value: value, Next: s.Top}
	s.Top = node
}

func (s *ByteStack) Pop() byte {
	node := s.Top
	s.Top = s.Top.Next
	return node.Value
}

type CoordNode struct {
	Value Coord
	Next  *CoordNode
}

type CoordStack struct {
	Top  *CoordNode
	size int
}

func (s *CoordStack) Push(value Coord) {
	node := &CoordNode{Value: value, Next: s.Top}
	s.Top = node
	s.size++
}

func (s *CoordStack) Pop() Coord {
	node := s.Top
	s.Top = s.Top.Next
	s.size--
	return node.Value
}

func (s *CoordStack) IsEmpty() bool {
	return s.size == 0
}

func getKnotHash(text string) []byte {
	var lengths = append([]byte(text), hash.SUFFIX...)
	var sparseHash = hash.BuildSparseHash(lengths, NUM_ROUNDS)
	var denseHash = hash.ReduceHash(sparseHash)
	return denseHash
}

func getDiskState(keyString string) map[Coord]bool {
	var grid = make(map[Coord]bool, 0)
	for y := 0; y < NUM_KEYS; y++ {
		var rowKey = fmt.Sprintf("%s-%d", keyString, y)
		var knotHash = getKnotHash(rowKey)
		for x, bit := range getBits(knotHash) {
			var coord = Coord{x, y}
			grid[coord] = bit == '1'
		}
	}
	return grid
}

func getBits(knotHash []byte) (bits string) {
	for _, knot := range knotHash {
		bits += fmt.Sprintf("%08b", knot)
	}
	return
}

func countUsed(grid map[Coord]bool) (count int) {
	for _, isUsed := range grid {
		if isUsed {
			count++
		}
	}
	return
}

func countIslands(grid map[Coord]bool) (count int) {
	var visited = make(map[Coord]bool, 0)

	for coord, isUsed := range grid {
		if !isUsed || visited[coord] {
			continue
		}

		count++
		var stack CoordStack
		stack.Push(coord)
		visited[coord] = true

		for !stack.IsEmpty() {
			coord = stack.Pop()
			for _, neighbor := range getConnectedNeighbors(coord) {
				if !visited[neighbor] && grid[neighbor] {
					stack.Push(neighbor)
				}
				visited[neighbor] = true
			}
		}
	}
	return
}

func getConnectedNeighbors(coord Coord) (neighbors []Coord) {
	for _, offset := range NEIGHBORS {
		var neighbor = move(coord, offset)
		if isInbounds(neighbor) {
			neighbors = append(neighbors, neighbor)
		}
	}
	return
}

func move(coord Coord, offset Coord) Coord {
	return Coord{coord[0] + offset[0], coord[1] + offset[1]}
}

func isInbounds(coord Coord) bool {
	return 0 <= coord[0] && coord[0] < 128 &&
		0 <= coord[1] && coord[1] < 128
}

func Part1(args []string) interface{} {
	var grid = getDiskState(args[0]) // args[0] isn't a filename :(
	return countUsed(grid)
}

func Part2(args []string) interface{} {
	var grid = getDiskState(args[0]) // args[0] isn't a filename :(
	return countIslands(grid)
}
