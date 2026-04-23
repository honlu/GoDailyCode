package main

// LCR 190. 加密运算
// 不使用加减乘除实现加法运算
// a ^ b 异或计算，表示无进位的加法
// （a & b）<< 1  表示进位
func encryptionCalculate(dataA int, dataB int) int {
	// 递归版
	// if dataB == 0 {
	// 	return dataA
	// }
	// return encryptionCalculate(dataA^dataB, (dataA&dataB)<<1)

	// 循环版
	for dataB != 0 {
		dataA, dataB = dataA^dataB, (dataA&dataB)<<1
	}
	return dataA
}
