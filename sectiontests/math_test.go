package sectiontests

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	type testCase struct {
		numbers  []int
		expected int
	}

	cases := []testCase{
		{[]int{1, 2}, 3},
		{[]int{2, 3}, 5},
		{[]int{-1, 2}, 1},
	}

	for _, val := range cases {
		x := sum(val.numbers...)
		if x != val.expected {
			t.Error("Expected", val.expected, "got", x)
		}
	}
}

func ExampleSum() {
	fmt.Println(sum(1, 2))
	// Output: 3
}
