package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[int]string{1:"a",3:"c",2:"b"}
	fmt.Println(m) // 输出map的顺序是无序的
	s := make([]int,len(m))
	fmt.Println(s)
	i := 0

	// key 顺序
	for k,_ := range m{
		s[i] = k // 无序取k
		i++
	}
	fmt.Println(s)
	sort.Ints(s)
	fmt.Println(s)

	// 键值交换
	n := make(map[string]int,len(m))
	for k,v := range m{
		n[v] = k
	}
	fmt.Println(n)
	//for _, _ = range m {
	//	s[i] = m[i+1]
	//	i++
	//}
	//fmt.Println(s)
}
