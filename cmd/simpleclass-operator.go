package main

import (
	"fmt"
	"gostu/configs"
	"math"
)

/*
	Go语言是强类型语言，不存在隐式转换，也不存在运算符的重载，表达式的解析是从左至右。
	只有两个相同类型的值才可以进行比较。

	基本类型与运算符
	基本类型：
	布尔值：bool
	数值：int,float32,float64 注：GO语言中没有float类型，也没有double类型

	int8（-128 -> 127）
	int16（-32768 -> 32767）
	int32（-2,147,483,648 -> 2,147,483,647）
	int64（-9,223,372,036,854,775,808 -> 9,223,372,036,854,775,807）

	float32 精确到小数点后 7 位，float64 精确到小数点后 15 位

	***格式化说明符***
	%t：布尔值
	%d: 整数（%x和%X用于格式化16进制表示的数字）
	%g: 浮点型（%f输出浮点数，%e输出科学计数法数）
	%0nd: 用于输出长度为n的整数，0开头的数字是必须的
	%n.mg: 表示数字n并精确到小数点后n位，也可以用 e,f %n.me,%n.mf
	%v 表示复数
	%b 是用于表示位的格式化标识符

*/

//类型转换
func typeConversion() {
	var a int16 = 5
	var b int32
	b = int32(a) //类型转换
	fmt.Printf("b is :%d\n", b)

	c := int64(0) //同时完成转换和赋值的操作
	fmt.Printf("c is :%d\n", c)

	fmt_util.FmtD("dadadad", 1)

}

// Unit8FromInt int 转 unit8
func Unit8FromInt(n int) (uint8, error) {
	fmt.Printf("maxuint8 is %d\n", 1<<8-1) //255
	if 0 <= n && n <= math.MaxUint8 {
		return uint8(n), nil
	}
	return 0, fmt.Errorf("%d is out of uint8 range", n)
}

// IntFromFloat64 float64 转 int
func IntFromFloat64(x float64) int {
	if math.MinInt32 <= x && x <= math.MaxInt32 { // x lies in the integer range
		whole, fraction := math.Modf(x)
		if fraction >= 0.5 {
			whole++
		}
		return int(whole)
	}
	panic(fmt.Sprintf("%g is out of the int32 range", x))
}
