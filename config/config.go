package config

var (
	nums    = []int{0, 0, 1, 3, 1, 2, 2, 3, 3, 4}
	numsLL  = []int{6, 7, 8, 3, 12, 2}
	numsInt = []int{2, 1, 3, 5, 6}
	numTree = []int{15, 10, 20, 8, 12, 17, 25}
)

func GetNums() []int {
	return nums
}

func GetLinkedList() []int {
	return numsLL
}

func GetNumInt() []int {
	return numsInt
}

func GetTree() []int {
	return numTree
}
