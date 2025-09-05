package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Digite el número: ")
	fmt.Scan(&n)

	root := n       // CAMBIAR
	pe := int(root) // parte entera
	wsw := 0

	for i := 2; i < pe; i++ {
		if n%i == 0 {
			wsw = 1
			break // ya encontramos un divisor, no hace falta seguir
		}
	}

	if wsw == 1 {
		fmt.Printf("%d no es número primo\n", n)
	} else {
		fmt.Printf("%d es número primo\n", n)
	}
}
