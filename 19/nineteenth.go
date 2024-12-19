package nineteenth

import (
	"adventofcode/2024-go/util"
	"fmt"
	"strings"
)

func parseInput(lines []string) ([]string, []string) {
	var towelTypes []string = []string{}
	var expectedTowels []string
	for i, line := range lines {
		if i == 0 {
			towelTypes = append(towelTypes, strings.Split(lines[i], ", ")...)
		}
		if i > 1 {
			expectedTowels = append(expectedTowels, line)
		}
	}
	return towelTypes, expectedTowels
}

func isPossible(target string, words []string) int {
	targetLength := len(target)
	canBeConstructed := make([]int, targetLength+1)
	// one can construct a 0-sized word
	canBeConstructed[0] = 1

	// for every letter check every word if it can be constructed from current letter
	for i := range targetLength {
		if canBeConstructed[i] > 0 {
			for _, candidate := range words {
				// skip too long words from this position
				if i+len(candidate) > len(target) {
					continue
				}
				// check if candidate word is in the string
				if target[i:i+len(candidate)] == candidate {
					// if yes, then it can be constructed at it's end index
					// ... and there will be exactly current amount of existing combos added to that index
					canBeConstructed[i+len(candidate)] += canBeConstructed[i]
				}
			}
		}
	}

	return canBeConstructed[targetLength]
}

func Run() {
	file := util.OpenFileOrPanicPlz("./19/input.txt")
	defer util.CloseFileOrPanicPlz(file)
	scanner := util.GiveMeAScannerPlz(file)
	input := util.ReadLines(scanner)

	towelTypes, expectations := parseInput(input)

	res := 0
	res2 := 0
	for _, ex := range expectations {
		possible := isPossible(ex, towelTypes)
		if possible > 0 {
			res += 1
			res2 += possible
		}
	}

	fmt.Println("19.1", res)
	fmt.Println("19.2", res2)
}
