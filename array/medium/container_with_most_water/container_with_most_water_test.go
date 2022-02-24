package containerwithmostwater

import "testing"

type TestData struct {
	Input        []int `json:"input"`
	ExpectedArea int   `json:"expected_area"`
}

func TestFindContainerWithMostWater(t *testing.T) {
	tests := []TestData{
		{
			Input:        []int{},
			ExpectedArea: 0,
		},
		{
			Input:        []int{1},
			ExpectedArea: 0,
		},
		{
			Input:        []int{1, 8, 6, 2, 9, 4},
			ExpectedArea: 24,
		},
		{
			Input:        []int{1, 8, 6, 2, 9, 10},
			ExpectedArea: 32,
		},
		{
			Input:        []int{6, 9, 3, 4, 5, 8},
			ExpectedArea: 32,
		},
	}
	for _, test := range tests {
		realOutput := findContainerWithMostWater(test.Input)
		if realOutput != test.ExpectedArea {
			t.Errorf("Expected %d got %d", test.ExpectedArea, realOutput)
		}
	}
}
