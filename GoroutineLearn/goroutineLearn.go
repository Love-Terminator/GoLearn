package main

import (
	"fmt"
	"runtime"
)

func main() {
	//var x int
	threads := runtime.GOMAXPROCS(0)
	fmt.Println(threads)
	/*
		for i := 0; i < threads; i++ {
			go func() {
				for { x++ }
			}()
		}
		time.Sleep(time.Second)
		fmt.Println("x =", x)

	*/
}
