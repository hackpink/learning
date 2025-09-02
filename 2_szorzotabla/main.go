package main

import (
	"fmt"
)

func main() {
	felhasznaloInput := 0
	fmt.Print("add meg melyik számot szorozzam = ")
	fmt.Scanln(&felhasznaloInput)
	for i := 1; i <= 10; i++ {
		szorzat := felhasznaloInput * i
		fmt.Println(felhasznaloInput, "x", i, "=", szorzat)
	}
}
