package main

import (
	"fmt"
	"reflect"
)

type My_struct struct {
	Name  *string
	Age   *int
	House *string
}

type Another_struct struct {
	House   *string
	Age     *int
	HaveDog *bool
	Name    *string
}

func main() {

	a := &My_struct{}

	str := "beijing"
	number := 20
	flag := true
	str2 := "james"
	b := &Another_struct{&str, &number, &flag, &str2}

	c := reflect.TypeOf(*a) //获取type
	d := reflect.TypeOf(*b)

	g := reflect.ValueOf(a).Elem() //获取value
	h := reflect.ValueOf(b).Elem()
	//双循环，对相同名字对字段进行赋值
	for i := 0; i < c.NumField(); i++ {
		for j := 0; j < d.NumField(); j++ {
			fmt.Println(h.Field(i))
			if c.Field(i).Name == d.Field(j).Name {
				g.Field(i).Set(h.Field(j))
			}
		}
	}
	fmt.Println(*a, *b)
	//{james 20 beijing} {beijing 20 true james}
}
