package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 1
	for  {
		a++
		if a > 3{ break}
		fmt.Println(a)
	}
	fmt.Println("over a:"+strconv.Itoa(a))
    /*
    比较  b和c的作用域  初始值的作用域
     */
	b := 2
	for ;b<3 ;b++  {
		fmt.Println(b)
	}

	fmt.Println(b)

	c := 0
	// 优化 c<len(a)  将len(a)保存于一个变量中
	for c:=1;c<3 ;c++  {
		fmt.Println(c)
	}
	fmt.Println(c)
}
