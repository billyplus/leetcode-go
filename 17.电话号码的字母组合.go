package leetcode

/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
func letterCombinations(digits string) []string {
	result := combination([]byte(digits))
	ret := make([]string, 0, len(result))
	for _, r := range result {
		ret = append(ret, string(r))
	}
	return ret
}

func combination(digits []byte) [][]byte {
	result := [][]byte{}
	n := len(digits)
	if n == 0 {
		return result
	}
	d := digits[0] - 0x30
	al := data[d]
	if n == 1 {
		for j := 0; j < len(al); j++ {
			result = append(result, []byte{al[j]})
		}
	} else {
		left := combination(digits[1:])
		for j := 0; j < len(al); j++ {
			for _, b := range left {
				result = append(result, append([]byte{al[j]}, b...))
			}
		}
	}
	return result
}

var (
	data = []string{
		"", "", "abc", "def",
		"ghi", "jkl", "mno",
		"pqrs", "tuv", "wxyz",
	}
)

// @lc code=end
