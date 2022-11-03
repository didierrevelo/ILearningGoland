package main


import (
	"fmt"
	"time"

	us "ejer10/user"

)


type pepe struct{
	us.Users
}
func main() {
	u := new(pepe)
	u.HigUser(1, "Juan", "Perez", "juanperez@gmail.com", time.Now(), true)
	fmt.Println(u.Users)
}
