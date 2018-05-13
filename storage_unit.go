package main

import "fmt"

// 利用iota和位运算实现计算机存储单位
const (
	B float64 = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB

)


func main() {
	fmt.Println(B,KB,MB)
}
