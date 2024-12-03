package third

import (
	"adventofcode/2024-go/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func parseMul(str string) (int64, int64) {
	modified, _ := strings.CutPrefix(str, "mul(")
	modified, _ = strings.CutSuffix(modified, ")")
	numbers := strings.Split(modified, ",")
	number1, _ := strconv.ParseInt(numbers[0], 10, 64)
	number2, _ := strconv.ParseInt(numbers[1], 10, 64)
	return number1, number2
}

func mulcheck(lines []string) int64 {
	muls := make([]int64, 0)
	r, _ := regexp.Compile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	for _, line := range lines {
		mulMatches := r.FindAllString(line, -1)

		fmt.Println(mulMatches)

		for _, mulMatch := range mulMatches {
			num1, num2 := parseMul(mulMatch)
			muls = append(muls, num1*num2)
			fmt.Println(num1, num2, muls)
		}
	}
	var sum int64
	for _, mul := range muls {
		sum += mul
	}
	return sum
}

func Run() {
	file := util.OpenFileOrPanicPlz("./3/input.txt")
	defer util.CloseFileOrPanicPlz(file)
	scanner := util.GiveMeAScannerPlz(file)
	lines := util.ReadLines(scanner)

	safeCount := mulcheck(lines)
	fmt.Println(safeCount)
}
