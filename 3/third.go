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

func matchAllLines(expression string, lines []string) []string {
	r, _ := regexp.Compile(expression)

	var matches []string
	for _, line := range lines {
		matches = append(matches, r.FindAllString(line, -1)...)
	}
	return matches
}

func sumMuls(mulMatches []string) int64 {
	var sum int64
	for _, mulMatch := range mulMatches {
		num1, num2 := parseMul(mulMatch)
		multiplied := num1 * num2
		// fmt.Println(num1, num2, multiplied)
		sum += multiplied
	}
	return sum
}

func mulcheck(lines []string) int64 {
	mulMatches := matchAllLines(`mul\([0-9]{1,3},[0-9]{1,3}\)`, lines)
	return sumMuls(mulMatches)
}

func mulcheckNoDonts(lines []string) int64 {
	matches := matchAllLines(`mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don't\(\)`, lines)
	var noDonts []string
	ignoreMode := false
	for _, match := range matches {
		if strings.Compare(match, "do()") == 0 {
			ignoreMode = false
		}
		if strings.Compare(match, "don't()") == 0 {
			ignoreMode = true
		}
		if !ignoreMode && strings.HasPrefix(match, "mul") {
			noDonts = append(noDonts, match)
		}
	}
	return sumMuls(noDonts)
}

func Run() {
	file := util.OpenFileOrPanicPlz("./3/input.txt")
	defer util.CloseFileOrPanicPlz(file)
	scanner := util.GiveMeAScannerPlz(file)
	lines := util.ReadLines(scanner)

	mulSum := mulcheck(lines)
	fmt.Println(mulSum)

	mulSumNoDonts := mulcheckNoDonts(lines)
	fmt.Println(mulSumNoDonts)
}
