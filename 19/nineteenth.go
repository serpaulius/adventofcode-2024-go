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

func isPossible(target string, words []string) bool {
	targetLength := len(target)
	canBeConstructed := make([]bool, targetLength+1)
	// one can construct a 0-sized word
	canBeConstructed[0] = true
	// for every letter check every word if it can be constructed from current letter
	for i := range targetLength {
		if canBeConstructed[i] {
			for _, word := range words {
				// skip too long words
				if i+len(word) > len(target) {
					continue
				}
				// check if candidate word is in the string
				if target[i:i+len(word)] == word {
					// if yes, then it can be constructed
					canBeConstructed[i+len(word)] = true
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
	for _, ex := range expectations {
		possible := isPossible(ex, towelTypes)
		// fmt.Println(ex, possible)
		if possible {
			res += 1
		}
	}

	fmt.Println("19.1", res)
}
