package main


import "fmt"

func main() {
	fmt.Println(uno(5))
	numero, estado := dos(1)
	fmt.Println(numero)
	fmt.Println(estado)
	fmt.Println(calculo(5, 46))
	fmt.Println(calculo(5, 46, 32, 25))
	fmt.Println(calculo(5, 46, 36, 46, 24, 33))
	fmt.Println(calculo(5, 46, 8, 66, 45, 64))
}

func uno(numero int) (z int) {
	z = numero * 2
	return
}

func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	} else {
	return 10, true
	}
}

func calculo(numero ...int) int {
	total := 0
	for _, num := range numero {
		total += num
	// 	fmt.Printf("\n Item %d \n", Item)
	}
	return total
	
}

