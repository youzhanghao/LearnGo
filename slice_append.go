package main

import "fmt"

func main() {
	/**
	当append超过slice容量时，内存将会重新分配地址，并拷贝原有的值
	 */
	s1 := make([]int, 3 ,6)
	fmt.Printf("%p\n",s1)
	s1 = append(s1, 1,2,3)
	fmt.Printf("%v %p\n",s1, s1)
	s1 = append(s1, 4,5,6)
	fmt.Printf("%v %p\n",s1, s1)

	/**
	多个slice应用时需要注意的
	 */
	 a := []int{1,2,3,4,5,6}
	 s3 := a[2:5]
	 s2 := a[1:3]
	 fmt.Println(s3,s2)
	 //s2 = append(s2, 1,2,1,1,1)  // TODO 解开注释查看效果 内存已经被重新分配
	 s3[0] = 9
	 fmt.Println(s3,s2)
}
