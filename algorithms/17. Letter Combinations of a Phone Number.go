package algorithms

//Given a digit string, return all possible letter combinations that the number could represent.
//
//A mapping of digit to letters (just like on the telephone buttons) is given below.

var table = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	var result []string
	n(&result, digits, "", 0)
	return result
}

func n(result *[]string, digits, prefix string, index int) {
	if index >= len(digits) {
		if len(prefix) > 0 {
			*result = append(*result, prefix)
		}
		return
	}
	digit := digits[index] - '0'
	if digit < 2 {
		n(result, digits, prefix, index+1)
	} else {
		for _, v := range table[digit] {
			n(result, digits, prefix+string(v), index+1)
		}
	}
}
