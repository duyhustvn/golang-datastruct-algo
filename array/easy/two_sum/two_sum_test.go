package twosum

import (
	"reflect"
	"testing"
)

type TestData struct {
	Nums          []int
	Sum           int
	ExpectedIndex []int
}

func generateTestCase() []TestData {
	return []TestData{
		{
			Nums:          []int{1, 3, 7, 9, 2},
			Sum:           11,
			ExpectedIndex: []int{3, 4},
		},
		{
			Nums:          []int{1, 3, 7, 9, 2},
			Sum:           6,
			ExpectedIndex: nil,
		},
		{
			Nums:          []int{5},
			Sum:           5,
			ExpectedIndex: nil,
		},
		{
			Nums:          []int{},
			Sum:           5,
			ExpectedIndex: nil,
		},
	}
}

func TestTwoSum(t *testing.T) {
	tests := generateTestCase()
	for _, test := range tests {
		foundIndex := twoSum(test.Nums, test.Sum)
		if !reflect.DeepEqual(foundIndex, test.ExpectedIndex) {
			t.Errorf("Expected: %+v instead of %+v", test.ExpectedIndex, foundIndex)
		}
	}
}

func TestTwoSumUsingMap(t *testing.T) {
	tests := generateTestCase()
	for _, test := range tests {
		foundIndex := twoSumUsingMap(test.Nums, test.Sum)
		if !reflect.DeepEqual(foundIndex, test.ExpectedIndex) {
			t.Errorf("Expected: %+v instead of %+v", test.ExpectedIndex, foundIndex)
		}
	}
}
