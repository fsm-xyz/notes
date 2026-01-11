package main

import "fmt"

func main() {

	a := []byte("abc")

	fmt.Println(string(a), len(a), cap(a))

	b := append(a, []byte("d")...)

	fmt.Println(string(a), string(b), len(b), cap(b))

	a[2] = 69

	c := b[1:2]

	fmt.Println(string(a), string(b), string(c), len(c), cap(c))

	c = append(c, []byte("ef")...)

	fmt.Println(string(a), string(b), string(c))
}
