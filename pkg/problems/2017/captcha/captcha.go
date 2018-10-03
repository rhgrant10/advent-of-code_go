package captcha

import (
	"io/ioutil"
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

func parseCaptcha(filename string) (captcha []int) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var sequence = strings.TrimSpace(string(data))
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
func Part1(args []string) interface{} {
	var captcha = parseCaptcha(args[0])
	return perform(captcha, oneAhead)
}

// Part2 is here
func Part2(args []string) interface{} {
	var captcha = parseCaptcha(args[0])
	return perform(captcha, halfwayAhead)
}
