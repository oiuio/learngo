package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div1(a, b) //不需要用到的返回值可以用"_"
		return q, nil
	default:
		//panic("unsupported operation: " + op)
		return 0, fmt.Errorf("unsupported operation: " + op)
	}
}

// 13/3 = 4....1
func div1(a, b int) (int, int) {
	return a / b, a % b
}

func div2(a, b int) (q, r int) {
	return a / b, a % b
}

func div3(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d,%d)", opName, a, b)
	return op(a, b)
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func swap(a, b *int) {
	*b, *a = *a, *b
}

func swap2(a, b int) (int, int) {
	return b, a
}

func main() {
	//fmt.Println(eval(3, 4, "x"))
	if result, err := fmt.Println(eval(3, 4, "x")); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(result)
	}

	fmt.Println(div1(13, 3))
	q, r := div2(13, 3)
	fmt.Println(q, r)
	fmt.Println(div3(13, 3))
	fmt.Println(
		apply(func(i int, i2 int) int {
			return i * i2
		}, 3, 4))
	fmt.Println(sum(1, 2, 3, 4, 5))

	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
	a, b = swap2(a, b)
	fmt.Println(a, b)
}
