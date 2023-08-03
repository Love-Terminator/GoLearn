package main

import "fmt"

func main() {

	var arr1 [3]int
	arr1 = [3]int{1, 2, 3}
	fmt.Println(arr1)

	var arr2 [3]string = [3]string{"你", "好", "啊"}
	fmt.Println(arr2)

	arr3 := [...]int{5, 6, 7, 8, 9, 10}
	fmt.Println(arr3)

	arr4 := [10]int{2: 6, 5: 6, 0: 6}
	fmt.Println(arr4)

	for i := 0; i < len(arr4); i++ {
		fmt.Println(arr4[i])
	}

	for k, v := range arr2 {
		fmt.Println(k, v)
	}

	var arr6 [3][3]int = [3][3]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}
	fmt.Println(arr6)

	var slice1 []int
	slice1 = []int{1, 2, 3, 4, 5}
	slice2 := []string{"work", "dog"}
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Printf("len: %d, cap:%d\n", len(slice1), cap(slice1))

	slice3 := arr4[2:5]
	fmt.Printf("len: %d, cap:%d\n", len(slice3), cap(slice3))

	slice4 := make([]int, 5, 10)
	fmt.Println(slice4)
	fmt.Printf("len: %d, cap:%d\n", len(slice4), cap(slice4))

	slice5 := []string{"泰国", "新加坡", "印度尼西亚"}
	fmt.Printf("len: %d, cap:%d\n", len(slice5), cap(slice5))
	slice5 = append(slice5, "印度")
	fmt.Printf("len: %d, cap:%d\n", len(slice5), cap(slice5))

	x1 := []int{1, 3, 5}
	x2 := x1
	x3 := make([]int, 3, 5)
	copy(x3, x1)
	fmt.Println(x1, x2, x3)

	x5 := []int{1, 3, 5, 7, 9, 11}
	x5 = append(x5[:2], x5[4:]...)
	fmt.Println(x5)
	fmt.Println(cap(x5))

	arr5 := [...]int{1, 3, 5, 7, 9, 11}
	slice6 := arr5[:]
	fmt.Println(slice6)
	slice6 = append(slice6[:1], slice6[2:]...)
	fmt.Println(slice6)
	fmt.Println(arr5)

	fmt.Println("Test")
}
