package main

import "fmt"

func main() {
	a := 2
	// 块作用域
	if a:= 1; a > 0 {
		fmt.Println(a)
	}
	fmt.Println(a)
}
