package main

import (
	"errors"
	"fmt"
	"os"
)

type Result struct {
	Type   string
	Lenght int
	Value  string
}

func generateStruct(s []rune) (*Result, error) {
	r := Result{
		Type:   "",
		Value:  "",
		Lenght: 0,
	}

	r.Type = string(s[0]) + string(s[1])
	if r.Type != "TX" && r.Type != "NN" {
		return &r, errors.New("No es un tipo valido")
	}
	r.Lenght = int(s[2]-'0') + int(s[3]-'0')

	if r.Lenght != len(s)-4 {
		return &r, errors.New("El tamaño es incorrecto")
	}

	for i := 4; i < len(s); i++ {
		r.Value = r.Value + string(s[i])
	}

	return &r, nil
}

func main() {
	str := os.Args[1]
	chars := []rune(str)
	fmt.Println(generateStruct(chars))

}
