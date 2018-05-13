package main

import "fmt"

/*
  位运算
  6: 0110
  11: 1011
------------
& 或 两个都为1才为1
| 与 其中一个为1才为1
^    只有一个1才为1
&^  0100 第二个数为1则强制将第一个数改为0
 */
func main() {
	fmt.Println(6 &^ 11)
}