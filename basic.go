package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

//定义变量
func variableZeroValue() {
	var a int;
	var s string;
	fmt.Printf("%d %q\n", a, s) //%q 空串会输出引号
}

//定义变量时赋初值
func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

//类型推断，不同类型也能写在一行
func variableTypeDeduction() {
	var a, b, s = 3, 4, "abc"
	fmt.Println(a, b, s)
}

//:=定义变量，变量第一次出现需要:定义，var必须用在函数内
func variableShorter() {
	a, b, s := 3, 4, "abc"
	b = 5
	fmt.Println(a, b, s)
}

//函数外定义变量，不能用:=，不是全局变量，二是包内部变量。没有全局变量说法
var aa = 3
var ss = "kkk"
//bb:= true
var bb = true
// 简略
var (
	aaa = 3
	sss = "kkk"
	bbb = true
)

//复数：欧拉公式
func euler() {
	//c:=3+4i
	//fmt.Println(cmplx.Abs(c)) //绝对值 = 5
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	fmt.Println(cmplx.Exp(1i+math.Pi) + 1)
	fmt.Printf("%3f", cmplx.Pow(math.E, 1i*math.Pi)+1)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func main() {
	fmt.Println("Hello World")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()

	euler()
	triangle()
}
