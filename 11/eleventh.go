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

func naiveBlink(stones []int64, blinks int) int {
	log.Println("start", stones)
	for i := 0; i < int(blinks); i++ {
		stones = blink(stones)
		log.Println(i, stones)
	}
	log.Println("end", stones)
	return len(stones)
}

const blinksToDo int64 = 75

type StoneKey struct {
	stone  int64
	blinks int64
}

func recursivelyCachedBlink(stone int64, blinksLeft int64, cache map[StoneKey]int64) int64 {
	currentBlink := blinksToDo - blinksLeft + 1
	stoneKey := StoneKey{stone: stone, blinks: currentBlink}

	if cached := cache[stoneKey]; cached > 0 {
		return cached
	}
	if blinksLeft == 0 {
		return 1
	}
	if stone == 0 {
		count := recursivelyCachedBlink(1, blinksLeft-1, cache)
		cache[stoneKey] = count
		return count
	}
	if length := util.IntLength(stone); length%2 == 0 {
		halfway := util.AddZeros(1, length/2)
		left := stone / halfway
		right := stone % halfway
		countL := recursivelyCachedBlink(left, blinksLeft-1, cache)
		countR := recursivelyCachedBlink(right, blinksLeft-1, cache)
		cache[stoneKey] = countL + countR
		return countL + countR
	}
	count := recursivelyCachedBlink(stone*2024, blinksLeft-1, cache)
	cache[stoneKey] = count
	return count
}

func Run() {
	file := util.OpenFileOrPanicPlz("./11/input.txt")
	defer util.CloseFileOrPanicPlz(file)
	scanner := util.GiveMeAScannerPlz(file)
	input := util.ReadNumberLines(scanner)[0]

	result1 := naiveBlink(input, 25)
	fmt.Println("11.1 - blink25", result1)

	var sum int64
	var cache map[StoneKey]int64 = make(map[StoneKey]int64)
	for _, one := range input {
		sum += recursivelyCachedBlink(one, blinksToDo, cache)
	}
	fmt.Println("11.2 - blink75", sum)
}
