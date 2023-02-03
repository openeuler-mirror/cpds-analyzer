package utils

import (
	"strings"
	"time"

	mapset "github.com/deckarep/golang-set"
	"github.com/pkg/errors"
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

func GetDifferentIntArray(sourceList, sourceList2 []int) (result []int) {
	for _, src := range sourceList {
		var find bool
		for _, target := range sourceList2 {
			if src == target {
				find = true
				continue
			}
		}
		if !find {
			result = append(result, src)
		}
	}
	return
}

func ExistIntArray(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ExistStringArray(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func GetDifferentStringArray(sourceList, sourceList2 []string) (result []string) {
	for _, src := range sourceList {
		var find bool
		for _, target := range sourceList2 {
			if src == target {
				find = true
				continue
			}
		}
		if !find {
			result = append(result, src)
		}
	}
	return
}

func MergeDuplicateStringArray(slice []string, elems []string) []string {
	listPId := append(slice, elems...)
	t := mapset.NewSet()
	for _, i := range listPId {
		t.Add(i)
	}
	var result []string
	for i := range t.Iterator().C {
		result = append(result, i.(string))
	}
	return result
}

func DuplicateStringArray(m []string) []string {
	s := make([]string, 0)
	smap := make(map[string]string)
	for _, value := range m {
		length := len(smap)
		smap[value] = value
		if len(smap) != length {
			s = append(s, value)
		}
	}
	return s
}

func ParseTimeByTimeStr(str, errPrefix string) (time.Time, error) {
	p := strings.TrimSpace(str)
	if p == "" {
		return time.Time{}, errors.Errorf("%s cannot be empty", errPrefix)
	}

	t, err := time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)
	if err != nil {
		return time.Time{}, errors.Errorf("%s wrong format", errPrefix)
	}

	return t, nil
}

func ParseTimeToInt64(t ...time.Time) int64 {
	if len(t) == 0 {
		return time.Now().UnixNano() / 1e6
	} else {
		return t[0].UnixNano() / 1e6
	}
}

func ParseSecondTimeToInt64() int64 {
	return time.Now().Unix()
}

func ParseHourTimeToInt64() int64 {
	return time.Now().Unix() / 3600 * 3600
}
