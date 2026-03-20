package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	separator := strings.Repeat("-", 24)

	fmt.Println("== Strings ==")
	interpreted := "this is an escape character:\n it created a new line"
	raw := `this is an escape character: \n it created a new line`
	fmt.Println(interpreted)
	fmt.Println(raw)
	fmt.Println(separator)

	fmt.Println("== Numbers ==")
	var i int = -42
	var u uint = 42
	var f32 float32 = 3.14
	var f64 float64 = 3.14159
	c := complex(1, 2)
	fmt.Printf("int=%d uint=%d float32=%.2f float64=%.5f complex=%v\n", i, u, f32, f64, c)
	fmt.Println(separator)

	fmt.Println("== Booleans ==")
	var ok bool
	fmt.Println("zero value:", ok)
	ok = true
	fmt.Println("set:", ok)
	fmt.Println(separator)

	fmt.Println("== Errors ==")
	var err error
	fmt.Println("zero value:", err)
	err = errors.New("something went wrong")
	fmt.Println("set:", err)
	fmt.Println(separator)

	fmt.Println("== Variables ==")
	var name string = "Felipe"
	score := 87
	fmt.Println(name, score)
	fmt.Println(separator)

	fmt.Println("== Constants ==")
	const a = 42
	const b string = "hello, world"
	const cConst = a * 2
	fmt.Println(a, b, cConst)
}
