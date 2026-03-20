package main

import (
	"fmt"
	"strings"
)

func main() {
	separator := strings.Repeat("-", 24)

	fmt.Println("== Arrays ==")
	arr := [3]int{1, 2, 3}
	arr2 := arr
	arr[0] = 99
	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println(separator)

	fmt.Println("== Slices ==")
	s := []string{"foo", "bar", "baz"}
	s2 := s
	s[0] = "qux"
	fmt.Println(s, s2)
	fmt.Println(separator)

	fmt.Println("== Maps ==")
	m := map[string]int{"foo": 1, "bar": 2}
	m2 := m
	m["bar"] = 99
	delete(m, "foo")
	fmt.Println(m, m2)
	v, ok := m["foo"]
	fmt.Println(v, ok)
	fmt.Println(separator)

	fmt.Println("== Structs ==")
	type score struct {
		name  string
		score int
	}
	score1 := score{name: "Dent, Arthur", score: 87}
	score2 := score1
	score2.name = "MacMillan, Tricia"
	fmt.Println(score1, score2)
}
