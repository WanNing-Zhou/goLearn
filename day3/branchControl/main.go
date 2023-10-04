package main

import "fmt"

func main() {
	// 流程控制语句
	// 分支控制 :
	// 语法: if <表达式 值为false | true >  {
	//	 执行代码
	//} else if <表达式> {
	// 执行代码
	//} ... else {
	// 执行代码
	//}

	var age int
	fmt.Println("请输入年龄")
	fmt.Scanln(&age)
	if age >= 18 {
		fmt.Println("您已经成年了")
	}

	// 支持定义变量
	if age2 := 20; age2 >= 18 {
		fmt.Println("您已经成年了")
	} else {
		fmt.Println("您还未成年")
	}

	// switch 分支结构
	// 语法: switch 不需要break
	/**
	switch <表达式> {

	case <表达式1>, <表达式2>, <...> :
		执行代码
	case <表达式1>, <表达式2>, <...> :
		执行代码
	default :
		执行代码
	}

	如果匹配就执行相应的case然后退出, 如果没有任何表达式匹配成功, 则执行default 语言块
	*/

	switch age {
	case 19, 20:
		fmt.Println("您已经成年了")
	case age, 21:
		fmt.Println("您说的对")
	default:
		fmt.Println("您没有成年")
	}

	// for 循环控制
	// 语法
	/**
	 for 初始化; 循环条件; 循环变量迭代 {
		执行代码
	}
	*/

	for i := 1; i < 10; i++ {
		fmt.Println(i)
		if i == 2 {
			goto label1
		}
	}

	// goto : 一般不使用goto, 容易造程序的混乱

label1:
	fmt.Println("kakxi ")

}
