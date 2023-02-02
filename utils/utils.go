package utils

import (
	mapset "github.com/deckarep/golang-set"
)

func MergeDuplicateIntArray(slice []int, elems []int) []int {
	listPId := append(slice, elems...)
	t := mapset.NewSet()
	for _, i := range listPId {
		t.Add(i)
	}
	var result []int
	for i := range t.Iterator().C {
		result = append(result, i.(int))
	}
	return result
}

func DuplicateIntArray(m []int) []int {
	s := make([]int, 0)
	smap := make(map[int]int)
	for _, value := range m {
		length := len(smap)
		smap[value] = 1
		if len(smap) != length {
			s = append(s, value)
		}
	}
	return s
}
