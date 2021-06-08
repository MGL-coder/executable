package main

import (
	"fmt"
	"github.com/MGL-coder/library"
)

func main () {
	fmt.Print("Changing the case of \"computer\" string:")
	fmt.Println(library.ChangeCase("computer"))

	fmt.Print("Changing the case of \"USA\" string:")
	fmt.Println(library.ChangeCase("USA"))

	x1, x2, _ := library.ToSolveQuadraticFunc(1, 5, 3)
	fmt.Print("The solution for the \"x^2 + 5x + 3 = 0\" is ")
	fmt.Printf("x1 = %v, x2 = %v\n", x1, x2)

	fmt.Println("The UUID:")
	fmt.Println(library.ToGetUUID())
}

