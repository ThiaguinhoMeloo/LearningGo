package exampleimports

import (
	"fmt"
	"strconv"
)

func Strconv() {
	x, err := strconv.Atoi("4567")
	if err != nil {
		fmt.Println("Erro ao converter: ", err)
	}
	fmt.Println(x)

	y := strconv.Itoa(42)
	fmt.Println(y)

	b, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Erro ao converter a string em bool: ", err)
		return
	}
	fmt.Println(b)

	s := "verdadeiro"
	var r bool
	if s == "verdadeiro" {
		r = true
	} else {
		r = false
	}
	fmt.Println(r)
}
