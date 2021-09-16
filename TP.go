package main

import (
	"fmt"
	"os"
)

type Result struct {
	Type   string
	Lenght int
	Value  string
}

func generateStruct(s []rune) *Result {
	r := Result{
		Type:   "",
		Value:  "",
		Lenght: 0,
	}

	r.Type = string(s[0]) + string(s[1])
	r.Lenght = int(s[2]-'0') + int(s[3]-'0')

	for i := 4; i < len(s); i++ {
		r.Value = r.Value + string(s[i])
	}

	return &r
}

func main() {
	str := os.Args[1]
	chars := []rune(str)
	fmt.Println(generateStruct(chars))

}
