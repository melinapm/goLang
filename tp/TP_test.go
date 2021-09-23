package tp_test

import (
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
	"tpe2021.com/tp"
)

func TestCadena(t *testing.T) {
	var cases = []struct {
		Input   string // input string in order to be parsed
		Success bool   // paser result
		Type    string // the input type
		Value   string // the input value
		Length  int    // value length
	}{
		{"TX02AB", true, "TX", "AB", 2},
		{"NN100987654321", true, "NN", "0987654321", 10},
		{"TX06ABCDE", false, "", "", 0},
		{"NN04000A", false, "", "", 0},
	}

	for _, testData := range cases {
		Result, err := tp.GenerateStruct(testData.Input)
		if assert.Equal(t, err, nil) { //CHEQUEA SI EL SEGUNDO PARAMETRO ES IGUAL AL SEGUNDO Y RETORNA UN BOOLEANO
			if assert.Equal(t, Result.Type, testData.Type) {
				if assert.Equal(t, Result.Length, testData.Length) {
					if assert.Equal(t, Result.Value, testData.Value) {
						fmt.Println("Cumplio Test")
					}
				}
			}
		}

	}
}
