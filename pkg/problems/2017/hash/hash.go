package hash

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var SUFFIX = []byte{17, 31, 73, 47, 23}

type node struct {
	Value byte
	Next  *node
}

type Stack struct {
	Top *node
}

func (s *Stack) Push(value byte) {
	node := &node{Value: value, Next: s.Top}
	s.Top = node
}

func (s *Stack) Pop() byte {
	node := s.Top
	s.Top = s.Top.Next
	return node.Value
}

func readLengths(filename string) (lengths []byte) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	content := strings.Trim(string(data), "\n")
	for _, value := range strings.Split(content, ",") {
		length, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		lengths = append(lengths, byte(length))
	}
	return
}

func readBytes(filename string) []byte {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return append(bytes.TrimSpace(data), SUFFIX...)
}

// BuildSparseHash creates a sparse hash using a given number of hash rounds.
func BuildSparseHash(lengths []byte, numRounds int) [256]byte {
	var marks [256]byte
	for i := 0; i < 256; i++ {
		marks[i] = byte(i)
	}
	var index byte
	var skip int
	for r := 0; r < numRounds; r++ {
		for _, length := range lengths {
			Twist(&marks, index, length)
			index += byte(int(length) + skip)
			skip++
		}
	}
	return marks
}

// Twist length marks starting at index.
func Twist(marks *[256]byte, index byte, length byte) {
	var segment Stack
	var i byte
	for i = 0; i < length; i++ {
		segment.Push(marks[index+i])
	}
	for i = 0; i < length; i++ {
		marks[index+i] = segment.Pop()
	}
}

// ReduceHash reduces a sparse hash to a dense hash.
func ReduceHash(sparseHash [256]byte) (denseHash []byte) {
	var value byte
	for i := 0; i < 256; i++ {
		value ^= sparseHash[i]
		if (i+1)%16 == 0 {
			denseHash = append(denseHash, value)
			value = 0
		}
	}
	return
}

func toHexString(hash []byte) (hexString string) {
	for _, b := range hash {
		hexString += fmt.Sprintf("%02x", b)
	}
	return
}

func Part1(args []string) interface{} {
	var filename = args[0]
	var lengths = readLengths(filename)
	var sparseHash = BuildSparseHash(lengths, 1)
	return int(sparseHash[0]) * int(sparseHash[1])
}

func Part2(args []string) interface{} {
	var filename = args[0]
	var data = readBytes(filename)
	var sparseHash = BuildSparseHash(data, 64)
	var denseHash = ReduceHash(sparseHash)
	return toHexString(denseHash)
}
