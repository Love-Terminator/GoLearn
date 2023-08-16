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
	ch1 := make(chan string, 2)
	ch2 := make(chan string, 2)

	var animal1 Animal
	animal1 = &Cat{name: "cat", age: 3, sex: "female"}

	var animal2 Animal
	animal2 = &Dog{name: "dog", age: 2, sex: "male"}

	exit := make(chan string)

	go animalCat(animal1, ch1, ch2)
	go animalDog(animal2, ch1, ch2, exit)

	<-exit
}

func animalCat(animal1 Animal, ch1 chan string, ch2 chan string) {
	for {
		select {
		case msg1, ok := <-ch2:
			if ok {
				fmt.Println("Cat: " + msg1)
				ch1 <- animal1.info()
			} else {
				fmt.Println("Cat: I am also Leave~~~")
				close(ch1)
				return
			}
		}
	}
}

func animalDog(animal2 Animal, ch1 chan string, ch2 chan string, exit chan string) {
	ch2 <- animal2.info()
	for {
		// v, ok := <-ch1
		select {
		case msg1, ok := <-ch1:
			if ok {
				fmt.Println("Dog: " + msg1)
				time.Sleep(10 * time.Second)
				fmt.Println("Dog: I am leave")
				close(ch2)
			} else {
				defer close(exit)
				return
			}
		}
	}

}
