package stream

func parseStream(stream string) (score int, garbage int) {
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
					garbage++
				}
			} else if c == '<' {
				inGarbage = true
			} else if c == '{' {
				depth++
				score += depth
			} else if c == '}' {
				depth--
			}
		}
	}
	return
}

// Part1 function
func Part1(input string) interface{} {
	var score, _ = parseStream(input)
	return score
}

// Part2 function
func Part2(input string) interface{} {
	var _, garbage = parseStream(input)
	return garbage
}
