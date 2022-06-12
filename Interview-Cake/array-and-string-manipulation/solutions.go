package main

import "fmt"

/*
Problem:
- Given a list of unsorted, independent meetings, returns a list of a merged
  one.
Example:
- Input: []meeting{{1, 2}, {2, 3}, {4, 8}, {5, 6}}
  Output: []meeting{{1, 3}, {4, 8}}
- Input: []meeting{{1, 5}, {2, 3}}
  Output: []meeting{{1, 5}}a

Solution:

Cost:
O(n^2)
*/

type Meeting struct {
	Start int
	End   int
}

func IsCrossed(l, r Meeting) bool {
	return (l.End >= r.Start && l.End <= r.End) ||
		(l.Start >= r.Start && l.Start <= r.End) ||
		(l.Start <= r.End && l.End >= r.End)
}

func megreMeetings(meetings []Meeting) []Meeting {
	var res []Meeting = make([]Meeting, 0)

	if len(meetings) == 0 {
		return res
	}

	res = append(res, meetings[0])
	for i := 1; i < len(meetings); i++ {
		for j := 0; j < len(res); j++ {
			if IsCrossed(meetings[i], res[j]) {

				if meetings[i].Start < res[j].Start {
					res[j].Start = meetings[i].Start
				}

				if meetings[i].End > res[j].End {

					// if latest crossed meeting is later
					if res[len(res)-1].End > meetings[i].End {
						res[j].End = res[len(res)-1].End
					} else {
						res[j].End = meetings[i].End
					}
					res = res[:j+1]
				}
				break
			}
			if j == len(res)-1 {
				if meetings[i].End < res[j].Start {
					res = append([]Meeting{meetings[i]}, res...)
				} else {
					res = append(res, meetings[i])
				}
				break
			}
		}
	}

	return res
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

var DIVIDER string = "_"

func Swap[T any](array []T, lIndex int, rIndex int) {
	var temp T = array[rIndex]
	array[rIndex] = array[lIndex]
	array[lIndex] = temp
}

func _ReverseStringArray(array []string, start int, end int) []string {
	for l, r := start, end-1; l < (start + (end-start)/2); l, r = l+1, r-1 {
		Swap(array, l, r)
	}
	return array
}

func ReverseStringArray(array []string) []string {
	return _ReverseStringArray(array, 0, len(array))
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
func ReverseWordsArray(array []string) []string {
	length := len(array)
	rightSortedIndex := length
	leftSortedIndex := 0

	for l, r := 0, len(array)-1; rightSortedIndex != leftSortedIndex; l, r = l+1, r-1 {
		Swap(array, l, r)
		if l == r {
			_ReverseStringArray(array, leftSortedIndex+1, rightSortedIndex)
			break
		}
		if array[l] == DIVIDER {
			_ReverseStringArray(array, leftSortedIndex, l)
			leftSortedIndex = l
		}
		if array[r] == DIVIDER {
			_ReverseStringArray(array, r+1, rightSortedIndex)
			rightSortedIndex = r
		}
	}
	return array
}

func main() {
	fmt.Println(megreMeetings([]Meeting{{5, 10}, {12, 13}, {14, 16}, {8, 15}, {1, 3}}))
}
