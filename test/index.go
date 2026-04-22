package main

import "fmt"

type test struct {
	Name string
}

func main() {
	a := test{
		Name: "test1",
	}
	b := test{
		Name: "test2",
	}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(&a.Name)
}
