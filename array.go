package main

import "fmt"

func main() {
	var a [1]int
	var b [2]int

	// b = a  数据格式不一致  非法赋值
	fmt.Println(a,b)  // 数组默认值为0 字符串默认为空字符串

	c := [2]int{1} // 默认给第一个元素赋值
	d := [2]int{1,1}
	e := [20]int{18:1}  // int{索引:值}
	f := [...]int{1,3,6} // [...]表示由go计算该数组有多少个元素
	fmt.Println(c,d,e,f,len(e),len(f))

	// 指向数组的指针
	m := [...]int{99:1}
	var p *[100]int = &m
	fmt.Println(p)

	// 数组指针
	x,y := 1,2
	n := [...]*int{&x,&y}
	fmt.Println(n) // 保存地址

	j := [10]int{}
	k := new([10]int) // 使用new关键字操作
	j[1] = 1 // 索引操作
	fmt.Println(j,k)
}
