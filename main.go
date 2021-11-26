package main

import (
	"fmt"

	"github.com/sakash1ta90/my_type_pg/types"
)

func main() {
	// MyAny
	nilValAny := types.MyAny{}.New(nil)
	fmt.Printf("nil to MyAny: %v\n", nilValAny.Get())

	hogeValAny := types.MyAny{}.New("hoge")
	var hogeValAnyGet string = hogeValAny.Get().(string) // string型
	fmt.Printf("hoge to MyAny: %v\n", hogeValAnyGet)

	numValAny := types.MyAny{}.New(3)
	fmt.Printf("num to MyAny: %v\n", numValAny.Get())
}
