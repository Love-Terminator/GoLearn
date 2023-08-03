package main

import (
	"fmt"
	"math"
)

var student string

const pi = 3.1415926

const (
	n1 = iota
	n2
	n3
	n4
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
)

func main() {
	student = "YPC"
	fmt.Println("hello world")
	fmt.Printf("name:%s\n", student)
	var x = 20
	fmt.Printf("number:%d\n", x)
	sex := "male"
	fmt.Printf("sex:%s\n", sex)
	r := 2.5
	radis := pi * math.Pow(r, 2)
	fmt.Printf("radis:%f\n", radis)

	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)

	me := "硅谷亮城workdog"
	for k, v := range me {
		fmt.Printf("id:%d, addr:%c\n", k, v)
	}
	for x := 1; x < 10; x++ {
		for y := 1; y <= x; y++ {
			fmt.Printf("%d * %d = %d ", y, x, x*y)
		}
		fmt.Println()
	}

}
