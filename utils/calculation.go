package utils

import (
	"fmt"
	"strconv"

	"github.com/Knetic/govaluate"
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
		return -1, fmt.Errorf("the divisor cannot be 0")
	}
	return sum / float64(n), nil
}

// Variance: s^2 = ((x1 - m)^2 + (x2 - m)^2 + ... + (xn - m)^2) / n
func GetVariance(nums ...float64) (float64, error) {
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
	expr, _ := govaluate.NewEvaluableExpression(e)
	parameters := make(map[string]interface{})

	for index, value := range nums {
		p := fmt.Sprintf("x%s", strconv.Itoa(index+1))
		parameters[p] = value
	}
	m, err := GetMean(nums...)
	if err != nil {
		return -1, err
	}
	parameters["m"] = m
	parameters["n"] = len(nums)

	result, err := expr.Evaluate(parameters)
	if err != nil {
		return -1, err
	}
	return result.(float64), nil
}

func GetStandardDeviation(nums ...float64) (float64, error) {
	e := "[variance] ** 0.5"
	expr, _ := govaluate.NewEvaluableExpression(e)
	parameters := make(map[string]interface{})
	variance, err := GetVariance(nums...)
	if err != nil {
		return -1, err
	}
	parameters["variance"] = variance

	result, err := expr.Evaluate(parameters)
	if err != nil {
		return -1, err
	}
	return result.(float64), nil
}
