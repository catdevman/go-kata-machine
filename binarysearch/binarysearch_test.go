package binarysearch

import (
	"testing"
)

var intTestCases = []struct {
	name   string
	input  []int
	target int
	output int
}{
	{
		"element found at left hand side",
		[]int{1, 2, 3, 4, 5},
		2,
		1,
	},
	{
		"element found at right hand side",
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		7,
		6,
	},
	{
		"element not found",
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		10,
		-1,
	},
	{
		"element found at the middle",
		[]int{1, 2, 3, 4, 5},
		3,
		2,
	},
	{
		"searching in an empty slice",
		[]int{},
		3,
		-1,
	},
	{
		"negative",
		[]int{-9,-8,-7,-6,-5,-4, -3, -2, -1},
		-8,
		1,
	},
}

var stringTestCases = []struct{
    name string
    input []string
    target string
    output int
}{
    {
        "element found a left hand side",
        []string{"a", "b", "c", "d", "e"},
        "b",
        1,
    },
}

var int64TestCases = []struct {
	name   string
	input  []int64
	target int64
	output int
}{
	{
		"element found at left hand side",
		[]int64{1, 2, 3, 4, 5},
		2,
		1,
	},
	{
		"element found at right hand side",
		[]int64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		7,
		6,
	},
	{
		"element not found",
		[]int64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		10,
		-1,
	},
	{
		"element found at the middle",
		[]int64{1, 2, 3, 4, 5},
		3,
		2,
	},
	{
		"searching in an empty slice",
		[]int64{},
		3,
		-1,
	},
	{
		"negative",
		[]int64{-9,-8,-7,-6,-5,-4, -3, -2, -1},
		-8,
		1,
	},
}
func TestIntBinarySearch(t *testing.T) {
	for _, tc := range intTestCases {
		t.Run(tc.name, func(t *testing.T) {
			want := tc.output
			got := BinarySearch(tc.input, tc.target)
			if got != want {
				t.Errorf("got: %v, want: %v", got, want)
			}
		})
	}
}

func TestInt64BinarySearch(t *testing.T){
	for _, tc := range int64TestCases {
		t.Run(tc.name, func(t *testing.T) {
			want := tc.output
			got := BinarySearch(tc.input, tc.target)
			if got != want {
				t.Errorf("got: %v, want: %v", got, want)
			}
		})
	}
}

func TestStringBinarySearch(t *testing.T){
	for _, tc := range stringTestCases {
		t.Run(tc.name, func(t *testing.T) {
			want := tc.output
			got := BinarySearch(tc.input, tc.target)
			if got != want {
				t.Errorf("got: %v, want: %v", got, want)
			}
		})
	}
}
