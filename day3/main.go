package main

func main() {
	// 浮点分类: 单精度和双精度, float32 , float64
	// 浮点数 = 符号位 + 指数位 + 尾数位
	// 浮点数都是有符号的
	// 尾数部分可能精度缺失
	var num3 float32 = -123.0000901
	var num4 float64 = -123.0000901
	println(num3)
	println(num4)
	// float64 比 float32 精度更高一些
	// 要保存一个精度高的数, 应该选择float64
	// 浮点默认声明 float64

}
