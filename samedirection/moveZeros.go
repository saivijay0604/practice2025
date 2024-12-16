package samedirection

import "fmt"

func MoveZeroes(nums []int) {
	fmt.Print("\nMove zeros at the end of the list: ")
	z := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[z] = nums[i]
			z++
		}
	}

	for i := z; i < len(nums); i++ {
		nums[i] = 0
	}
	fmt.Println(nums)
}
