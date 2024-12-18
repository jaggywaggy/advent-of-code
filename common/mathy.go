package common

func SumOfIntSlice(toSum []int) int {
	var sum int
	for i := 0; i < len(toSum); i++ {
		sum += toSum[i]
	}
	return sum
}
