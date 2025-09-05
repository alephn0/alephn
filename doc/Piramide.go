package main

import (
	"fmt"
)

func main() {
	var n int
	// comentario
	fmt.Println("Digite la altura de la piramide ")
	fmt.Scanln(&n)
	for x := 1; x <= n; x++ { //CAMBIAR
		for z := 1; z <= n-x; z++ {
			fmt.Print(" ")
		}
		for z := 1; z <= 2*x-1; z++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}
