package main

import (
	"fmt"
	// "io/ioutil"
	"log"
	// "os"
)


func main(){
	// archivo := "prueba.txt"
	// f, err := os.Open(archivo)

	// defer f.Close()

	// if err != nil {
	// 	fmt.Println("error")
	// 	os.Exit(1)
	// }
	ejemploPanic()
}

func ejemploPanic(){
	defer func(){
		reco := recover()
		if reco != nil {
			fmt.Println("Error detectado", reco)
			log.Fatalf("ocurrio un error que genero panic \n %v", reco)
			log.Fatal(reco)
			// log.Println("Error detectado", reco)
			// log.Println("Error detectado", reco)
			// log.Fatalln("Error detectado", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("Se encontro el valor de 1")
	}
}
