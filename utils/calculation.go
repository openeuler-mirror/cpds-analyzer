package utils

func GetSum(nums ...float64) (sum float64) {
	for _, v := range nums {
		sum += v
	}
	return
}
