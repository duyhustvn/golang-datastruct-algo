package fib

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type DataTest struct {
	num    int
	expOut int
}

// var tests []DataTest
// tests = []DataTest{
// 	{1, 1},
// 	{2, 1},
// 	{3, 2},
// 	{4, 3},
// }

var tests = []struct {
	num    int
	expOut int
}{
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
}

func TestFib(t *testing.T) {
	for _, test := range tests {
		actualOutput := fib(test.num)
		assert.Equal(t, test.expOut, actualOutput)
	}
}

var result int

func BenchmarkFib(b *testing.B) {
	var r int
	for _, v := range tests {
		b.Run(fmt.Sprintf("input_size_%d", v.num), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// always record the result of Fib to prevent
				// the compiler eliminating the function call.
				r = fib(v.num)
			}
		})
	}

	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}
