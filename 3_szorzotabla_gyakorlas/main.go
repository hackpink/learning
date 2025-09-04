package main

import (
	"fmt"
	"math/rand"
)

func main() {

	szorzando := 0
	szorzat := 0

	//  felhasznaloValasztotaEredmenye := ""
	fmt.Print("ezt szorzom most, te add meg és az eredményt is= ")
	fmt.Scanln(&szorzando)

	for {
		szorzo := rand.Intn(10) + 1
		fmt.Print(szorzo, " * ", szorzando, "= ")
		fmt.Scanln(&szorzat)

		szorzatGep := szorzo * szorzando
		if szorzatGep == szorzat {
			fmt.Println("jó lett")
		} else {
			fmt.Println("elcseszted, próbáld másikkal, de ez a jó: ", szorzatGep)
		}
	}

}
