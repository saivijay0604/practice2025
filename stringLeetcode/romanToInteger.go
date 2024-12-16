package stringleetcode

// Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000

// Input: s = "LVIII"
// Output: 58
// Explanation: L = 50, V= 5, III = 3.

var romanMap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func RomanToInt(s string) int {
	total := 0

	for i := 0; i < len(s); i++ {
		currentValue := romanMap[string(s[i])]
		if i < len(s)-1 && romanMap[string(s[i+1])] > currentValue {
			total -= currentValue
		} else {
			total += currentValue

		}
	}

	return total

}
