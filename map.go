package main

import "fmt"

func main() {
	//var m map[int]string = make(map[int]string)
	m := make(map[int]string)
	m[1] = "ok"
	fmt.Println(m,m[1])
	delete(m,1)
	fmt.Println(m)
}
