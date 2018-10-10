package captcha

import (
	"strconv"
	"strings"
)

func perform(captcha []int, getAhead func([]int) int) (sum int) {
	var size = len(captcha)
	var j = getAhead(captcha)
	for i := 0; i < size; i++ {
		if captcha[i] == captcha[j%size] {
			sum += captcha[i]
		}
		j++
	}
	return
}

func parseCaptcha(input string) (captcha []int) {
	var sequence = strings.TrimSpace(input)
	for _, char := range sequence {
		var digit, err = strconv.Atoi(string(char))
		if err != nil {
			panic(err)
		}
		captcha = append(captcha, digit)
	}
	return
}

func oneAhead(input []int) int {
	return 1
}

func halfwayAhead(input []int) int {
	return len(input) / 2
}

// Part1 is here
func Part1(input string) interface{} {
	var captcha = parseCaptcha(input)
	return perform(captcha, oneAhead)
}

// Part2 is here
func Part2(input string) interface{} {
	var captcha = parseCaptcha(input)
	return perform(captcha, halfwayAhead)
}
