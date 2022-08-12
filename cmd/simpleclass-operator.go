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

	***位运算符***

	按位与 &
	二者都为1 才为1 其他情况都是0
	1 & 1 -> 1
	1 & 0 -> 0
	0 & 1 -> 0
	0 & 0 -> 0

	按位或 |
	其中一个为1 则为1 都是0 则为0
	1 | 1 -> 1
  	1 | 0 -> 1
  	0 | 1 -> 1
  	0 | 0 -> 0

	按位异或 ^
	二者同为1 或者同为0 则为0
	1 ^ 1 -> 0
	1 ^ 0 -> 1
	0 ^ 1 -> 1
	0 ^ 0 -> 0

	位清除 &^：将指定位置上的值设置为 0
	15 &^ 4
	15表示 1111
	4表示  0100
	4转换为二进制时，将位置为1对应15的位置设置为0，剩余的位置与15的位置相同 得到  1011
	所谓的位清空操作，把4中1对应位置在15位置上清空

	***算数运算符***
	Go 在进行字符串拼接时允许使用对运算符 + 的重载
	Go 本身不允许开发者进行自定义的运算符重载
	取余运算符只能作用于整数：9 % 4 -> 1
	浮点数除以 0.0 会返回一个无穷尽的结果，使用 +Inf 表示
	i++ i-- :只能用于后缀，带有++或者--的只能用于语句，而非表达式
	i++ -> i += 1 -> i = i + 1
	i-- -> i -= 1 -> i = i - 1
	在运算时 溢出 不会产生错误，Go 会简单地将超出位数抛弃，如果需要范围大的数 使用big包

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

//算数运算符
func numOperator() {
	i := 1
	i++

}
