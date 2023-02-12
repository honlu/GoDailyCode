/*
数组
*/
package main

import "fmt"

func main() {
	var a = [5]int{1, 2, 3, 4, 5}
	var b = []int{1, 2, 3, 4, 5}
	fmt.Println(a)
	testArray(a)
	testSlice(b)
	fmt.Println(a)
	fmt.Println(b)
}

func testArray(a [5]int) {
	for i := 0; i < len(a); i++ {
		a[i] *= 2
	}
}

func testSlice(b []int) {
	for i := 0; i < len(b); i++ {
		b[i] *= 2
	}
}
