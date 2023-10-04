package utils

// Add 用于计算两数之和
// 首字母需要大写, 表示该函数可导出
func Add(a int, b int) int {
	return a + b
}

var Age int
var Name string

// 如果引入的文件 和main 中都有init函数, 它的执行流程是
// 被引用变量定义 , 被引用init , main中变量定义 ,main中的init main中 main函数

func init() {
	Age = 18
	Name = "张三"
}
