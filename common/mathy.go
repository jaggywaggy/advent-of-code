package common

func SumOfIntSlice(toSum []int) int {
	var sum int
	for i := 0; i < len(toSum); i++ {
		sum += toSum[i]
	}
	return sum
}

func Abs(x int) int {
  if x < 0 {
    return -x
  }
  return x
}
