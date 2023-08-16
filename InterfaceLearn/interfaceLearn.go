package main

import (
	"bytes"
	"fmt"
	"github.com/ugorji/go/codec"
)

type People struct {
	Name string
	Age  int
	High float32
}

/*
type Animal struct {
	animalPeople *People
}

func newAnimal(ctx context.Context) (tx interface{}, err error) {
	people := People{
		name: "Wang",
		age:  23,
		high: 172.6,
	}
	return &Animal{&people}, nil
}

*/

func MsgPackMarshal(v interface{}) ([]byte, error) {
	fmt.Println(v)
	var buf = new(bytes.Buffer)
	fmt.Println(buf)
	enc := codec.NewEncoder(buf, new(codec.MsgpackHandle))
	people := &People{
		Name: "Jack",
		Age:  18,
		High: 178.6,
	}
	fmt.Println(*people)
	err := enc.Encode(*people)
	fmt.Println(buf)
	fmt.Println(buf.Bytes())

	return buf.Bytes(), err
}

func MsgPackUnMarshal(data []byte, v interface{}) error {
	var buf = bytes.NewBuffer(data)
	dec := codec.NewDecoder(buf, new(codec.MsgpackHandle))
	return dec.Decode(v)
}

type A struct {
	Name string
	Age  int
	High int
}
type B float64

/*
var v1 A
var v2 *A = &v1
var v3 int = 9
var v4 bool = false
var v5 interface{} = v3
var v6 interface{} = nil
var v7 B
var v8 *B = &v7

*/

func main() {
	//tx, _ := newAnimal(context.Background())
	//fmt.Println(tx)
	//fmt.Println(tx.(*Animal).animalPeople)
	//fmt.Println(string([]byte{0xFE}))

	people := &People{
		Name: "Li",
		Age:  20,
		High: 172.6,
	}

	marshal, err := MsgPackMarshal(people)
	if err != nil {
		return
	}
	fmt.Println(marshal)

	newPeople := &People{}
	err = MsgPackUnMarshal(marshal, newPeople)
	if err != nil {
		return
	}
	fmt.Println(newPeople)

	v1 := A{
		Name: "king",
		Age:  50,
		High: 176,
	}

	var buf = new(bytes.Buffer)
	// var b = make([]byte, 0, 64)
	var h = new(codec.MsgpackHandle)
	var enc = codec.NewEncoder(buf, h)
	err = enc.Encode(v1)
	fmt.Println(buf)

	v20 := &A{}

	// var b1 []byte
	// ... assume b contains the bytes to decode from
	var h1 codec.Handle = new(codec.MsgpackHandle)
	var dec *codec.Decoder = codec.NewDecoder(buf, h1)
	dec.Decode(v20) //
	fmt.Println(v20)
}
