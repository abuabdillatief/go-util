package slices

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}, expected: []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}},
		{input: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{input: []int{5, 5, 5, 5, 5}, expected: []int{5, 5, 5, 5, 5}},
		{input: []int{}, expected: []int{}},
	}

	for _, test := range tests {
		QuickSort(test.input)
		if !reflect.DeepEqual(test.input, test.expected) {
			t.Errorf("QuickSort(%v) = %v; expected %v", test.input, test.input, test.expected)
		}
	}
}

func TestGetMax(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{input: []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}, expected: 9},
		{input: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, expected: 10},
		{input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, expected: 10},
		{input: []int{5, 5, 5, 5, 5}, expected: 5},
	}

	for _, test := range tests {
		result := GetMax(test.input)
		if result != test.expected {
			t.Errorf("GetMax(%v) = %v; expected %v", test.input, result, test.expected)
		}
	}
}

func TestGetMin(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{input: []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}, expected: 1},
		{input: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, expected: 1},
		{input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, expected: 1},
		{input: []int{5, 5, 5, 5, 5}, expected: 5},
	}

	for _, test := range tests {
		result := GetMin(test.input)
		if result != test.expected {
			t.Errorf("GetMin(%v) = %v; expected %v", test.input, result, test.expected)
		}
	}
}

func TestGetSum(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{input: []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}, expected: 44},
		{input: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, expected: 55},
		{input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, expected: 55},
		{input: []int{5, 5, 5, 5, 5}, expected: 25},
	}

	for _, test := range tests {
		result := GetSum(test.input)
		if result != test.expected {
			t.Errorf("GetSum(%v) = %v; expected %v", test.input, result, test.expected)
		}
	}
}

func TestGetAverage(t *testing.T) {
	tests := []struct {
		input    []int
		expected float64
	}{
		{input: []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}, expected: 4.0},
		{input: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, expected: 5.5},
		{input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, expected: 5.5},
		{input: []int{5, 5, 5, 5, 5}, expected: 5.0},
	}

	for _, test := range tests {
		result := GetAverage(test.input)
		if result != test.expected {
			t.Errorf("GetAverage(%v) = %v; expected %v", test.input, result, test.expected)
		}
	}
}

func TestGetMedian(t *testing.T) {
	tests := []struct {
		input    []int
		expected float64
	}{
		{input: []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}, expected: 4.0},
		{input: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, expected: 5.5},
		{input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, expected: 5.5},
		{input: []int{5, 5, 5, 5, 5}, expected: 5.0},
	}

	for _, test := range tests {
		result := GetMedian(test.input)
		if result != test.expected {
			t.Errorf("GetMedian(%v) = %v; expected %v", test.input, result, test.expected)
		}
	}
}