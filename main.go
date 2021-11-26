package main

import (
	"fmt"
)

func main() {
	// MyAny
	nilValAny := types.MyAny{}.NewNil()
	fmt.Printf("nil to MyAny: %s\n", nilValAny.Get())

	hogeValAny := types.MyAny{}.New("hoge")
	var hogeValAnyGet string = hogeValAny.Get() // string型
	fmt.Printf("hoge to MyAny: %s\n", hogeValAnyGet)

	numValAny := types.MyAny{}.New(3)
	fmt.Printf("num to MyAny: %s\n", numValAny.Get())
}
