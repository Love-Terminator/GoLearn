package main

import "fmt"

func main() {
	n := 18
	ptr := &n
	copy_n := *ptr
	fmt.Println(ptr)
	fmt.Printf("%T\n", ptr)
	fmt.Println(copy_n)
}
