package plumber

import "bufio"

import "io"
import "os"
import "strings"
import "strconv"

type Graph map[int][]int

type node struct {
	Value int
	Next  *node
}

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

func readGraph(filename string) Graph {
	fp, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var graph = make(Graph)
	var reader = bufio.NewReader(fp)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

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

func Part1(args []string) interface{} {
	var graph = readGraph(args[0])
	var programs = getProgramsInGroup(graph, 0)
	return len(programs)
}

func Part2(args []string) interface{} {
	var graph = readGraph(args[0])
	var numGroups = countProgramGroups(graph)
	return numGroups
}
