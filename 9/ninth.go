package ninth

import (
	"adventofcode/2024-go/util"
	"fmt"
	"log"
	"strconv"
)

const EMPTY_SPACE int64 = -1

func expandDiskMap(diskMap string) []int64 {
	var newDisk []int64
	var fileId int64
	for i, ch := range diskMap {
		isFile := i%2 == 0
		length, _ := strconv.ParseInt(string(ch), 10, 64)
		for i := 0; i < int(length); i++ {
			if isFile {
				newDisk = append(newDisk, fileId)
			}
			if !isFile {
				newDisk = append(newDisk, -1)
			}
		}
		if isFile {
			fileId++
		}
	}
	return newDisk
}

func swapValues(runes []int64, i, j int) []int64 {
	if i < 0 || j < 0 || i >= len(runes) || j >= len(runes) {
		return runes
	}
	runes[i], runes[j] = runes[j], runes[i]
	return runes
}

func parseInput(disk string) int {
	log.Println(disk)
	expandedDisk := expandDiskMap(disk)
	log.Println(expandedDisk)

	dataLength := len(expandedDisk)
	for i := 0; i < dataLength; i++ {
		currIndex := dataLength - 1 - i
		if expandedDisk[currIndex] != EMPTY_SPACE {
			var firstEmptySpace int
			for f := 0; f < dataLength; f++ {
				if expandedDisk[f] == EMPTY_SPACE {
					firstEmptySpace = f
					break
				}
			}
			expandedDisk = swapValues(expandedDisk, currIndex, firstEmptySpace)
		}
		emptySpaceCount := 0
		for j := 0; j < currIndex; j++ {
			if expandedDisk[j] == EMPTY_SPACE {
				emptySpaceCount += 1
			}
		}
		if emptySpaceCount == 0 {
			log.Println(expandedDisk)
			break
		}
	}

	var checksum int64
	for i := 0; i < dataLength; i++ {
		if expandedDisk[i] != -1 {
			checksum += int64(i) * expandedDisk[i]
		}
	}
	return int(checksum)
}

func Run() {
	file := util.OpenFileOrPanicPlz("./9/input.txt")
	defer util.CloseFileOrPanicPlz(file)
	scanner := util.GiveMeAScannerPlz(file)
	lines := util.ReadLines(scanner)

	checksum := parseInput(lines[0])
	fmt.Println("9.1 - disk fragmenter", checksum)
}
