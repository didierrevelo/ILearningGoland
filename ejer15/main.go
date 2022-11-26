package main 

import (
	"fmt"
	"strings"
	"time"
)


func main(){
	go miNombreLento("Didier")
	fmt.Println("Estoy aqui")
	var wait string
	fmt.Scanln(&wait)
}

func miNombreLento(nombre string) {
	letras := strings.Split(nombre, "")
	for _, letra := range letras {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(letra)
	}
}
