package main

import "fmt"

//arreglos estaticos y slices
var tabla [10]int
var matriz [5][7]int
func main() {
	matriz2 := []int{2, 5, 4}
	
	//arreglos estaticos
	tabla[0] = 1
	tabla[1] = 2
	tabla[2] = 3
	tabla[3] = 4
	tabla[4] = 5
	tabla[5] = 6
	tabla[6] = 7
	tabla[7] = 8
	tabla[8] = 9
	tabla[9] = 10

	//arreglos dinamicos
	//slice
	tabla2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i<len(tabla); i++{
		fmt.Printf("tabla 1 espacio %d valor %d \n", i, tabla[i])
	}
	for i := 0; i<len(tabla2); i++{
		fmt.Printf("tabla 2 espacio %d valor %d \n", i, tabla2[i])
	}
	fmt.Println(tabla)
	fmt.Println(tabla2)

	//matrices
	matriz[0][0] = 1
	matriz[0][1] = 2
	matriz[0][2] = 3
	matriz[0][3] = 4
	matriz[0][4] = 5
	matriz[0][5] = 6
	matriz[0][6] = 7
	matriz[1][0] = 8
	matriz[1][1] = 9
	matriz[1][2] = 10
	matriz[1][3] = 11
	matriz[1][4] = 12
	matriz[1][5] = 13
	matriz[1][6] = 14
	matriz[2][0] = 15
	matriz[2][1] = 16
	matriz[2][2] = 17
	matriz[2][3] = 18
	matriz[2][4] = 19
	matriz[2][5] = 20
	matriz[2][6] = 21
	matriz[3][0] = 22
	matriz[3][1] = 23
	matriz[3][2] = 24
	matriz[3][3] = 25
	matriz[3][4] = 26
	matriz[3][5] = 27
	matriz[3][6] = 28
	matriz[4][0] = 29
	matriz[4][1] = 30
	matriz[4][2] = 31
	matriz[4][3] = 32
	matriz[4][4] = 33
	matriz[4][5] = 34
	matriz[4][6] = 35

	for i := 0; i<len(matriz); i++{
		for j := 0; j<len(matriz[i]); j++{
			if j == 6 {
				fmt.Printf("%d \n", matriz[i][j])
			}else{
				fmt.Printf("%d ", matriz[i][j])
			}
		}
	}

	fmt.Println(matriz2)
	variante2()
	variante3()
	variante4()
}
func variante2(){
	elementos := [5]int{1, 2, 3, 4, 5}
	porcion := elementos[:3]
	fmt.Println(porcion)
}
func variante3(){
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d \n", len(elementos), cap(elementos))
}
func variante4(){
	nums := make([]int, 0, 0)
	for i := 0; i<100; i++{
		nums = append(nums, i)
	}
	fmt.Printf("Largo %d, Capacidad %d\n", len(nums), cap(nums))
	
}