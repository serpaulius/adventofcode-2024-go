package util

func Abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func IntLength(i int64) int64 {
	if i == 0 {
		return 1
	}
	var count int64
	for i != 0 {
		i /= 10
		count++
	}
	return count
}

func AddZeros(value int64, zeros int64) int64 {
	var i int64
	for i = 0; i < zeros; i++ {
		value = value * 10
	}
	return value
}

func ConcatInts(a int64, b int64) int64 {
	return AddZeros(a, IntLength(b)) + b
}

// greatest common divider
func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
