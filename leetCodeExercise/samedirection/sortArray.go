package leetcode

import (
	"fmt"
	"sort"
)

func SortArray(array interface{}) (interface{}, error) {
	switch a := array.(type) {
	case []int:
		fmt.Println("Type Int: ")
		sort.Ints(a)
		return a, nil
	case []string:
		fmt.Println("Type string: ")
		sort.Strings(a)
		return a, nil
	default:
		return nil, fmt.Errorf("error - does not type")
	}

}

//alternative funtion without sort function
// func Sorta(a interfacte{})interface{} {
// 	for i:=0; i<len(a)-1;i++{
// 		for j:=0;j<len(a)-1-i;j++{
// 			if a[j] > a[j+1]{
// 				a[j],a[j+1] =a[j+1],a[j]
// 			}

// 		}
// 	}
//     return a
// }
