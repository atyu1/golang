package main

import (
	"fmt"
)

type person struct {
	Name string
	Age  int8
}

func (p *person) speak() {
	fmt.Println("My name is", p.Name, "and ", p.Age, "years old.")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		Name: "John Hendricks",
		Age:  34,
	}
  //Dont work
	//saySomething(p1)

  // Works
	saySomething(&p1)
}
