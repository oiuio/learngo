package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6] = ", arr[2:6]) //arr[2:6] =  [2 3 4 5]
	fmt.Println("arr[:6] = ", arr[:6])   //arr[:6] =  [0 1 2 3 4 5]
	fmt.Println("arr[2:] = ", arr[2:])   //arr[2:] =  [2 3 4 5 6 7]
	fmt.Println("arr[:] = ", arr[:])     //arr[:] =  [0 1 2 3 4 5 6 7]

	s := arr[2:6]
	s[0] = 100
	fmt.Println("s1 = ", s)    //s1 =  [100 3 4 5]
	fmt.Println("arr = ", arr) //arr =  [0 1 100 3 4 5 6 7]

	s1 := arr[:]
	fmt.Println("s1 = ", s1) //s1 =  [0 1 100 3 4 5 6 7]
	s1 = s1[:5]
	fmt.Println("s1 = ", s1) //s1 =  [0 1 100 3 4]
	s1 = s1[2:]
	fmt.Println("s1 = ", s1) //s1 =  [100 3 4]

	arr[2] = 2
	s1 = arr[2:6]
	fmt.Println("s1 = ", s1)
	fmt.Printf("s1=%v, len(s1)=%d ,cap(s1)=%d\n", s1,len(s1),cap(s1))
	s2 := s1[3:5]
	fmt.Printf("s2=%v, len(s2)=%d ,cap(s2)=%d", s2,len(s2),cap(s2))
}
