package main

import "fmt"

func main() {
}

func legjobbBaratom() string {
	return "Zsófi"
}

func legjobbBaratnoim() (string, string) {
	return "Zsofi", "Panka"
}
func enEletkorom() int {
	return 12
}

// tanulás:
// Amikor ezek megvannak, akkor gyere és szólj nekem, hogy csekkolhatom :) Jó  munkát kicsi szívem :)
// 1. Csinálj egy függvényt, ami visszaadja a kedvenc lovaidnak a nevét és teszteld is le
// 2. csinláj olyan függvényt, ami visszaadja a kedvenc lovad nevét és korát (amennyit élt) és tesztben irasd is ki a kettőt
// 3. csinálj egy olyan függvényt, ami visszaadja a kedvenc fagyi ízeid nevét sorrendben a legkedvencebbel (legalább3 fagyi legyen benne) és teszteld is le
// 4. csinálj egy olyan függvényt, ami visszaad egy viccet és képes vagy kiiratni azt egy tesztelési eljárásban
// 5. csinálj egy olyan függvényt, ami visszaadja a legjobb barátnőd nevét és annak korát és irattasd is ki egy tesztelési eljárásban teljes mondatban (pl Peti 46 éves)
// 6. csinálj egy olyan függvényt, aminek neve az, hogy fruzsiSzuperProgramozo és csak igazat ad vissza és irattasd is ki egy tesztelési eljárásban

func kedvencLovaimNevei() (string, string) {
	return "Ispán", "Berci"
}
func lovakEletkoraiNevei() (int, string, int, string) {
	return 16, "Berci", 13, "Ispán"
}
func kedvencFagyiIzeim() (string, string, string, string) {
	return "Afterait", "Vanillia", "Sztracsatella", "Csoki"
}
func vicc() string {
	return "fehér vagy és szereted a vért, sose alszol, tudom mi vagy... mond ki... ápoló"

}
func baratomNevKor() (string, int, string) {
	return "Zsófi", 12, "éves"
}
func fruzsiSzuperProgramozo() bool {
	return true
}
func fruzsiTortEletkora() float64 {
	return 11.5
}
func getFamily(mother string) string {
	return mother + ": Fruzsi, Frida"
	//return fmt.Sprint(mother, ": Fruzsi, Frida")
}

func greeting(firstName string, lastName string, country string) string {
	if country == "hu" {
		return fmt.Sprint("Szia", " ", lastName, " ", firstName, "!")

	} else {
		return fmt.Sprint("Hello", " ", firstName, " ", lastName, ",")
	}
}

// Csinálj egy függvéynt, ami várja a felhasználó keresztnevét és vezetéknevét és a korát és visszaadja, így pl:
// Szia Gebri Fruzsi, Te 11 éves vagy

// Csinálj egy függvényt, ami várja a kedvenc két lovad nevét az istálló (hely) nevét és visszaadja, hogy
// az Lovas Klub istállóban Berci és Fekete szépség áll

// Csinálj egy olyan függvényt, ami vár egy level (szint) paramétert, ami integer kell legyen és egy robot
// mozgásának energia szintjét adja vissza. Ha a level nagyobb mint 50 akkor a robot 3 egységet mehet előre,
// ha kevesebb mint 50 akkor csak 2 lépést mehet előre. Szóval itt egy level int -et kell átadnod paraméterként
// és egy movement (mozgás) int-et kell visszaadnod visszatérési értékként. Itt is használnod kell majd feltételt

func korNev(firstName string, lastName string, age int) string {
	return fmt.Sprint("Szia", " ", lastName, " ", firstName, ",", "te ", 11, " ", "éves vagy")

}
func lovakIstaloi(loNev1 string, loNev2 string, sistallo string) string {
	return fmt.Sprint(loNev1, " ", "és ", loNev2, " ", "a", " ", sistallo, " ", "laknak")
}
func robot(szint int) int {
	if szint > 50 {
		return 3
	} else {
		return 2
	}
}

// Csinálj egy olyan függvényt, ami várja egy ló életkorát int-ként és visszaadja, a ló korát szövegesen.
// Ha a ló kora kisebb vagy egyenlő 3-el, akkor ő csikó
// ha a ló kora nagyobb mint 3 akkor ő felnőtt

// Csinálj egy olyan függvényt, ami várja egy robot méretét int-ben pl 10-15-100 és visszaadja, szövegesen, hogy
// az adott robot alacsony, vagy magas-e. Ha a robot mérete 20 alatt van, akkor adja vissza, hogy "alacsony"
// ha 20 vagy fölötte van, akkor adja vissza, hogy magas.

func loEletkora(loEletkor int) string {
	if loEletkor <= 3 {
		return "a ló még csikó, most még nem kezelhető annyira"
	} else {
		return "A ló már elérte a felnőtt kort, akkor lett a legkezelhetőbb"

	}

}
func robotMerete(robotMerete int) string {
	if robotMerete <= 20 {
		return "alacsony"
	} else {
		return "magas"
	}

}
func tobbFeltetel(meret int) string {
	if meret < 1 {
		return "kicsi"
	} else if meret < 3 {
		return "közepes"
	} else if meret <= 10 {
		return "nagy"
	}
	return "oriasi"
}
func elsoNyeregProject(nev string) (int, string) {
	if nev == "apa" {
		return 160, "-at kell apának befűznie"
	} else if nev == "fruzsi" {
		return 80, "-at kell befűznie"
	} else if nev == "frida" {
		return 60, "-at kell befűznie"
	} else if nev == "anya" {
		return 120, "-at kell befűznie"
	}
	return 170, "-et kell befűznie"
}
