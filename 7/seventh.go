package seventh

import (
	"adventofcode/2024-go/util"
	"fmt"
	"strconv"
	"strings"
)

type Equation struct {
	result      int64
	members     []int64
	memberCount int
}

func parseInput(lines []string) []Equation {
	var equations []Equation
	for _, line := range lines {
		split := strings.Split(line, ": ")
		result, _ := strconv.ParseInt(split[0], 10, 64)
		membersStrs := strings.Split(split[1], " ")
		var members = make([]int64, len(membersStrs))
		for i, str := range membersStrs {
			converted, _ := strconv.ParseInt(str, 10, 64)
			members[i] = converted
		}
		equations = append(equations, Equation{result: result, members: members, memberCount: len(members)})
	}
	return equations
}

func traverseEquation(equation Equation, result int64, memberIndex int) bool {
	if memberIndex == equation.memberCount {
		if result == equation.result {
			return true
		} else {
			return false
		}
	}

	var multiplyResult int64
	if memberIndex == 0 {
		multiplyResult = equation.members[memberIndex]
	} else {
		multiplyResult = result * equation.members[memberIndex]
	}
	res1 := traverseEquation(equation, multiplyResult, memberIndex+1)

	if res1 {
		return res1
	}

	addResult := result + equation.members[memberIndex]

	res2 := traverseEquation(equation, addResult, memberIndex+1)
	return res2
}

func traverseEquations(equations []Equation) int64 {
	var sum int64
	for _, equation := range equations {
		result := traverseEquation(equation, 0, 0)
		if result {
			sum += equation.result
		}
	}
	return sum
}

func Run() {
	file := util.OpenFileOrPanicPlz("./7/input.txt")
	defer util.CloseFileOrPanicPlz(file)
	scanner := util.GiveMeAScannerPlz(file)
	lines := util.ReadLines(scanner)

	equations := parseInput(lines)
	result := traverseEquations(equations)
	fmt.Println("7.1 - operators on the bridge", result)
}
