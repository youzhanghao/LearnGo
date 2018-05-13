package main

import "fmt"

func main() {

	s1 := []int{1,2,3,4,5,6}
	s2 := []int{7,8,9,0,1,4,0}
	copy(s2[1:4],s1[3:])
	s3 := make([]int, len(s1))
	copy(s3,s1)
	s4 := s3[:] // slice
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
}
