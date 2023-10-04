package main

import (
	"fmt"
)

func main() {
	var a = 30
	var b = 30
	if a == b {
		fmt.Println("这是真的")
	}

	var name string
	var age int
	// 键盘输入语句
	// 获取一行的数据
	fmt.Scanln(&name)
	fmt.Println(name)

	fmt.Scanf("姓名 %s 年龄 %d", &name, &age)
	fmt.Println(name, age)

}
