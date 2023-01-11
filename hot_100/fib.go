package hot_100

func fib(n int) int64 {
	if n < 2 {
		return int64(n)
	}
	var res int64
	a := int64(0)
	b := int64(1)
	for i := 1; i < n ; i++ {
		res = a + b
		a = b
		b = res
	}

	return res % 1000000007
}
