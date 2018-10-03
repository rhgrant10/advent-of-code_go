package circus

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type node struct {
	Name     string
	Weight   int
	Children []*node
}

type level struct {
	Node    node
	Weights []int
}

// NodeStack is a stack of nodes
type NodeStack struct {
	data []*node
}

func (s *NodeStack) Push(node *node) {
	s.data = append(s.data, node)
}

func (s *NodeStack) Pop() *node {
	i := len(s.data) - 1
	node := s.data[i]

	s.data[i] = nil // avoid memory leaks
	s.data = s.data[:i]

	return node
}

func (s *NodeStack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *NodeStack) Size() int {
	return len(s.data)
}

//
// Level stack
//
type LevelStack struct {
	data []*level
}

func (s *LevelStack) Push(level *level) {
	s.data = append(s.data, level)
}

func (s *LevelStack) Pop() *level {
	i := len(s.data) - 1
	level := s.data[i]

	s.data[i] = nil
	s.data = s.data[:i]

	return level
}

func (s *LevelStack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *LevelStack) Size() int {
	return len(s.data)
}

func parseTree(filename string) (map[string]string, map[string][]string, map[string]int) {
	fp, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	parents := make(map[string]string)
	weights := make(map[string]int)
	children := make(map[string][]string)

	reader := bufio.NewReader(fp)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		node, weight, childNodes := parseLine(line)
		weights[node] = weight
		for _, childNode := range childNodes {
			parents[childNode] = node
			children[node] = append(children[node], childNode)
		}
	}

	return parents, children, weights
}

func parseLine(line string) (name string, weight int, childNames []string) {
	fields := strings.Fields(line)

	name = fields[0]

	weightField := fields[1]
	weight, err := strconv.Atoi(weightField[1 : len(weightField)-1])
	if err != nil {
		panic(err)
	}

	if strings.Contains(line, "->") {
		for _, childField := range fields[3:] {
			child := strings.Trim(childField, ",")
			childNames = append(childNames, child)
		}
	}

	return
}

func buildTree(name string, weights map[string]int, children map[string][]string) *node {
	var root = &node{Name: name, Weight: weights[name]}
	var stack NodeStack
	stack.Push(root)

	for !stack.IsEmpty() {
		parent := stack.Pop()
		for _, child := range children[parent.Name] {
			childNode := &node{Name: child, Weight: weights[child]}
			parent.Children = append(parent.Children, childNode)
			stack.Push(childNode)
		}
	}

	return root
}

func findImbalance(root *node) (int, error) {
	var stack LevelStack
	stack.Push(&level{Node: *root})

	for !stack.IsEmpty() {
		top := stack.Pop()

		if !top.IsComplete() {
			// Visit the next child.
			i := len(top.Weights)
			stack.Push(top)
			stack.Push(&level{Node: *top.Node.Children[i]})
		} else {
			// All children have been visited. Are they balanced?
			common, unique, uniqueFound := getCommonAndUnique(top.Weights)
			if uniqueFound {
				// Children are not balanced. Return the corrected weight for
				// the uniquely weighted child.
				diff := common - unique
				for i, weight := range top.Weights {
					if weight == unique {
						return top.Node.Children[i].Weight + diff, nil
					}
				}
				return 0, fmt.Errorf("could not find the uniquely weighted child")
			}

			// Children are balanced, so sum them with the node and add the total
			// weight to the parent's list of child weights.
			total := top.Node.Weight
			for _, w := range top.Weights {
				total += w
			}
			top = stack.Pop()
			top.Weights = append(top.Weights, total)
			stack.Push(top)
		}
	}

	return 0, fmt.Errorf("no imbalance found")
}

func getCommonAndUnique(weights []int) (common int, unique int, uniqueFound bool) {
	if len(weights) < 3 {
		return common, unique, false
	}

	var i = 1
	for i < len(weights) {
		if weights[i] != weights[0] {
			uniqueFound = true
			if i > 1 || weights[i+1] == weights[0] {
				common = weights[0]
				unique = weights[i]
			} else {
				common = weights[i]
				unique = weights[0]
			}
			break
		}
		i += 1
	}

	return
}

// True if we have all children's cumulative weights.
func (level *level) IsComplete() bool {
	return len(level.Node.Children) == len(level.Weights)
}

func findRoot(parents map[string]string, weights map[string]int) (string, error) {
	for node, _ := range weights {
		_, present := parents[node]
		if !present {
			return node, nil
		}
	}
	return "", fmt.Errorf("No root found! You sure this is a tree?")
}

// Part1 is here
func Part1(args []string) interface{} {
	filename := os.Args[1]
	parents, _, nodes := parseTree(filename)
	root, err := findRoot(parents, nodes)
	if err != nil {
		panic(err)
	}
	return root
}

// Part2 is here
func Part2(args []string) interface{} {
	filename := os.Args[1]
	parents, children, weights := parseTree(filename)
	root, err := findRoot(parents, weights)
	if err != nil {
		panic(err)
	}

	tree := buildTree(root, weights, children)
	imbalance, err := findImbalance(tree)
	if err != nil {
		panic(err)
	}
	return imbalance
}
