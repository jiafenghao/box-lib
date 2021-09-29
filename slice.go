/**
 * @brief
 * @file slice
 * @author zhangpeng
 * @version 1.0
 * @date
 */

package box_lib

import "math"

func UniqueInt64Slice(s []int64) []int64 {
	occurred := map[int64]bool{}
	var result []int64
	for _, value := range s {
		if occurred[value] != true {
			occurred[value] = true
			result = append(result, value)
		}
	}
	return result
}

func UniqueStringSlice(s []string) []string {
	occurred := map[string]bool{}
	var result []string
	for _, value := range s {
		if occurred[value] != true {
			occurred[value] = true
			result = append(result, value)
		}
	}
	return result
}

func ExistsForInt64(v int64, s []int64) bool {
	for _, val := range s {
		if val == v {
			return true
		}
	}
	return false
}

func ExistsForString(v string, s []string) bool {
	for _, val := range s {
		if val == v {
			return true
		}
	}
	return false
}

func SplitInt64Slice(s []int64, size int) [][]int64 {
	if size < 1 {
		size = 1
	}
	length := len(s)
	chunks := int(math.Ceil(float64(length) / float64(size)))
	var n [][]int64
	for i, end := 0, 0; chunks > 0; chunks-- {
		end = (i + 1) * size
		if end > length {
			end = length
		}
		n = append(n, s[i*size:end])
		i++
	}
	return n
}
