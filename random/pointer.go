package random

import "fmt"

func pointer2() {

	a := 5

	fmt.Println("a -", a)
	printPointer(&a)

}

func printPointer(a *int) {
	fmt.Println("a pointer 1 - ", a)
	fmt.Println("a pointer 2 - ", &a)
	fmt.Println("a pointer 3 - ", *a)
	fmt.Println("a pointer value 4 - ", *&a)
	fmt.Println("a pointer 5 - ", &*a)
	b := &a
	fmt.Println("b -", b)
	c := &b
	fmt.Println("c -", c)
	fmt.Println("actual value ", ***c)

}
