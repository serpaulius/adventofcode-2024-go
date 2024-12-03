package second

import (
	"adventofcode/2024-go/util"
	"fmt"
)

func checkRow(row []int64) bool {
	rowDirection := (row[1] - row[0])
	if rowDirection == 0 {
		return false
	}

	rowIsAscending := rowDirection > 0
	for previousIndex, currentVal := range row[1:] {
		previousVal := row[previousIndex]
		delta := util.Abs(currentVal - previousVal)
		if delta > 3 ||
			delta < 1 ||
			(rowIsAscending && (previousVal > currentVal)) ||
			(!rowIsAscending && (previousVal < currentVal)) {
			return false
		}
	}
	fmt.Println(row)
	return true
}

func safetyCheck(rows [][]int64) int {
	safeCount := 0
	for _, row := range rows {
		isValid := checkRow(row)

		if isValid {
			safeCount += 1
		}
	}
	return safeCount
}

func Run() {
	file := util.OpenFileOrPanicPlz("./2/input.txt")
	defer util.CloseFileOrPanicPlz(file)
	scanner := util.GiveMeAScannerPlz(file)
	lines := util.ReadNumberLines(scanner)

	safeCount := safetyCheck(lines)
	fmt.Println(safeCount)

	problemDampenerSafeCount := safetyCheck(lines)
	fmt.Println(problemDampenerSafeCount)
}
