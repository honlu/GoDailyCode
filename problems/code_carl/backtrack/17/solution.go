package backtrack17

var digitLetters = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	res := []string{}
	path := make([]byte, 0, len(digits))
	var backtrack func(int)
	backtrack = func(index int) {
		if index == len(digits) {
			res = append(res, string(path))
			return
		}
		for i := 0; i < len(digitLetters[digits[index]]); i++ {
			path = append(path, digitLetters[digits[index]][i])
			backtrack(index + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return res
}
