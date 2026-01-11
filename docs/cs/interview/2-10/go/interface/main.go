// You can edit this code!
// Click here and start typing.
package main

import (
	"errors"
	"fmt"
)

func main() {

	f1()
	f2()
	fmt.Println(f3())
}

type T interface{}

func f1() {
	a := []int{1, 2, 3, 4}
	ex(a)
	fmt.Println(a)
	add(a)
	fmt.Println(a)
}
func ex(a []int) {
	a = append(a, 5)
}

func add(a []int) {
	for i := 0; i < len(a); i++ {
		a[i] = a[i] + 5
	}
}

func f2() {
	var (
		t T
		p *T

		i1 interface{} = t
		i2 interface{} = p
	)

	fmt.Println(i1 == t, i1 == nil)
	fmt.Println(i2 == p, i2 == nil)
}

func f3() (err error) {
	defer func() {
		fmt.Println(err)
		err = errors.New("a")
	}()

	defer func(err error) {
		fmt.Println(err)
		err = errors.New("b")
	}(err)

	return errors.New("c")
}
