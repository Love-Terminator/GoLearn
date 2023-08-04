package main

import (
	"fmt"
	"strconv"
	"time"
)

type Animal interface {
	info() string
}

type Cat struct {
	name string
	age  int
	sex  string
}

type Dog struct {
	name string
	age  int
	sex  string
}

func (cat *Cat) info() string {
	return "Hello, My name is " + cat.name + ", My age is " + strconv.Itoa(cat.age) + ", My sex is " + cat.sex + ", Miao~~~"
}

func (dog *Dog) info() string {
	return "Hello, My name is " + dog.name + ", My age is " + strconv.Itoa(dog.age) + ", My sex is " + dog.sex + ", Wang~~~"
}

func main() {

	// 1、没有缓存的通道，再发送数据之后，必须从通道中接收数据，不然通道会堵塞
	/*
		ch := make(chan int)

		ch <- 1
		ch <- 2

		fmt.Println(<-ch)
		fmt.Println(<-ch)
	*/

	// 2、缓存为2元素的通道，这时通道中可以缓存2元素的数据，不会发生堵塞
	/*
		ch := make(chan string, 2)

		ch <- "Hello"
		ch <- "World"

		v, ok := <-ch
		fmt.Println(v, ok)
		v, ok = <-ch
		fmt.Println(v, ok)
		close(ch)
		v, ok = <-ch
		fmt.Println(v, ok)

	*/

	var animal1 Animal
	animal1 = &Cat{name: "cat", age: 3, sex: "female"}

	var animal2 Animal
	animal2 = &Dog{name: "dog", age: 2, sex: "male"}

	ch := make(chan string)
	exit := make(chan string)

	go animalCat(animal1, ch)
	go animalDog(animal2, ch, exit)

	<-exit

}

func animalCat(animal1 Animal, ch chan string) {
	for {
		if v, ok := <-ch; ok {
			fmt.Println("Cat: " + v)
			ch <- animal1.info()
		} else {
			fmt.Println("Cat: Bye!")
			return
		}
	}
}

func animalDog(animal2 Animal, ch chan string, exit chan string) {
	ch <- animal2.info()
	for {
		if v, ok := <-ch; ok {
			fmt.Println("Dog: " + v)
			time.Sleep(10 * time.Second)
			close(ch)
		} else {
			defer close(exit)
			fmt.Println("Dog: Bye!")
			return
		}
	}
}
