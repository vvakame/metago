// Code generated by metago. DO NOT EDIT.

//+build !metago

package main

import (
	"fmt"
)

type Foo struct {
	ID   int64
	Name string
}

func main() {
	obj := &Foo{1, "vvakame"}

	{
		fmt.Println("ID", obj.ID)
	}
	{
		fmt.Println("Name", obj.Name)
	}
}