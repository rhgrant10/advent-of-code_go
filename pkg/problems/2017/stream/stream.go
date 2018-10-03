package stream

import (
	"io/ioutil"
	"strings"
)

func readInputFile(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return strings.Trim(string(data), "\n")
}

func Parse(stream string) (score int, garbage int) {
	var depth = 0
	var inGarbage = false
	var skip = false

	for _, c := range stream {
		if skip {
			skip = false
		} else if c == '!' {
			skip = true
		} else {
			if inGarbage {
				inGarbage = c != '>'
				if inGarbage {
					garbage += 1
				}
			} else if c == '<' {
				inGarbage = true
			} else if c == '{' {
				depth += 1
				score += depth
			} else if c == '}' {
				depth -= 1
			}
		}
	}
	return
}

func Part1(args []string) interface{} {
	var filename = args[0]
	var stream = readInputFile(filename)
	var score, _ = Parse(stream)
	return score
}

func Part2(args []string) interface{} {
	var filename = args[0]
	var stream = readInputFile(filename)
	var _, garbage = Parse(stream)
	return garbage
}
