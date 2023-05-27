package tpmath

import (
	"math"
)

type Number interface {
	int | float64
}

func DistributionFunc[T Number](arr []T) map[T]float64 {
	res := make(map[T]float64)
	for _, v := range arr {
		_, ok := res[v]
		if !ok {
			res[v] = 0
		}
		res[v]++
	}
	unique := make([]T, 0, len(res))
	for key, _ := range res {
		unique = append(unique, key)
		res[key] = res[key] / float64(len(arr))
	}
	Sort(unique)
	for i, _ := range unique {
		if i == 0 {
			continue
		}
		res[unique[i]] += res[unique[i-1]]
	}
	return res
}

func Probability[T Number](arr []T, a, b float64) float64 {
	f := DistributionFunc(arr)
	return f[arr[BinarySearch(arr, b, 1)]] - f[arr[BinarySearch(arr, a, -1)]]
}

func Excess[T Number](arr []T) float64 {
	m4 := float64(SumMinValPower(arr, T(Expect(arr)), 4)) / float64(len(arr))
	sig := Sigma(arr)
	return m4/(sig*sig*sig*sig) - 3
}

func Asymmetry[T Number](arr []T) float64 {
	m3 := float64(SumMinValPower(arr, T(Expect(arr)), 3)) / float64(len(arr))
	sig := Sigma(arr)
	return m3 / (sig * sig * sig)
}

func Median[T Number](arr []T) float64 {
	switch len(arr) % 2 {
	case 0:
		return float64(arr[len(arr)/2]+arr[-1+len(arr)/2]) / 2
	default:
		return float64(arr[len(arr)/2])
	}
}

func Variance[T Number](arr []T) float64 {
	return float64(SumMinValPower(arr, T(Expect(arr)), 2)) / float64(len(arr))
}

func Sigma[T Number](arr []T) float64 {
	return math.Sqrt(Variance(arr))
}

func Expect[T Number](arr []T) float64 {
	return float64(Sum(arr) / T(len(arr)))
}
