/*
数组
*/
package main

import "fmt"

func main() {
	var a = [5]int{1, 2, 3, 4, 5}
	var b = []int{1, 2, 3, 4, 5}
	var c []int
	var d [5]int
	c = b
	d = a
	fmt.Println(a)
	testArray(a)
	testSlice(b)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
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
