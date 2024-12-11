package eleventh

import (
	"adventofcode/2024-go/util"
	"fmt"
	"log"
)

func blink(stones []int64) []int64 {
	var newArr []int64
	// apply relevant rule to each stone, return new array
	for i := 0; i < len(stones); i++ {
		intLength := util.IntLength(stones[i])
		if stones[i] == 0 {
			newArr = append(newArr, 1)
		} else if intLength%2 == 0 {
			half2 := stones[i] % util.AddZeros(1, intLength/2)
			half1 := stones[i] / util.AddZeros(1, (intLength/2))
			newArr = append(newArr, half1, half2)
		} else {
			newArr = append(newArr, stones[i]*2024)
		}
	}

	return newArr
}

func blinkAtStones(stones []int64, blinks int) int {
	log.Println("start", stones)
	for i := 0; i < int(blinks); i++ {
		stones = blink(stones)
		log.Println(i, stones)
	}
	log.Println("end", stones)
	return len(stones)
}

func Run() {
	file := util.OpenFileOrPanicPlz("./11/input.txt")
	defer util.CloseFileOrPanicPlz(file)
	scanner := util.GiveMeAScannerPlz(file)
	input := util.ReadNumberLines(scanner)[0]

	result1 := blinkAtStones(input, 25)
	fmt.Println("11.1 - stones", result1)
}
