package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
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

	if int(s[3]-'0') == 0 {
		r.Lenght = int(s[2]-'0') + 9
	} else {
		r.Lenght = int(s[2]-'0') + int(s[3]-'0')
	}

	if r.Lenght != len(s)-4 {
		return &r, errors.New("El tamaño es incorrecto")
	}

	for i := 4; i < len(s); i++ {
		r.Value = r.Value + string(s[i])
	}

	if r.Type == "TX" {
		if _, err := strconv.Atoi(r.Value); err == nil {
			return &r, errors.New("No es un texto")
		}
	} else if r.Type == "NN" {
		if _, err := strconv.Atoi(r.Value); err != nil {
			return &r, errors.New("No es un numero")
		}
	}

	return &r, nil
}

func main() {
	str := os.Args[1]
	chars := []rune(str)
	fmt.Println(generateStruct(chars))

}
