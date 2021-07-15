package math

import "testing"

type testpair struct {
	values  []float64
	average float64
}

type minTestPair struct {
	values []float64
	min    float64
}

type maxTestPair struct {
	values []float64
	max    float64
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
}

var minTests = []minTestPair{
	{[]float64{1, 5}, 1},
	{[]float64{6, 2, 5}, 2},
	{[]float64{9, 22, 3}, 3},
}

var maxTests = []maxTestPair{
	{[]float64{1, 5}, 5},
	{[]float64{6, 2, 5}, 6},
	{[]float64{9, 22, 3}, 22},
}

func TestAverage(t *testing.T) {
	var v float64
	v = Average([]float64{1, 2})
	if v != 1.5 {
		t.Error("Expected 1.5, got ", v)
	}
}

func TestAverage2(t *testing.T) {
	for _, pair := range tests {
		v := Average(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}

func TestMin(t *testing.T) {
	for _, pair := range minTests {
		v := Min(pair.values)
		if v != pair.min {
			t.Error(
				"For", pair.values,
				"expected", pair.min,
				"got", v,
			)
		}
	}
}

func TestMax(t *testing.T) {
	for _, pair := range maxTests {
		v := Max(pair.values)
		if v != pair.max {
			t.Error(
				"For", pair.values,
				"expected", pair.max,
				"got", v,
			)
		}
	}
}
