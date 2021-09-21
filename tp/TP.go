package tp

import (
	"errors"
	"strconv"
)

type Result struct {
	Type   string
	Lenght int
	Value  string
}

func GenerateStruct(str string) (*Result, error) {
	r := Result{
		Type:   "",
		Value:  "",
		Lenght: 0,
	}

	s := []rune(str)

	r.Type = string(s[0]) + string(s[1])
	if r.Type != "TX" && r.Type != "NN" {
		return &r, errors.New("Is not a valid type")
	}

	if int(s[3]-'0') == 0 {
		r.Lenght = int(s[2]-'0') + 9
	} else {
		r.Lenght = int(s[2]-'0') + int(s[3]-'0')
	}

	if r.Lenght != len(s)-4 {
		return &r, errors.New("The lenght is invalid")
	}

	for i := 4; i < len(s); i++ {
		r.Value = r.Value + string(s[i])
	}

	if r.Type == "TX" {
		if _, err := strconv.Atoi(r.Value); err == nil {
			return &r, errors.New("Is not a text")
		}
	} else if r.Type == "NN" {
		if _, err := strconv.Atoi(r.Value); err != nil {
			return &r, errors.New("Is not a number")
		}
	}

	return &r, nil
}
