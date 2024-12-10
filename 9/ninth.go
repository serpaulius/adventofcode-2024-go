package ninth

import (
	"adventofcode/2024-go/util"
	"fmt"
	"log"
	"strconv"
)

const EMPTY_SPACE int64 = -1

type FileInfo struct {
	id     int64
	index  int64
	length int64
}

type FreeSpaceInfo struct {
	index  int64
	length int64
}

func expandDiskMap(diskMap []int64) ([]int64, []FileInfo, []FreeSpaceInfo) {
	var expanded []int64
	var fileId int64
	var files []FileInfo
	var spaces []FreeSpaceInfo
	var currIndex int64
	for i, length := range diskMap {
		isFile := i%2 == 0
		for i := 0; i < int(length); i++ {
			if isFile {
				expanded = append(expanded, fileId)
			} else {
				expanded = append(expanded, -1)
			}
		}
		if isFile {
			files = append(files, FileInfo{id: fileId, index: currIndex, length: length})
			fileId++
		} else {
			spaces = append(spaces, FreeSpaceInfo{index: currIndex, length: length})
		}
		currIndex += length
	}
	return expanded, files, spaces
}

func swapValues(arr []int64, i, j int) {
	if i < 0 || j < 0 || i >= len(arr) || j >= len(arr) {
		fmt.Println("Invalid positions or range.")
		return
	}
	arr[i], arr[j] = arr[j], arr[i]
}

func swapNValues(arr []int64, a, b, n int) {
	if a < 0 || b < 0 || a+n > len(arr) || b+n > len(arr) {
		fmt.Println("Invalid positions or range.")
		return
	}
	for i := 0; i < n; i++ {
		arr[a+i], arr[b+i] = arr[b+i], arr[a+i]
	}
}

func fragmentDisk(disk []int64) int {
	log.Println(disk)
	expandedDisk, _, _ := expandDiskMap(disk)
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
			swapValues(expandedDisk, currIndex, firstEmptySpace)
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

	return int(diskChecksum(expandedDisk))
}

func diskChecksum(expandedDisk []int64) int64 {
	var checksum int64
	for i := 0; i < len(expandedDisk); i++ {
		if expandedDisk[i] != -1 {
			checksum += int64(i) * expandedDisk[i]
		}
	}
	return checksum
}

func parseDisk(line string) []int64 {
	var disk []int64
	for _, ch := range line {
		entity, _ := strconv.ParseInt(string(ch), 10, 64)
		disk = append(disk, entity)
	}
	return disk
}

func nextSpace(s []FreeSpaceInfo, length int64) *FreeSpaceInfo {
	for i := 0; i < len(s); i++ {
		space := &s[i]
		if space.length >= length {
			return space
		}
	}
	return nil
}

func printDisk(disk []int64) {
	for _, d := range disk {
		if d != -1 {
			fmt.Print(d)
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}

// for each file with decreasing id
// find an empty space from start
// move if found a space
// continue with another file backwards
// fixme: works, though very improveable
func defragDisk(disk []int64) int64 {
	expandedDisk, files, spaces := expandDiskMap(disk)
	for i := 0; i < len(files); i++ {
		file := files[len(files)-1-i]
		// printDisk(expandedDisk)
		space := nextSpace(spaces, file.length)
		if space != nil && space.index < file.index {
			swapNValues(expandedDisk, int(file.index), int(space.index), int(file.length))
			file.index = space.index
			space.length -= file.length
			space.index += file.length
		}
	}
	return diskChecksum(expandedDisk)
}

func Run() {
	file := util.OpenFileOrPanicPlz("./9/input.txt")
	defer util.CloseFileOrPanicPlz(file)
	scanner := util.GiveMeAScannerPlz(file)
	lines := util.ReadLines(scanner)

	disk := parseDisk(lines[0])
	checksum := fragmentDisk(disk)
	fmt.Println("9.1 - disk fragmenter", checksum)

	checksum2 := defragDisk(disk)
	fmt.Println("9.2 - disk defragmenter", checksum2)
}
