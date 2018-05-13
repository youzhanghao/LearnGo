package main

import "fmt"

func main() {
	// 定义slice
	// slice速度最快
	var s1 []int
	fmt.Println(s1)

	a := [10]int{1,2,3,4,5,6,7,8,9,10}
	s2 := a[5:10] // [5,10)
	s3 := a[5]
	s4 := a[5:]
	s5 := a[:5]
	s6 := a[5:len(a)]
	fmt.Println(s2,s3,s4,s5,s6)
}
