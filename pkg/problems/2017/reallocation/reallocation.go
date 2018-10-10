package reallocation

import "os"
import "io/ioutil"
import "strings"
import "strconv"

func banksFromString(data string) (banks []int) {
	fields := strings.Fields(string(data))
	for _, field := range fields {
		bank, err := strconv.Atoi(field)
		if err != nil {
			panic(err)
		}
		banks = append(banks, bank)
	}
	return
}

func banksToString(banks []int) string {
	var elements []string
	for _, bank := range banks {
		elements = append(elements, strconv.Itoa(bank))
	}
	return strings.Join(elements, "\t")
}

func redistribute(data string) string {
	banks := banksFromString(data)

	// find the largest bank and it's index
	var index, maxBlocks int
	for i, blocks := range banks {
		if blocks > maxBlocks {
			index = i
			maxBlocks = blocks
		}
	}

	// redistrubte the largest bank
	banks[index] = 0
	for maxBlocks > 0 {
		index = (index + 1) % len(banks)
		banks[index]++
		maxBlocks--
	}

	return banksToString(banks)
}

func countReallocationCycles(banks string) (count int) {
	seen := make(map[string]bool)

	hasKey := false
	for !hasKey {
		seen[banks] = true
		banks = redistribute(banks)
		_, hasKey = seen[banks]
		count++
	}
	return
}

func countReallocationLoopSize(banks string) (count int) {
	seen := make(map[string]bool)

	hasKey := false
	for !hasKey {
		seen[banks] = true
		banks = redistribute(banks)
		_, hasKey = seen[banks]
	}

	count = 1
	target := banks
	banks = redistribute(banks)
	for banks != target {
		banks = redistribute(banks)
		count++
	}

	return
}

func parseInputFile(filename string) string {
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	return banksToString(banksFromString(string(data)))
}

// Part1 function
func Part1(input string) interface{} {
	return countReallocationCycles(input)
}

// Part2 function
func Part2(input string) interface{} {
	return countReallocationLoopSize(input)
}
