package main

import (
	"fmt"
	"leetcode/practice2025/config"
	integerleetcode "leetcode/practice2025/integerLeetcode"
	"leetcode/practice2025/samedirection"
	stringleetcode "leetcode/practice2025/stringLeetcode"
)

const (
	multiplier = 2
	k          = 5
)

func main() {
	fmt.Println("Practice 2025 ")

	fmt.Println("Index :", config.GetNums())
	sameDirection()

	leetcodeInts()
	leetcodeStrings()

}

func sameDirection() {
	samedirection.MoveZeroes(config.GetNums())

	resultRD := samedirection.RemoveDuplicates(config.GetNums())
	fmt.Println("\nThe number of unique elements in numsRD: ", resultRD)

	headLL := config.CreateLinkedList(config.GetMLL())
	config.PrintLinkedList(headLL)

	val := samedirection.MiddleNode(headLL)
	fmt.Println("Middle of the list: ", val.Value)

}

func leetcodeStrings() {
	fmt.Println("\nString problems from leet Code")

	r := "LVIII"
	total := stringleetcode.RomanToInt(r)
	fmt.Println("\n Convert the roman number into Integer: ", total)

}

func leetcodeInts() {
	fmt.Println("\nInteger problems from leet Code")

	x := integerleetcode.GetFinalState(config.GetNumInt(), k, multiplier)
	fmt.Println("\n Performs the K operations on the array: ", x)
}
