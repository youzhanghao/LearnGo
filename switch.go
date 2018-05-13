package main

import "fmt"

func main() {
	a := 1
	switch a {
	case 0:
		fmt.Println("a = 0")
	case 1:
		fmt.Println("a = 1")
	default:
		fmt.Println("None")

	}

	b := 0
	switch  {
	case b >= 0:
		fmt.Println("b>=0")
		fallthrough // 满足该条件后继续往下执行
	case b <=1:
		fmt.Println("b<=1")
	case b>=1:
		fmt.Println("b>1")
	default:
		fmt.Println("b<0")
	}
}
