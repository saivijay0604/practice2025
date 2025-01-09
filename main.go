package main

import (
	"fmt"
	routine "leetcode/practice2025/concurrencyExercise/goRoutine"
	mutexexe "leetcode/practice2025/concurrencyExercise/mutexProblems"
	config "leetcode/practice2025/config"
	algomonster "leetcode/practice2025/leetCodeExercise"
	integerLC "leetcode/practice2025/leetCodeExercise/integerLeetcode"
	sd "leetcode/practice2025/leetCodeExercise/samedirection"
	stringLC "leetcode/practice2025/leetCodeExercise/stringLeetcode"
)

const (
	multiplier = 2
	k          = 5
)

var linkedList = config.CreateLinkedList(config.GetLinkedList())
var root *config.Node

func main() {
	fmt.Println("Practice 2025 ")

	config.PrintLinkedList(linkedList)

	// fmt.Println("Index :", config.GetNums())
	sameDirection()

	leetcodeInts()
	leetcodeStrings()
	algoMonsterPractice()

	concurrencyProblems()

	mutexProblems()

}

func sameDirection() {
	sd.MoveZeroes(config.GetNums())

	resultRD := sd.RemoveDuplicates(config.GetNums())
	fmt.Println("\nThe number of unique elements in numsRD: ", resultRD)

	val := sd.MiddleNode(linkedList)
	fmt.Println("Middle of the list: ", val.Value)

}

func leetcodeStrings() {
	fmt.Println("\nString problems from leet Code")

	roman := "LVIII"
	total := stringLC.RomanToInt(roman)
	fmt.Println("\n Convert the roman number into Integer: ", total)

	// strArray := []string{"flower", "flow", "flight"}
	// commandPrefix := stringLC.LongestCommonPrefix(strArray)
	// fmt.Println("\n longest common prefix: ", commandPrefix)

}

func leetcodeInts() {
	fmt.Println("\nInteger problems from leet Code")

	x := integerLC.GetFinalState(config.GetNumInt(), k, multiplier)
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
