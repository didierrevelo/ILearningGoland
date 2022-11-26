package main


import (
	"fmt"
	"bufio"
	"os"
	"io/ioutil"
)

func main(){
	leoArchivo()
	leoArchivo2()
	graboArchivo()
	graboArchivo2()
	graboArchivo3()
}

func leoArchivo(){
	archivo, err := ioutil.ReadFile("./archivo.txt")
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		fmt.Println(string(archivo))
	}
}

func leoArchivo2(){
	archivo, err := os.Open("./archivo.txt")
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		scanner := bufio.NewScanner(archivo)
		for scanner.Scan(){
			registro := scanner.Text()
			fmt.Println("Linea > " + registro)
		}
	}
	archivo.Close()
}

func graboArchivo(){
	archivo, err := os.Create("./nuevoArchivo.txt")
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		fmt.Fprintln(archivo, "Esta es una linea nueva")
		archivo.Close()
	}
	archivo.Close()
}

func graboArchivo2(){
	archivo, err := os.OpenFile("./nuevoArchivo.txt", os.O_WRONLY | os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		fmt.Fprintln(archivo, "Esta ees una linea nueva")
		archivo.Close()
	}
	archivo.Close()
}

func graboArchivo3(){
	fileName := "./nuevoArchivo.txt"
	if Append(fileName, "Esta es una 2 linea nueva") == false {
		fmt.Println("Error al grabar el archivo")
	}

}

func Append(fileName string, text string) bool {
	file, err := os.OpenFile(fileName, os.O_WRONLY | os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Hubo un error")
		return false
	}
	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println("Hubo un error")
		file.Close()
		return false
	}
	return true
}
