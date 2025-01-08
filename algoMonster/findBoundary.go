package algomonster

import (
	"fmt"
)

// /An array of boolean values is divided into two sections: The left section consists of all false, and the right section consists of all true.
// Find the First True in a Sorted Boolean Array of the right section, i.e., the index of the first true element.
// If there is no true element, return -1.

//Input: arr = [false, false, true, true, true]

//Output: 2

func FindBoundary(arr []bool) int {
	fmt.Println("Find the Boundary: ")
	if len(arr) == 0 {
		return -1
	}
	left := 0
	right := len(arr) - 1
	boundaryIndex := -1
	for left <= right {
		mid := (right + left) / 2
		if arr[mid] == false {
			left = mid + 1
		} else {
			boundaryIndex = mid
			right = mid - 1
		}
	}
	return boundaryIndex
}
