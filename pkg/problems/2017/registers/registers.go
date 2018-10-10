package registers

import "strconv"
import "strings"

type RegisterSet map[string]int

type Operator func(a int, b int) int

func inc(a int, b int) int { return a + b }
func dec(a int, b int) int { return a - b }

type Comparator func(a int, b int) bool

func eq(a int, b int) bool { return a == b }
func ne(a int, b int) bool { return a != b }
func lt(a int, b int) bool { return a < b }
func gt(a int, b int) bool { return a > b }
func le(a int, b int) bool { return a <= b }
func ge(a int, b int) bool { return a >= b }

var OPERATORS = map[string]Operator{
	"inc": inc,
	"dec": dec,
}

var COMPARATORS = map[string]Comparator{
	"==": eq,
	"!=": ne,
	"<":  lt,
	">":  gt,
	"<=": le,
	">=": ge,
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

func parseInstruction(line string) (func(RegisterSet), func(RegisterSet) bool) {
	parts := strings.Split(line, " if ")
	performOperation := parseOperation(parts[0])
	checkCondition := parseCondition(parts[1])
	return performOperation, checkCondition
}

func parseOperation(operation string) func(RegisterSet) {
	fields := strings.Fields(operation)
	register := fields[0]
	symbol := fields[1]
	operand, err := strconv.Atoi(fields[2])
	if err != nil {
		panic(err)
	}

	return func(registers RegisterSet) {
		registers[register] = OPERATORS[symbol](registers[register], operand)
	}
}

func parseCondition(condition string) func(RegisterSet) bool {
	fields := strings.Fields(condition)
	register := fields[0]
	symbol := fields[1]
	operand, err := strconv.Atoi(fields[2])
	if err != nil {
		panic(err)
	}

	return func(registers RegisterSet) bool {
		return COMPARATORS[symbol](registers[register], operand)
	}
}

func execute1(instructions []string, registers RegisterSet) (maxValue int) {
	for _, line := range instructions {
		performOperation, checkCondition := parseInstruction(line)
		if checkCondition(registers) {
			performOperation(registers)
		}
	}
	maxValue = findMaxValue(registers)
	return
}

func execute2(instructions []string, registers RegisterSet) (maxValue int) {
	for _, line := range instructions {
		performOperation, checkCondition := parseInstruction(line)
		if checkCondition(registers) {
			performOperation(registers)
			max := findMaxValue(registers)
			if max > maxValue {
				maxValue = max
			}
		}
	}
	return
}

func findMaxValue(registers RegisterSet) (max int) {
	for _, v := range registers {
		if v > max {
			max = v
		}
	}
	return
}

// Part1 function
func Part1(input string) interface{} {
	var instructions = parseInput(input)
	var registers = make(RegisterSet)
	return execute1(instructions, registers)
}

// Part2 function
func Part2(input string) interface{} {
	var instructions = parseInput(input)
	var registers = make(RegisterSet)
	return execute2(instructions, registers)
}
