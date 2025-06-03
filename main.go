package main

import "fmt"

func main() {
	name := "fruzsi"
	age := 11
	eletkora := "éves"
	fmt.Println(name, age, eletkora)
	age = age + 1
	// fruzsiJovoeviKora := age + 1
	eves := "lesz"

	fmt.Println(name, age, eletkora, eves)
	energy := calculateEnergy(5)
	fmt.Println(energy)
	speed := calculateSpeed(15)
	fmt.Println(speed)
}

func calculateEnergy(level int) int {
	energy := level * 100
	return energy

}

// a függvény a bot szintje alapján vissakell adja a mozgási sebességét
// a függvény neve lenyen calculateSpeed
// a robot alapsebessége 5
// a robot sebessége = az alapsebesség+az aktuális szint 2*-ese
func calculateSpeed(level int) int {
	speed := (level * 2) + 5
	return speed
}
