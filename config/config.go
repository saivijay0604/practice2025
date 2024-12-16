package config

var (
	nums    = []int{0, 0, 1, 3, 1, 2, 2, 3, 3, 4}
	numsMLL = []int{6, 7, 8, 3, 12, 2}
	numsInt = []int{2, 1, 3, 5, 6}
)

func GetNums() []int {
	return nums
}

func GetMLL() []int {
	return numsMLL
}

func GetNumInt() []int {
	return numsInt
}
