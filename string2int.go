package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 65
	b := string(a)
	fmt.Println("b="+b)// 转换的65数字对应的A了
	c := strconv.Itoa(a) // 需要调用strconv方法
	fmt.Println("c="+c)
	var d,_ = strconv.Atoi(b)
	fmt.Println(d)
}
