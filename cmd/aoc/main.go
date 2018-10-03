package main

import (
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

	"fmt"
	"os"
	"strconv"
)

type part func([]string) interface{}
type problem []part

var problems = make(map[string]problem)

func buildProblemMap() {
	problems["captcha"] = problem{captcha.Part1, captcha.Part2}
	problems["checksum"] = problem{checksum.Part1, checksum.Part2}
	problems["spiral"] = problem{spiral.Part1, spiral.Part2}
	problems["passphrase"] = problem{passphrase.Part1, passphrase.Part2}
	problems["trampolines"] = problem{trampolines.Part1, trampolines.Part2}
	problems["reallocation"] = problem{reallocation.Part1, reallocation.Part2}
	problems["circus"] = problem{circus.Part1, circus.Part2}
	problems["registers"] = problem{registers.Part1, registers.Part2}
	problems["stream"] = problem{stream.Part1, stream.Part2}
	problems["hash"] = problem{hash.Part1, hash.Part2}
	problems["hexed"] = problem{hexed.Part1, hexed.Part2}
	problems["plumber"] = problem{plumber.Part1, plumber.Part2}
	problems["scanner"] = problem{scanner.Part1, scanner.Part2}
	problems["defragmentation"] = problem{defragmentation.Part1, defragmentation.Part2}
}

func parseArgs(args []string) (string, int, []string) {
	var name = args[0]
	part, err := strconv.Atoi(args[1])
	if err != nil {
		panic(err)
	}
	return name, part, args[2:]
}

func main() {
	buildProblemMap()
	name, part, args := parseArgs(os.Args[1:])
	var problem = problems[name][part-1]
	var answer = problem(args)
	fmt.Println(answer)
}
