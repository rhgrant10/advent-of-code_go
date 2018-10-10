package plumber

import "strings"
import "strconv"

// Graph type
type Graph map[int][]int

type node struct {
	Value int
	Next  *node
}

// Stack type
type Stack struct {
	Top *node
}

func (s *Stack) Push(value int) {
	node := &node{Value: value, Next: s.Top}
	s.Top = node
}

func (s *Stack) Pop() int {
	node := s.Top
	s.Top = s.Top.Next
	return node.Value
}

func (s *Stack) IsEmpty() bool {
	return s.Top == nil
}

func readGraph(input string) Graph {
	var graph = make(Graph)
	for _, line := range strings.Split(input, "\n") {
		var fields = strings.Fields(line)
		var children []int
		for _, field := range fields[2:] {
			var childName = strings.Trim(field, ",")
			child, err := strconv.Atoi(childName)
			if err != nil {
				panic(err)
			}
			children = append(children, child)
		}

		var nodeName = fields[0]
		node, err := strconv.Atoi(nodeName)
		if err != nil {
			panic(err)
		}

		graph[node] = children
	}
	return graph
}

func getProgramsInGroup(graph Graph, group int) (programs []int) {
	visited := make([]bool, len(graph))
	var stack Stack
	stack.Push(group)

	for !stack.IsEmpty() {
		var node = stack.Pop()
		visited[node] = true
		for _, child := range graph[node] {
			if !visited[child] {
				stack.Push(child)
			}
		}
	}

	for i, v := range visited {
		if v {
			programs = append(programs, i)
		}
	}
	return
}

func findUnvisited(visited []bool) (int, bool) {
	for i, v := range visited {
		if !v {
			return i, true
		}
	}
	return 0, false
}

func countProgramGroups(graph Graph) (numGroups int) {
	visited := make([]bool, len(graph))

	for {
		start, found := findUnvisited(visited)
		if !found {
			break
		}
		var stack Stack
		stack.Push(start)
		for !stack.IsEmpty() {
			var node = stack.Pop()
			visited[node] = true
			for _, child := range graph[node] {
				if !visited[child] {
					stack.Push(child)
				}
			}
		}
		numGroups++
	}
	return
}

// Part1 function
func Part1(input string) interface{} {
	var graph = readGraph(input)
	var programs = getProgramsInGroup(graph, 0)
	return len(programs)
}

// Part2 function
func Part2(input string) interface{} {
	var graph = readGraph(input)
	var numGroups = countProgramGroups(graph)
	return numGroups
}
