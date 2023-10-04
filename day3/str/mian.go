package main

import "fmt"

func main() {
	// 字符串就是一串固定长度的字符连接起来的字符序列, go 的字符串由单个字节连接起来的, 也就是说对于传统的字符串是由字符组成的
	// go的字符串不同, 它是由字节组成的

	var c1 = '0'
	var c2 = 'k'
	var c3 = '卡'
	fmt.Println(c1, c2, c3)
	fmt.Printf("c1: %T, c2: %T, c3: %T", c1, c2, c3)

	var name = "卡卡西"
	fmt.Println('\n', name)
}
