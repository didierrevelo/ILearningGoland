package main


import (
	"fmt"
	"os"
	"bufio"
)

var num1, num2 int
var result int
var leyend string

func main() {
	fmt.Println("Enter a number: ")
	fmt.Scanf("%d", &num1)

	fmt.Println("Enter another number: ")
	fmt.Scanf("%d", &num2)

	fmt.Println("Desccrition : ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		leyend = scanner.Text()
	}

	result = num1 + num2

	fmt.Println(leyend , result)
}
