package leetcode

// Write a function to find the longest common prefix string amongst an array of strings.

// If there is no common prefix, return an empty string ""

// Input: strs = ["flower","flow","flight"]
// Output: "fl"

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 1; i < len(strs); i++ {
		for len(strs[0]) > 0 && strs[i][:len(strs[0])] != strs[0] {
			strs[0] = strs[0][:len(strs[0])-1]
		}

		if len(strs[0]) == 0 {
			return ""
		}

	}
	return strs[0]

}
