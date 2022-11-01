package main

import "fmt"

var numero int
var texto string
var status bool

func main() {
	number1, number2, number3, texto := 1, 2, 3, "Hola"
	var moneda float32 = 1.5
	var status bool = true
	fmt.Println(number1, number2, number3, texto, moneda, status)
	fmt.Println(number1)
	fmt.Println(number2)
	fmt.Println(number3)	
	fmt.Println(texto)
}
