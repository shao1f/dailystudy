package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	Name string
}

func main() {
	var a int
	typeOfA := reflect.TypeOf(a)
	fmt.Println(typeOfA.Kind(), typeOfA.Name())
	s := MyStruct{}
	typeOfS := reflect.TypeOf(s)
	fmt.Println(typeOfS.Kind(), typeOfS.Name())
	type cat struct {
		Name   string `json:"name"`
		Age    int    `json:"age"`
		Gender int    `json:"gender"`
		birth  string `json:"birth"`
	}
	c := &cat{}
	typeOfC := reflect.TypeOf(c)
	// fmt.Println(typeOfC.Name(), ":", typeOfC.Kind())
	elemC := typeOfC.Elem()
	// fmt.Println(elemC.Name(), ":", elemC.Kind())
	fmt.Println(elemC.Field(3))
}
