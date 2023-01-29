package hot_100

import "testing"

func TestMedianFinder(t *testing.T) {
	finder := Constructor()
	finder.AddNum(-1)
	finder.AddNum(-2)
	res := finder.FindMedian()
	println(res)
}
