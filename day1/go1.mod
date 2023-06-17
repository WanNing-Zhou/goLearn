// 模块名
module day1

//go sdk 版本

go 1.20

//当前module项目的依赖包

require (
	//dependency latest
)

//排除第三库

exclude (
	//dependency latest
)

//修改依赖包的路径或版本
//依赖包发生秦阿姨
//原始包无法访问
//使用本地包替换原始包
replace (
	//source latest => target latest
)

//撤回
//当前项目作为其他项目的依赖,如果某个版本出现了问题,则撤回该版本
retract (
	v1.0.0
)
