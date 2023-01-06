package hot_100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinArray(t *testing.T) {
	min := minArray([]int{3, 4, 5, 1, 2})
	assert.Equal(t, 1, min)

	min = minArray([]int{3, 4, 5, 1, 2, 3})
	assert.Equal(t, 1, min)

	min = minArray([]int{5, 1, 2, 3, 4, 5})
	assert.Equal(t, 1, min)

	min = minArray([]int{3,1,1})
	assert.Equal(t, 1, min)

	min = minArray([]int{3,1,3})
	assert.Equal(t, 1, min)

	min = minArray([]int{1,3,3})
	assert.Equal(t, 1, min)

}
