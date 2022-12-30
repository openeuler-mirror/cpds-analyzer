package utils

import (
	"testing"

	"gotest.tools/assert"
)

func TestGetSum(t *testing.T) {
	arr := make([]float64, 0)
	arr = append(arr, 1.1, 2.2, 3.3, 4.4, 5.5)

	var sum float64
	for _, n := range arr {
		sum += n
	}
	assert.Equal(t, sum, GetSum(arr...))
}
