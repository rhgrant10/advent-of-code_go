package scanner

import "io/ioutil"

import "strings"
import "strconv"

func toInt(s string) int {
	i, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		panic(err)
	}
	return i
}

func readFirewall(filename string) map[int]int {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var content = strings.TrimSpace(string(data))
	var lines = strings.Split(content, "\n")

	var firewall = make(map[int]int, 0)
	for _, line := range lines {
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

func Part1(args []string) interface{} {
	var firewall = readFirewall(args[0])
	return cross(firewall)
}

func Part2(args []string) interface{} {
	var firewall = readFirewall(args[0])
	return getDelay(firewall)
}
