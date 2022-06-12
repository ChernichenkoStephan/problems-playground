package main

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestMergeMeetings(t *testing.T) {
	tests := []struct {
		in       []Meeting
		expected []Meeting
	}{
		{[]Meeting{}, []Meeting{}},
		{[]Meeting{{1, 2}}, []Meeting{{1, 2}}},
		{[]Meeting{{1, 2}, {2, 3}}, []Meeting{{1, 3}}},
		{[]Meeting{{1, 5}, {2, 3}}, []Meeting{{1, 5}}},
		{[]Meeting{{1, 2}, {4, 5}}, []Meeting{{1, 2}, {4, 5}}},
		{[]Meeting{{4, 5}, {1, 5}, {2, 3}}, []Meeting{{1, 5}}},
		{[]Meeting{{1, 2}, {2, 3}, {4, 5}}, []Meeting{{1, 3}, {4, 5}}},
		{[]Meeting{{1, 6}, {2, 3}, {4, 5}}, []Meeting{{1, 6}}},
		{[]Meeting{{4, 5}, {2, 3}, {1, 6}}, []Meeting{{1, 6}}},
		{[]Meeting{{5, 10}, {12, 13}, {14, 16}, {8, 15}, {1, 3}}, []Meeting{{1, 3}, {5, 16}}},
	}

	for _, tt := range tests {
		result := megreMeetings(tt.in)
		common.Equal(t, tt.expected, result)
	}
}

/*
Problem:
- Given a list of string, reverse its order in place.
Example:
- Input: []string{"a", "b", "c", "d"}
  Output: []string{"d", "c", "b", "a"}

Solution:

Cost:
O(n)
*/
func TestReverseStringArray(t *testing.T) {
	tests := []struct {
		in       []string
		expected []string
	}{
		{[]string{}, []string{}},
		{[]string{"a"}, []string{"a"}},
		{[]string{"a", "b"}, []string{"b", "a"}},
		{[]string{"a", "b", "c"}, []string{"c", "b", "a"}},
		{[]string{"a", "b", "c", "d"}, []string{"d", "c", "b", "a"}},
	}

	for _, tt := range tests {
		result := ReverseStringArray(tt.in)
		common.Equal(t, tt.expected, result)
	}
}

/*
Problem:
- Given a list of string that is made up of word but in reverse, return the
  correct order in-place.
Example:
- Input: []string{"w", "o", "r", "l", "d", "", "h", "e", "l", "l", "o", "", "s", "a", "y"}
  Output: []string{"s", "a", "y", "", "h", "e", "l", "l", "o", "", "w", "o", "r", "l", "d"}
Solution:
Cost:
*/
func TestReverseWord(t *testing.T) {
	tests := []struct {
		in       []string
		expected []string
	}{
		{
			[]string{"w", "o", "r", "l", "d", DIVIDER, "h", "e", "l", "l", "o", DIVIDER, "s", "a", "y"},
			[]string{"s", "a", "y", DIVIDER, "h", "e", "l", "l", "o", DIVIDER, "w", "o", "r", "l", "d"},
		},
	}

	for _, tt := range tests {
		result := ReverseWordsArray(tt.in)
		common.Equal(t, tt.expected, result)
	}
}
