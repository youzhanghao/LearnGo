package main

import "fmt"

func main() {
	s1 := make([]int, 3, 10)
	// 切片大小 3 容量 10 当超过容量 则在上一次原有容量基础上扩容一倍，并重新分配内存条，这种情况下，效率则较低，
	// 因此若一开始大概知道 切片所占据容量是比较好的
	// 若不设置默认为 当前切片大小
	fmt.Println(len(s1),cap(s1))
	a := [6]int{1,2,3,4,5,6}
	s1 = a[4:]
	s2 := s1[1:2] // 重新reslice 不可超过原slice的切片边界 s1[1:3]
	fmt.Println(s1, len(s1),cap(s1))
	fmt.Println(s2,len(s2),cap(s2))

}
