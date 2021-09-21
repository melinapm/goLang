package main

import (
	"fmt"

	"tpe2021.com/tp"
)

func main() {
	r, e := tp.GenerateStruct("TX04000A")
	if e == nil {
		fmt.Println(r)
	} else {
		fmt.Println(e)
	}
}
