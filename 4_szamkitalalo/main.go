package main

import (
	"fmt"
	"math/rand"
)

func main() {
	botValasztasa := rand.Intn(100) + 1
	for i := 0; i <= 7; i++ {
		userValasztasa := 0
		if i == 7 {
			fmt.Println(botValasztasa, " volt a megfelelő")
			break
		}
		fmt.Print("adj meg egy számot: ")
		fmt.Scanln(&userValasztasa)
		if userValasztasa == botValasztasa {
			fmt.Println("Eltaláltad")
		}
		if userValasztasa < botValasztasa {
			fmt.Println("több")
		}
		if userValasztasa > botValasztasa {
			fmt.Println("kevessebb")
		}

	}

}
