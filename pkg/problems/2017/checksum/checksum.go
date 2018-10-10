package checksum

import (
	"strconv"
	"strings"
)

func readSpreadsheet(input string) (spreadsheet [][]int) {
	for _, line := range strings.Split(input, "\n") {
		var row []int
		for _, field := range strings.Fields(line) {
			n, err := strconv.Atoi(field)
			if err != nil {
				panic(err)
			}
			row = append(row, n)
		}
		spreadsheet = append(spreadsheet, row)
	}
	return
}

type rowFunction func([]int) int

func calculateChecksum(rows [][]int, getRowValue rowFunction) (checksum int) {
	for _, row := range rows {
		checksum += getRowValue(row)
	}
	return
}

func getMaxDifference(values []int) int {
	min := values[0]
	max := values[0]
	for _, value := range values {
		if value < min {
			min = value
		} else if value > max {
			max = value
		}
	}
	return max - min
}

func getMultipleOfEvenlyDivisible(row []int) int {
	for i := 0; i < len(row); i++ {
		for j := i + 1; j < len(row); j++ {
			small, large := row[i], row[j]
			if small > large {
				small, large = large, small
			}
			if large%small == 0 {
				return large / small
			}
		}
	}
	return 0
}

// Part1 is here
func Part1(input string) interface{} {
	var spreadsheet = readSpreadsheet(input)
	return calculateChecksum(spreadsheet, getMaxDifference)
}

// Part2 is here
func Part2(input string) interface{} {
	var spreadsheet = readSpreadsheet(input)
	return calculateChecksum(spreadsheet, getMultipleOfEvenlyDivisible)
}
