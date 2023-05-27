package tpmath

import "testing"

func TestMedian(t *testing.T) {
	var arr = []int{1, 2, 3, 4, 5, 6}
	var got float64 = Median(arr)
	if got != 3.5 {
		t.Errorf("Median{1,2,3,4,5,6} = 3.5, got: %f", got)
	} else {
		t.Log("Success 1")
	}
	arr = []int{1, 2, 3, 4, 5}
	got = Median(arr)
	if got != 3 {
		t.Errorf("Median{1,2,3,4,5} = 3, got: %f", got)
	} else {
		t.Log("Success 2")
	}
	var farr = []float64{1.0, 2.23, 3.5, 4.5, 5.12, 6.2323}
	got = Median(farr)
	if got != 4 {
		t.Errorf("Median{farr(1)} = 8, got: %f", got)
	} else {
		t.Log("Success 3")
	}
	farr = []float64{1.1, 2.1, 3.49, 4.1, 5.1}
	got = Median(farr)
	if got != 3.49 {
		t.Errorf("Median{1.1,2.1,3.49,4.1,5.1} = 3.49, got: %f", got)
	} else {
		t.Log("Success 4")
	}
}

func TestBinarySearch(t *testing.T) {
	var arr []int
	var got int
	var expect int
	var testNo int = 1

	arr = []int{0, 1, 2, 3, 4, 5}
	got = BinarySearch(arr, 3, 0)
	expect = 3
	if got != expect {
		t.Errorf("TEST %d ERROR got: %d, expected: %d ", testNo, got, expect)
	}
	testNo++

	arr = []int{0, 1, 2, 4, 5}
	got = BinarySearch(arr, 3, 0)
	expect = -1
	if got != expect {
		t.Errorf("TEST %d ERROR got: %d, expected: %d ", testNo, got, expect)
	}
	testNo++

	arr = []int{0, 1, 2, 4, 5}
	got = BinarySearch(arr, 3, -1)
	expect = 2
	if got != expect {
		t.Errorf("TEST %d ERROR got: %d, expected: %d ", testNo, got, expect)
	}
	testNo++

	arr = []int{0, 1, 2, 4, 5}
	got = BinarySearch(arr, 3, 1)
	expect = 3
	if got != expect {
		t.Errorf("TEST %d ERROR got: %d, expected: %d ", testNo, got, expect)
	}
	testNo++

	arr = []int{1, 0, 1, 1, 0, 3, 1, 4, 1, 3, 2, 4, 0, 0, 5, 3, 1, 1, 1, 4, 2, 1, 0, 0, 1, 2, 5, 2, 6, 3, 6, 2, 4,
		1, 2, 1, 1, 1, 3, 2, 2, 0, 0, 0, 3, 3, 7, 1, 2, 1}
	Sort(arr)
	got = BinarySearch(arr, 1.72, 1)
	expect = 25
	if got != expect {
		t.Errorf("TEST %d ERROR got: %d, expected: %d ", testNo, got, expect)
	}
	testNo++

	got = BinarySearch(arr, 2.28, -1)
	expect = 33
	if got != expect {
		t.Errorf("TEST %d ERROR got: %d, expected: %d ", testNo, got, expect)
	}
	testNo++
}
