// 当前程序的包
package main
// 引入其他包  未调用，引入会抛错
import "fmt"

// 常量的定义
const PI = 3.14

// 组定义
const (
	const1 = "1"
	const2 = 2
	)

// 全局变量的声明与赋值
var name = "gopher"

// 一般类型声明
type newTpye int

// 结构的声明
type gopher struct{}

// 接口的声明
type golang interface {

}

// 由main函数作为程序入口  函数名小写 ==> private
func main(){
	 // 函数名大写 ==> public
	fmt.Println("hello go")
	//var b int = 1
	b := 1 // 由系统自动推断
	fmt.Println(b)
    // _ 省略
    a, _,c ,d := 2,3,4,5
    fmt.Println(a)
    fmt.Println(c)
    fmt.Println(d)
}

