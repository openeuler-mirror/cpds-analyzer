package utils

import (
	"fmt"
	"math"
	"strconv"

	"github.com/Knetic/govaluate"
)

func GetSum(nums ...float64) (float64, error) {
	var sum float64
	if len(nums) == 0 {
		return math.NaN(), fmt.Errorf("invalid argument: array cannot be empty")
	}

	for _, v := range nums {
		sum += v
	}
	return sum, nil
}

func GetMean(nums ...float64) (float64, error) {
	sum, err := GetSum(nums...)
	if err != nil {
		return math.NaN(), err
	}

	n := len(nums)
	if n == 0 {
		return math.NaN(), fmt.Errorf("the divisor cannot be 0")
	}
	return sum / float64(n), nil
}

// Variance: s^2 = ((x1 - m)^2 + (x2 - m)^2 + ... + (xn - m)^2) / n
func GetVariance(nums ...float64) (float64, error) {
	if len(nums) == 0 {
		return math.NaN(), fmt.Errorf("invalid argument: array cannot be empty")
	}

	e := fmt.Sprintf("(%s) / n", func(nums ...float64) (s string) {
		for index := range nums {
			if index == 0 {
				s = fmt.Sprintf("(x%s - m)**2", strconv.Itoa(index+1))
				continue
			}
			s = fmt.Sprintf("%s + (x%s - m)**2", s, strconv.Itoa(index+1))
		}
		return
	}(nums...))
	expr, err := govaluate.NewEvaluableExpression(e)
	if err != nil {
		return math.NaN(), err
	}
	parameters := make(map[string]interface{})

	for index, value := range nums {
		p := fmt.Sprintf("x%s", strconv.Itoa(index+1))
		parameters[p] = value
	}
	m, err := GetMean(nums...)
	if err != nil {
		return math.NaN(), err
	}
	parameters["m"] = m
	parameters["n"] = len(nums)

	result, err := expr.Evaluate(parameters)
	if err != nil {
		return math.NaN(), err
	}
	return result.(float64), nil
}

func GetStandardDeviation(nums ...float64) (float64, error) {
	if len(nums) == 0 {
		return math.NaN(), fmt.Errorf("invalid argument: array cannot be empty")
	}

	e := "[variance] ** 0.5"
	expr, err := govaluate.NewEvaluableExpression(e)
	if err != nil {
		return math.NaN(), err
	}
	parameters := make(map[string]interface{})
	variance, err := GetVariance(nums...)
	if err != nil {
		return math.NaN(), err
	}
	parameters["variance"] = variance

	result, err := expr.Evaluate(parameters)
	if err != nil {
		return math.NaN(), err
	}
	return result.(float64), nil
}

func GetMaxValue(nums ...float64) (float64, error) {
	if n := len(nums); n == 0 {
		return math.NaN(), fmt.Errorf("invalid argument: array cannot be empty")
	}

	max := nums[0]
	for _, v := range nums {
		max = math.Max(max, v)
	}
	return max, nil
}

func GetMinValue(nums ...float64) (float64, error) {
	if n := len(nums); n == 0 {
		return -1, fmt.Errorf("invalid argument: array cannot be empty")
	}

	min := nums[0]
	for _, v := range nums {
		min = math.Min(min, v)
	}
	return min, nil
}
