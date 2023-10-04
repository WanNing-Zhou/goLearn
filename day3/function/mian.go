package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

var (
	// 函数中-defer
	// 在函数中, 程序员经常创建资源 (比如: 数据库连接, 文件句柄等, 锁等), 为了函数执行完毕后, 及时释放资源, go的设计者提供defer(延时机制)

	fun1 = func(n1 int, n2 int) int {
		// 当执行到defer时, 会将defer后面的语句压入到独立的栈中(defer栈中), 暂时不执行, 按照先入后出的方式执行
		// 当执行完毕后, 再从defer栈中, 依次出栈执行 (注意先入后出机制)
		// 再defer 将语句放入到栈时, 也会将相关的值拷贝同时入栈
		defer fmt.Println("n1", n1)
		defer fmt.Println("n2", n2)
		n1 += 30
		fmt.Println("n1 + n2", n1+n2)
		return n1 + n2
	}
)

func main() {

	fmt.Println(fun1(3, 44))

	var a int
	var b int
	fmt.Println("请输入两个数")
	fmt.Scanln(&a, &b)
	fmt.Println("两个数的和值为", add(a, b))

	// 匿名函数的使用
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(1, 2)
	fmt.Println(res1)

	fun := func(n1 int, n2 int) int {
		return n1 - n2
	}

	res2 := fun(20, 10)

	fmt.Println(res2)

}
