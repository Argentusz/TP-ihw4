package tpmath

import (
	"sort"
)

func Sum[T Number](arr []T) T {
	var res T
	for _, v := range arr {
		res += v
	}
	return res
}

func SumMinValPower[T Number](arr []T, val T, power int) T {
	var res T = 0
	for _, v := range arr {
		var temp T = 1
		for i := 0; i < power; i++ {
			temp *= v - val
		}
		res += temp
	}
	return res
}

func CountIf[T Number](arr []T, f func(T) bool) T {
	var res T
	for _, v := range arr {
		if f(v) {
			res++
		}
	}
	return res
}

func SumIf[T Number](arr []T, f func(T) bool) T {
	var res T
	for _, v := range arr {
		if f(v) {
			res += v
		}
	}
	return res
}

func Indicator[T Number](x T, f func(T) bool) int {
	if f(x) {
		return 1
	}
	return 0
}

func Sort[T Number](arr []T) {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
}

func Sorted[T Number](arr []T) []T {
	var res []T
	copy(res, arr)
	sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })
	return res
}

func BinarySearch[T Number](arr []T, val float64, mode int) int {
	i, j := 0, len(arr)
	for i < j {
		h := int(uint(i+j) >> 1)
		if float64(arr[h]) < val {
			i = h + 1
		} else {
			j = h
		}
	}
	if float64(arr[i])-val != 0 {
		if float64(arr[i]) < val && mode == 1 {
			return i + 1
		} else if float64(arr[i]) > val && mode == -1 {
			return i - 1
		}
		if mode == 0 {
			return -1
		}
	}
	return i
}
