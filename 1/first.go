package first

import (
	"adventofcode/2024-go/util"
	"fmt"
	"slices"
)

type T = int64

func sumSortedNumberDifferences(col1unsorted, col2unsorted []T) T {
	var col1, col2 = make([]T, len(col1unsorted)), make([]T, len(col2unsorted))
	copy(col1, col1unsorted)
	copy(col2, col2unsorted)
	slices.Sort(col1)
	slices.Sort(col2)
	var sum T
	for i, val := range col1 {
		absas := util.Abs(val - col2[i])
		sum += absas
	}
	return sum
}

func mapOccurences(arr []T) map[T]T {
	occurences := map[T]T{}
	for _, val := range arr {
		occurences[val] += 1
	}
	return occurences
}

func similarityScore(col1, col2 []T) T {
	col2valueCounts := mapOccurences(col2)
	var similaritySum T
	for _, value := range col1 {
		similaritySum += value * col2valueCounts[value]
	}
	return similaritySum
}

func Run() {
	file := util.OpenFileOrPanicPlz("./1/input.txt")
	defer util.CloseFileOrPanicPlz(file)
	scanner := util.GiveMeAScannerPlz(file)

	col1, col2 := util.ReadNumberColumns(scanner)
	sum := sumSortedNumberDifferences(col1, col2)

	fmt.Println("1.1 - sum of diffs", sum)

	similarityIndex := similarityScore(col1, col2)
	fmt.Println("1.2 - similarity score", similarityIndex)
}
