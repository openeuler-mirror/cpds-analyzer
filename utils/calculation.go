package utils

import (
	"fmt"
)

func GetSum(nums ...float64) (sum float64) {
	for _, v := range nums {
		sum += v
	}
	return
}

func GetMean(nums ...float64) (float64, error) {
	sum := GetSum(nums...)
	n := len(nums)
	if n == 0 {
		return 0, fmt.Errorf("the divisor cannot be 0")
	}
	return sum / float64(n), nil
}
