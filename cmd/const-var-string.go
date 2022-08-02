package main

import (
	"fmt"
	"os"
	"runtime"
)

//常量
func constFunc() {
	const a = "abc"         //隐式定义类型
	const b string = "abcd" //显式定义类型
	/*
		const cc int = Number(a) //定义的值必须是在编译期就能确定的值
		因为在编译期间自定义函数均属于未知，因此无法用于常量的赋值，但内置函数可以使用，如：len()
		当常量赋值给一个精度过小的数字型变量时，可能会因为无法正确表达常量所代表的数值而导致溢出
	*/
	const c int = len(a)

	//iota是一个预先声明的标识符 默认值为0
	//每当iota在新的一行使用时，它的值就会自动+1
	//没有赋值的常量会自动应用上一行的赋值
	//每遇到一个新的常量块或单个常量声明时， iota 都会重置为 0（ 简单地讲，每遇到一次 const 关键字，iota 就重置为 0 ）
	//eg.
	const (
		aa = iota //0
		bb        //1
		cc        //2
		dd = 5    //5
		ee        //5
	)
	println(aa)
	println(bb)
	println(cc)
	println(dd)
	println(ee)

	// 赋值两个常量，iota 只会增长一次，而不会因为使用了两次就增长两次
	const (
		aaa, bbb = iota + 1, iota + 2 // aaa=1 bbb=2
		ccc, ddd                      // ccc=2 ddd=3
		eee, fff                      // ddd=3 fff=4
	)

}

//变量
func varFunc() {
	/*
		当一个变量被声明之后，系统自动赋予它该类型的零值：
		int 为 0，
		float32(64) 为 0.0，
		bool 为 false，
		string 为空字符串，
		指针为 nil。
		记住，所有的内存在 Go 中都是经过初始化的

		变量的命名规则遵循骆驼命名法，即首个单词小写，每个新单词的首字母大写 eg. unmShips

		var a int = 15
		var (
				a int
				b bool
				c string
			)
	*/

	/*
		变量的作用域
		1. 当一个变量声明在函数体之外称之为全局变量
		2. 声明在函数体之内称之为局部变量，参数和返回值也是局部变量
	*/

	var a = 5
	print(a)

	//声明包级别的全局变量 语法 :=
	println("声明包级别的全局变量 语法 :=")
	b := 1
	println(b)
	getGoos()
}

//值类型引用类型
func valueRefFunc() {

	var a int = 3
	//打印内存地址
	fmt.Println(&a) //0xc0000160a8 16进制表示
}

//
func stringFunc() {

}

func Number(a string) int {
	return len(a)
}

func getGoos() {
	var goos string = runtime.GOOS
	fmt.Printf("The operating system is : %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
}
