package main

import "fmt"

type person struct {
	name string
	city string
	age  int8
}

func (p person) test() person {
	p.name = "lucy"

	return p
}

func main() {
	p := &person{
		name: "name",
		city: "shanghai",
		age:  1,
	}

	res := p.test()
	fmt.Println(res)

	var a int
	a = *new(int)
	fmt.Println(a)
}

//:= 声明局部变量
//var声明全局或者局部变量
