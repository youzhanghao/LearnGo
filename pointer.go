package main

import "fmt"

func main() {
	a := 1
	a++ // 只能作为语句 也就是只能放在变量的右边
	var p *int = &a
    fmt.Println(p,*p)
}