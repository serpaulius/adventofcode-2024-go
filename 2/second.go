package second

import (
	"adventofcode/2024-go/util"
	"fmt"
	"log"
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
	return true
}

func recheckRowWithDampener(row []int64) bool {
	log.Println("using dampener for ", row)
	for i := range row {
		// fixme: slicing is ugly here
		sliced := []int64{}
		sliced = append(sliced, row[:i]...)
		sliced = append(sliced, row[i+1:]...)
		log.Println(i, sliced)
		if checkRow(sliced) {
			return true
		}
	}
	return false
}

func safetyCheck(rows [][]int64, tolerateOne bool) int {
	safeCount := 0
	for _, row := range rows {
		isValid := checkRow(row)
		if tolerateOne && !isValid {
			isValid = recheckRowWithDampener(row)
		}

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

	safeCount := safetyCheck(lines, false)
	fmt.Println("2.1 - safety check", safeCount)

	problemDampenerSafeCount := safetyCheck(lines, true)
	fmt.Println("2.2 - safety with dampener", problemDampenerSafeCount)
}
