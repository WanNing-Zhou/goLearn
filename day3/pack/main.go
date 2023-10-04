package main

import (
	"fmt"
	"goLearn/day3/pack/utils"
)

// 每一个源文件都可以包含一个init 函数 , 该函数会在main函数执行前, 被go运行框架调用, 也就是说init 会在main之前被调用
// 通常用来做初始化操作

// 如果一个文件中包含全局变量定义 , 同时包含init , main 先执行全局变量定义 => init => main

var age int = test()

func test() int {
	fmt.Println("kakxi")
	return 90
}
func init() {
	fmt.Println("init")
}

func main() {
	var a int
	var b int
	fmt.Println("请输入两个数")
	fmt.Scanln(&a, &b)
	fmt.Println("两数之和", utils.Add(a, b))
}
