package main

import "fmt"

func main() {
	var m map[int]map[int]string
	m = make(map[int]map[int]string)
	//m[1] = make(map[int]string)
	//m[1][1] = "ok"
	a, ok := m[2][1] // 第一个参数返回值  第二个参数返回代表是否存在
	// 不存在则初始化map  检查
	if !ok{
		m[2] = make(map[int]string)
	}
	m[2][1] = "Good"
	a, ok = m[2][1]
	fmt.Println(a)
}
