package main


import "fmt"


func main(){
	paises := make(map[string]string)
	fmt.Println(paises)
	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"
	paises["Colombia"] = "Bogota"
	fmt.Println(paises)
	campeonato := map[string]int{
		"Barcelona": 39,
		"Real Madrid": 38,
		"Chivas": 37,
		"Boca Juniors": 30,
	}
	fmt.Println(campeonato)

	campeonato["River Plate"] = 25
	campeonato["Chivas"] = 55
	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato{
		fmt.Printf("El equipo %s, tiene un puntaje de: %d\n", equipo, puntaje)
	}
	puntaje, existe := campeonato["Chivas"]
	fmt.Printf("el puntaje capturado es %d, y el equipo existe %t \n", puntaje, existe)

}
