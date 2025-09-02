package main

import (
	"fmt"
	"math/rand"
	"time"
)

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
	return fmt.Sprint("Szia", " ", lastName, " ", firstName, ",", "te ", age, " ", "éves vagy")

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
func loFajtak(fajta string) string {
	if fajta == "appalusa" {
		return "fekete pöttyös fehér alapon"
	} else if fajta == "mustang" {
		return "lehet bármien színü, kivéve tarka"
	} else if fajta == "hidegvérü" {
		return "lehet bármien, de a leggyakoribb a tarka ami annyit tesz, hogy fekete pacás fehér alapon, és szőrös pata"
	}
	return "nem lehet tudni mijen, mert nem adtál meg fajtát"
}
func loElohelyek(elohelye string) string {
	if elohelye == "mongolia" {
		return "valószínüleg telivér"
	} else if elohelye == "island" {
		return "izlandon a lovak/pónik többsége izlandi póni a hosszu szőrük miatt ami melegen tartja őket, de akár musztángok is élhetnek ott"
	}
	return "nem lehet tudni, hogy mijen őshonos fajta él ott, de ami biztos van az a hidegvéru, mivel ahol megtalálhatóak a lovak ott biztos van hidegvérü, mivel nyugott és erős, szőval tökéletes hegyimentő ló, vagy csak akár oktató ló"
}

func nevemEsLovamNeve(loNev string) string {

	if loNev == "Stella" {
		return "Stella gazdája Fruzsina"

	} else {
		return "Nincs adat a lóról"
	}

}

func loKengyelBefuzes(emberNev string) int {

	switch emberNev {
	case "Fruzsi":
		return 11
	case "Frida":
		return 3
	case "Péter":

		return 15
	case "Ági":
		return 13
	}

	return 0
}

func tanulas(tantargy int) string {

	switch tantargy {

	case 11:
		return "töri"

	case 20:
		return "magyar"

	case 30:
		return "matek"
	}
	return "nincs megadva a 3 helyes szám egyike"

}

func ismetles(switch1 string) string {

	switch switch1 {

	case "switch":
		return "a switch a gyorsabb változata az if vagy if else ágaknak"

	case "if":
		return "az if ág több lehetöség kiiratását csinálja"

	default:
		return "nem tudom megmondani, nem adtál meg megfelelő típust"

	}

}

func gyakorlasIsmetles() {

	i := 1

	for {

		i++
		kocka := rand.Intn(6) + 1
		fmt.Print("\r", kocka)
		time.Sleep(300 * time.Millisecond)

		if i == 7 {
			break
		}
	}
}
