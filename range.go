package main

import "fmt"

func main() {
	sm := make([]map[int]string, 5)
	for _, v := range sm {
		v = make(map[int]string, 1)
		v[1] = "ok" // value值只是作为拷贝
		fmt.Println(v)
	}
	fmt.Println(sm)

	for i,_ := range sm{
		sm[i] = make(map[int]string, 1)
		sm[i][1] = "ok" // 对索引操作
		fmt.Println(sm[i])
	}
	fmt.Println(sm)
}
