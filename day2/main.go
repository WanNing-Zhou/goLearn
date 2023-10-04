package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a = 200
	fmt.Println(a)
	// 判断数据是什么类型的数据
	// fmt.Printf() 这个可以用来做格式化输出
	fmt.Printf("a 的类型 %T", a)
	// 在程序中查看变量的字节大小以及数据类型
	var a2 int = 200
	// unsafe.Sizeof() 可以返回变量占用的字节数
	fmt.Printf("a2 的 数据类型 %T 占用的字节数是 %d", a2, unsafe.Sizeof(a2))
	// golang程序中 整形变量 在使用时, 遵守保小不保大的原则, 在保证程序正常运行的情况下尽量使用占用空间小的数据类型
	// bit 在计算机中最小的存储单位, byte: 计算机中基础的存储单元

}
