package main

import "fmt"

type person struct {
	name string
}

func main() {
	fmt.Println("== Value semantics ==")
	a := 42
	b := a
	b = 27
	fmt.Println("a:", a, "b:", b)

	fmt.Println("== Pointer semantics ==")
	c := 42
	d := &c
	*d = 27
	fmt.Println("c:", c, "*d:", *d)

	fmt.Println("== Structs ==")
	p1 := person{name: "Ana"}
	p2 := p1
	p2.name = "Bia"
	fmt.Println("p1:", p1, "p2:", p2)

	p3 := &p1
	p3.name = "Carla"
	fmt.Println("p1:", p1, "p3:", *p3)
}
