package main

import (
	fm "fmt"
	testDemo "gostu/cmd/testdemo"
)

func main() {
	// 字符串 基本类型和运算符 字符串 strings strconv包 时间和日期 指针

	//常量
	constFunc()

	//变量
	varFunc()

	//值类型引用类型
	//valueRefFunc()

	//typeConversion()

	Unit8FromInt(256)

	//测试导入其他目录的包
	testDemo.Test()
}

// init 初始化函数
func init() {
	fm.Println("初始化函数先执行")
}
