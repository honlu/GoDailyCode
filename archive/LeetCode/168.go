/*
168. excel表列名称
2023-3-31 by lu
link:https://leetcode.cn/problems/excel-sheet-column-title/
question:
	给你一个整数 columnNumber ，返回它在 Excel 表中相对应的列名称。
answer:
	关于计算过程细节要掌握住。
*/
func convertToTitle(columnNumber int) string {
	res := []byte{}
	for columnNumber > 0 {
		columnNumber--
		res = append(res, 'A'+byte(columnNumber%26))
		columnNumber /= 26
	}
	reverse(res)
	return string(res)
}

func reverse(arr []byte) {
	size := len(arr)
	for i := 0; i < size/2; i++ {
		arr[i], arr[size-1-i] = arr[size-1-i], arr[i]
	}
}