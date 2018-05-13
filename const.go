package main

import "fmt"



const(
	TEST_1 = iota
	cTEST_2  // 前缀加"c"或"_"不会被外部包所引用
	_TEST_3
)
func main() {
	fmt.Println(TEST_1)
	fmt.Println(cTEST_2)
	fmt.Println(_TEST_3)
}
