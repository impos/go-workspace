package main

import (
	"fmt"

	wspace "github.com/impos/go-workspace"
)

func main() {
	err := wspace.Error("test error")
	fmt.Println(err)

	fb := wspace.FooBarFn(1, 2)
	fmt.Printf("%+v\n", fb)
}
