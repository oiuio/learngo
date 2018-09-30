package main

import (
	"io/ioutil"
	"fmt"
)

//if不需要括号
func if1() {
	const filename = "abc.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

//if条件里可以赋值
//if条件里赋值的变量作用域仅在if语句里
func if2() {
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	//fmt.Printf("%s\n", contents)
}

//switch会自动break（默认有），除非使用fallthrough
func eval(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	default:
		panic("unsupported op: " + op)
	}
	return result
}

//
func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {

	if1()
	if2()

	fmt.Println(eval(10, 20, "+"))
	fmt.Println(eval(10, 20, "-"))
	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(80),
		grade(90),
		grade(100),
		//grade(101),//中断程序执行，报错
	)
}
