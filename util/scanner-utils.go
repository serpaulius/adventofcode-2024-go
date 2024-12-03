package util

import (
	"bufio"
	"strconv"
	"strings"
)

type T = int64

func ReadNumberColumns(scanner *bufio.Scanner) ([]T, []T) {
	var col1 = make([]T, 0)
	var col2 = make([]T, 0)

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		val1, _ := strconv.ParseInt(fields[0], 10, 64)
		val2, _ := strconv.ParseInt(fields[1], 10, 64)
		col1 = append(col1, val1)
		col2 = append(col2, val2)
	}

	return col1, col2
}
