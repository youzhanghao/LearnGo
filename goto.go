package main

import "fmt"

func main() {
LABEL1:
	for{
		for i:=1;i<10 ;i++  {
			fmt.Println("hello")
			if i>3 {
				//goto LABEL1 // 仍旧是跳到外层循环 无限循环  goto调整循环层次
				goto LABEL2
				break LABEL1 // 不使用label则只跳出当前层的循环，外层还是无限循环，使用标签则跳至与标签同级的循环
				// 若是continue 则走有限循环
			}
		}
	}
LABEL2:
	fmt.Println("ok")
}
