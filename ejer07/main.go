package main

import "fmt"


//funciones anonimas
// Declaring a variable called Calculo that is a function that takes two ints and returns an int.
//se declar al interior de una variable
var Calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 +	num2
}
func main() {
	fmt.Printf("Sumo 5 + 7 = %d \n", Calculo(5, 7))

	//restamos
	Calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}
	fmt.Printf("Resto 7 - 5 = %d \n", Calculo(7, 5))

	//dividimos
	Calculo = func(num1 int, num2 int) int {
		return num1 / num2
	}
	fmt.Printf("divido 12 / 3 = %d \n", Calculo(12, 3))

	Operaciones()

	//Closures
	tablaDel := 5
	MiTabla := Tabla(tablaDel)
	for i := 1; i < 11; i++ {
		fmt.Println(MiTabla())
	}
}

func Operaciones(){
	resultado := func() int {
		var a int = 23
		var b int = 14
		return a + b
	}
	fmt.Println(resultado())
}

func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia += 1
		return numero * secuencia
	}
}
