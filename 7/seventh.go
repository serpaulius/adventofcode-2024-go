package seventh

import (
	"adventofcode/2024-go/util"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Equation struct {
	result       int64
	operands     []int64
	operandCount int
}

func parseInput(lines []string) []Equation {
	var equations []Equation
	for _, line := range lines {
		split := strings.Split(line, ": ")
		result, _ := strconv.ParseInt(split[0], 10, 64)
		operandsStrs := strings.Split(split[1], " ")
		var operands = make([]int64, len(operandsStrs))
		for i, str := range operandsStrs {
			converted, _ := strconv.ParseInt(str, 10, 64)
			operands[i] = converted
		}
		equations = append(equations, Equation{result: result, operands: operands, operandCount: len(operands)})
	}
	return equations
}

var concatEnabled = false

func traverseEquation(equation Equation, result int64, operandIndex int) bool {
	if operandIndex == equation.operandCount {
		if result == equation.result {
			log.Println("OK", equation)
			return true
		} else {
			return false
		}
	}
	operand := equation.operands[operandIndex]

	var multiplyResult int64
	if operandIndex == 0 {
		multiplyResult = operand
	} else {
		multiplyResult = result * operand
	}
	if traverseEquation(equation, multiplyResult, operandIndex+1) {
		return true
	}

	addResult := result + operand
	if traverseEquation(equation, addResult, operandIndex+1) {
		return true
	}

	if concatEnabled {
		return traverseEquation(equation, util.ConcatInts(result, operand), operandIndex+1)
	}

	return false
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
	fmt.Println("7.1 - 2 operators on the bridge", result)

	// fixme: ouch
	concatEnabled = true
	result2 := traverseEquations(equations)
	fmt.Println("7.2 - 3 operators on the bridge", result2)
}
