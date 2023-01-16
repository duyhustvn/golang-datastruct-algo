package subsequence

import (
	"testing"
	"reflect"
	"github.com/stretchr/testify/assert"
	"fmt"
)


// the order is not important
func compare2DSlices(sl1 [][]int, sl2 [][]int) bool {
	if len(sl1) != len(sl2) {
		return false;
	}

	for _, sub1 := range(sl1) {
		isFound := false
		for j, sub2 := range (sl2) {
			if reflect.DeepEqual(sub1, sub2) {
				// remove sub2 from sl2
				copy(sl2[j:], sl2[j+1:])
				sl2 = sl2[:len(sl2)-1]
				isFound = true
				break
			}
		}
		if !isFound {
			return false
		}
	}

	return true
}

func print2DSlice(sl [][]int) {
	for _, sub := range(sl) {
		for _, e := range(sub) {
			fmt.Printf("%d ", e)
		}
		fmt.Println()
	}
}

func TestGenerateAllSubsequenceRecursive(t *testing.T) {
	type DataTest struct {
		input 		[]int
		expectedSubsequences [][]int
	}

	tests := []DataTest{
		{
			[]int{1,2}, [][]int{{1,2},{}, {2}, {1}},
		},
		{
			[]int{1,2,2}, [][]int{ {1},{1,2},{1,2,2},{1,2},{2,2},{2},{2},{} },
		},
		{
			[]int{1,2,3}, [][]int{ {1,2,3},{1,2},{1,3},{1},{2,3},{2},{3},{} },
		},
	}

	for _, test := range(tests) {
		realSubsequences := generateSubsequence(test.input)
		
		assert.Equal(t, compare2DSlices(realSubsequences, test.expectedSubsequences), true)
	}
}

func TestSubsequenceWithFixedLength(t *testing.T) {
	type DataTest struct {
		input 		[]int
		length      int
		expectedSubsequences [][]int
	}

	tests := []DataTest{
		{
			[]int{1,2}, 2, [][]int{{1,2}},
		},
		{
			[]int{1,2}, 1, [][]int{{1}, {2}},
		},
		{
			[]int{1,2,3}, 1, [][]int{ {1},{2},{3} },
		},
		{
			[]int{1,2,3}, 2, [][]int{ {1,2},{1,3},{2,3} },
		},
		{
			[]int{1,2,3}, 3, [][]int{ {1,2, 3} },
		},
		{
			[]int{1,2,2}, 1, [][]int{ {1},{2},{2} },
		},
		{
			[]int{1,2,2}, 2, [][]int{ {1,2},{1,2},{2,2} },
		},
		{
			[]int{1,2,2}, 3, [][]int{ {1,2,2} },
		},
	}

	for _, test := range(tests) {
		realSubsequences := generateSubsequenceWithFixedLength(test.input, test.length)

		fmt.Println("Expected: ")
		print2DSlice(test.expectedSubsequences)
		fmt.Println("Real: ")
		print2DSlice(realSubsequences)

		assert.Equal(t, compare2DSlices(realSubsequences, test.expectedSubsequences), true)
	}
}
