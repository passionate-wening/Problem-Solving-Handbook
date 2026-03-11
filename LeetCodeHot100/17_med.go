package LeetCodeHot100

func letterCombinations(digits string) []string {
	var res []string
	array := make([]string, 8)
	array = []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	for i := 0; i < len(digits); i++ {
		num := digits[i] - '0' - 2
		if len(res) == 0 {
			for j := 0; j < len(array[num]); j++ {
				res = append(res, array[num][j:j+1])
			}
		} else {
			var newRes []string
			for _, s := range res {
				for j := 0; j < len(array[num]); j++ {
					newRes = append(newRes, s+array[num][j:j+1])
				}
			}
			res = newRes
		}
	}
	return res
}
