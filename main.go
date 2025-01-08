package main

import (
	"fmt"
	algomonster "leetcode/practice2025/algoMonster"
	routine "leetcode/practice2025/concurrencyExercise/goRoutine"
	mutexexe "leetcode/practice2025/concurrencyExercise/mutexProblems"
	"leetcode/practice2025/config"
	integerleetcode "leetcode/practice2025/integerLeetcode"
	"leetcode/practice2025/samedirection"
	stringleetcode "leetcode/practice2025/stringLeetcode"
)

const (
	multiplier = 2
	k          = 5
)

var linkedList = config.CreateLinkedList(config.GetLinkedList())
var root *config.Node

func main() {
	fmt.Println("Practice 2025 ")

	//config.PrintLinkedList(linkedList)

	// fmt.Println("Index :", config.GetNums())
	// sameDirection()

	// leetcodeInts()
	// leetcodeStrings()
	//algoMonsterPractice()

	//concurrencyProblems()

	mutexProblems()

}

func sameDirection() {
	samedirection.MoveZeroes(config.GetNums())

	resultRD := samedirection.RemoveDuplicates(config.GetNums())
	fmt.Println("\nThe number of unique elements in numsRD: ", resultRD)

	val := samedirection.MiddleNode(linkedList)
	fmt.Println("Middle of the list: ", val.Value)

}

func leetcodeStrings() {
	fmt.Println("\nString problems from leet Code")

	roman := "LVIII"
	total := stringleetcode.RomanToInt(roman)
	fmt.Println("\n Convert the roman number into Integer: ", total)

	// strArray := []string{"flower", "flow", "flight"}
	// commandPrefix := stringleetcode.LongestCommonPrefix(strArray)
	// fmt.Println("\n longest common prefix: ", commandPrefix)

}

func leetcodeInts() {
	fmt.Println("\nInteger problems from leet Code")

	x := integerleetcode.GetFinalState(config.GetNumInt(), k, multiplier)
	fmt.Println("\n Performs the K operations on the array: ", x)
}

func algoMonsterPractice() {
	// arr := []bool{false, false, true, true, true}
	// index := algomonster.FindBoundary(arr)
	// fmt.Println("The first true is at : ", index)

	//Create a Tree
	for _, v := range config.GetTree() {
		root = config.CreateTree(root, v)
	}
	config.InOrderTraversal(root)
	depth := algomonster.TreeMaxDepth(root)
	fmt.Println("Max depth of a binary tree : ", depth)

}

func concurrencyProblems() {
	fmt.Println("Concurrency ")
	//routine.GoRoutines()
	//routine.Waitgroup()
	//routine.ClosureExe2()
	//routine.UnBufferedChannel()
	//routine.BufferedChannel()
	//routine.ChannelDirection()
	//routine.ChannelOwnership()
	routine.SelectPractice()
}

func mutexProblems() {
	fmt.Println(" ** Mutex Practice **")
	mutexexe.MutexPractice()
	mutexexe.AtomicPractice()
	mutexexe.ConditionSignal()
	mutexexe.ConditionalBoardcast()

}
