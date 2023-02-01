package utils

import (
	"math"
	"testing"

	"gotest.tools/assert"
)

func TestGetSum(t *testing.T) {
	arr := make([]float64, 0)
	sum, err := GetSum(arr...)

	assert.Equal(t, sum, float64(-1))
	assert.Error(t, err, "invalid argument: array cannot be empty")

	arr = append(arr, 1.1, 2.2, 3.3, 4.4, 5.5)
	sum, err = GetSum(arr...)

	assert.Check(t, err)
	assert.Equal(t, sum, 16.5)
}

func TestGetMean(t *testing.T) {
	arr := make([]float64, 0)
	arr = append(arr, 1.1, 2.2, 3.3, 4.4, 5.5)

	var m1, sum float64
	for _, v := range arr {
		sum += v
	}
	m1 = sum / float64(len(arr))
	m2, err := GetMean(arr...)

	assert.Check(t, err)
	assert.Equal(t, m1, m2)
}

func TestGetVariance(t *testing.T) {
	arr := make([]float64, 0)
	arr = append(arr, 1.1, 2.2, 3.3, 4.4, 5.5)

	var m, sum float64
	for _, v := range arr {
		sum += v
	}
	m = sum / float64(len(arr))
	v1 := ((arr[0]-m)*(arr[0]-m) + (arr[1]-m)*(arr[1]-m) + (arr[2]-m)*(arr[2]-m) +
		(arr[3]-m)*(arr[3]-m) + (arr[4]-m)*(arr[4]-m)) / float64(len(arr))
	v2, err := GetVariance(arr...)

	assert.Check(t, err)
	assert.Equal(t, v1, v2)
}

func TestGetStandardDeviation(t *testing.T) {
	arr := make([]float64, 0)
	arr = append(arr, 1.1, 2.2, 3.3, 4.4, 5.5)

	var m, sum float64
	for _, v := range arr {
		sum += v
	}
	m = sum / float64(len(arr))
	variance := ((arr[0]-m)*(arr[0]-m) + (arr[1]-m)*(arr[1]-m) + (arr[2]-m)*(arr[2]-m) +
		(arr[3]-m)*(arr[3]-m) + (arr[4]-m)*(arr[4]-m)) / float64(len(arr))
	sd, err := GetStandardDeviation(arr...)

	assert.Check(t, err)
	assert.Equal(t, math.Sqrt(variance), sd)
}

func TestGetMaxValue(t *testing.T) {
	arr := make([]float64, 0)
	max, err := GetMaxValue(arr...)
	assert.Error(t, err, "invalid argument: array cannot be empty")

	arr = append(arr, 1.1, 2.2, 5.5, 4.4, 3.3)
	max, err = GetMaxValue(arr...)
	assert.Check(t, err)
	assert.Equal(t, 5.5, max)
}

func TestGetMinValue(t *testing.T) {
	arr := make([]float64, 0)
	min, err := GetMaxValue(arr...)
	assert.Error(t, err, "invalid argument: array cannot be empty")

	arr = append(arr, 1.1, 2.2, 5.5, 4.4, 3.3)
	min, err = GetMaxValue(arr...)
	assert.Check(t, err)
	assert.Equal(t, 1.1, min)
}
