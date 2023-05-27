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
