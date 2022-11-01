package main

import "fmt"


func main() {
	
	for i:=1; i<=10; i++ {
		fmt.Println(i)
	}

	var i = 0
	for i < 10 {
		fmt.Printf("\n Valor de i: %d \n", i)
		if i == 5 {
			fmt.Printf("multiplicamos por 2 \n")
			i = i * 2
			fmt.Printf("Valor de i: %d \n", i)
			continue
		}
		fmt.Printf("Pasamos al siguiente valor \n")
		i++
	}

	var j int = 0
	RUTINA:
		for j < 10 {
			if j == 4 {
				j = j + 2
				fmt.Printf("\n Vamos a RUTINA sumando 2 a j = %d \n", j)
				goto RUTINA
			}
			fmt.Printf("\n Valor de j: %d \n", j)
			j++
		}
}
