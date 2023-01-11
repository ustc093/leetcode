package hot_100

import (
	"math"
	"strconv"
)

func translateNum(num int) int {
	if num <= 26 && num >= 10 {
		return 2
	} else if num <= 9{
		return 1
	}
	n := len(strconv.Itoa(num))

	first := num/int(math.Pow(10, float64(n-1))) % 10
	second := num/int(math.Pow(10, float64(n-2))) % 10
	tmp := first *10 + second
	if tmp <= 0 || tmp > 26 {
		i := num - first * int(math.Pow(10, float64(n-1)))
		return translateNum(i)
	}
	i := num - first * int(math.Pow(10, float64(n-1)))
	j := num - first * int(math.Pow(10, float64(n-1))) - second * int(math.Pow(10, float64(n-2)))
	return translateNum(i) + translateNum(j)
}

