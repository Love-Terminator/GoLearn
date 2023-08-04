package main

import (
	"fmt"
)

type feature interface {
	sound(name string)
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

func (cat *Cat) sound() {
	fmt.Printf("Hello, I am %s, Miao~~~", cat.name)
}

func (dog *Dog) sound() {
	fmt.Printf("Hello, I am %s, Wang~~~", dog.name)
}

func main() {

}
