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
}

