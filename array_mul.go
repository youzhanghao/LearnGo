package main

import "fmt"

func main() {
	// 冒泡排序
	a := [...]int{2,4,2,6,8,9}
	fmt.Println(a)

	num := len(a)
	for i:=0;i<num ; i++ {
		for j:=i+1;j<num ; j++ {
			if a[i] < a[j]{
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}

	}
	fmt.Println(a)
}