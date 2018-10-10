package main

import (
	"os"

	arg "github.com/alexflint/go-arg"
	"github.com/rhgrant10/advent-of-code_go/pkg/problems/2017/captcha"
	"github.com/rhgrant10/advent-of-code_go/pkg/problems/2017/checksum"
	"github.com/rhgrant10/advent-of-code_go/pkg/problems/2017/circus"
	"github.com/rhgrant10/advent-of-code_go/pkg/problems/2017/defragmentation"
	"github.com/rhgrant10/advent-of-code_go/pkg/problems/2017/hash"
	"github.com/rhgrant10/advent-of-code_go/pkg/problems/2017/hexed"
	"github.com/rhgrant10/advent-of-code_go/pkg/problems/2017/passphrase"
	"github.com/rhgrant10/advent-of-code_go/pkg/problems/2017/plumber"
	"github.com/rhgrant10/advent-of-code_go/pkg/problems/2017/reallocation"
	"github.com/rhgrant10/advent-of-code_go/pkg/problems/2017/registers"
	"github.com/rhgrant10/advent-of-code_go/pkg/problems/2017/scanner"
	"github.com/rhgrant10/advent-of-code_go/pkg/problems/2017/spiral"
	"github.com/rhgrant10/advent-of-code_go/pkg/problems/2017/stream"
	"github.com/rhgrant10/advent-of-code_go/pkg/problems/2017/trampolines"
	"github.com/rhgrant10/advent-of-code_go/pkg/utilities/johnny5"

	"fmt"
	"strconv"
)

type part func(string) interface{}
type problem []part

var problems = make(map[int]problem)

func buildProblemMap() {
	problems[1] = problem{captcha.Part1, captcha.Part2}
	problems[2] = problem{checksum.Part1, checksum.Part2}
	problems[3] = problem{spiral.Part1, spiral.Part2}
	problems[4] = problem{passphrase.Part1, passphrase.Part2}
	problems[5] = problem{trampolines.Part1, trampolines.Part2}
	problems[6] = problem{reallocation.Part1, reallocation.Part2}
	problems[7] = problem{circus.Part1, circus.Part2}
	problems[8] = problem{registers.Part1, registers.Part2}
	problems[9] = problem{stream.Part1, stream.Part2}
	problems[10] = problem{hash.Part1, hash.Part2}
	problems[11] = problem{hexed.Part1, hexed.Part2}
	problems[12] = problem{plumber.Part1, plumber.Part2}
	problems[13] = problem{scanner.Part1, scanner.Part2}
	problems[14] = problem{defragmentation.Part1, defragmentation.Part2}
}

func toInt(text string) int {
	value, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}
	return value
}

var args struct {
	Year  int    `arg:"positional,required"`
	Day   int    `arg:"positional,required"`
	Part  int    `arg:"positional,required"`
	Input string `arg:"-i"`
}

func main() {
	buildProblemMap()
	p := arg.MustParse(&args)
	if args.Day < 1 || args.Day > 25 {
		p.Fail("day must be between 1 and 25, inclusive")
	}
	if args.Part != 1 && args.Part != 2 {
		p.Fail("part must be 1 or 2")
	}
	if args.Day > len(problems) {
		fmt.Printf("day %v is not yet implemented\n", args.Day)
		os.Exit(2)
	}
	var problem = problems[args.Day][args.Part-1]
	if args.Input == "" {
		input, err := johnny5.GetInput(args.Year, args.Day)
		if err != nil {
			panic(fmt.Errorf("failed to get input", err))
		}
		args.Input = input
	}
	var answer = problem(args.Input)
	fmt.Println(answer)
}
