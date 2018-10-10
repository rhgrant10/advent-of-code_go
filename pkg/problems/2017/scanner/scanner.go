package scanner

import "strings"
import "strconv"

func toInt(s string) int {
	i, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		panic(err)
	}
	return i
}

func readFirewall(input string) map[int]int {
	var firewall = make(map[int]int, 0)
	for _, line := range strings.Split(input, "\n") {
		var keypair = strings.Split(line, ":")
		var depth = toInt(keypair[0])
		var firewallRange = toInt(keypair[1])
		firewall[depth] = firewallRange
	}

	return firewall
}

func cross(firewall map[int]int) int {
	var cost = 0
	for depth, firewallRange := range firewall {
		var period = 2 * (firewallRange - 1)
		if depth%period == 0 {
			cost += depth * firewallRange
		}
	}
	return cost
}

func getDelay(firewall map[int]int) (delay int) {
	for isCostly(firewall, delay) {
		delay++
	}
	return
}

func isCostly(firewall map[int]int, delay int) bool {
	for depth, firewallRange := range firewall {
		var period = 2 * (firewallRange - 1)
		if (depth+delay)%period == 0 {
			return true
		}
	}
	return false
}

// Part1 function
func Part1(input string) interface{} {
	var firewall = readFirewall(input)
	return cross(firewall)
}

// Part2 function
func Part2(input string) interface{} {
	var firewall = readFirewall(input)
	return getDelay(firewall)
}
