package main

import (
	"fmt"
	"testing"
)

func TestLegjobbBaratom(t *testing.T) {
	baratnom := legjobbBaratom()
	fmt.Println(baratnom)
}
func TestLegjobbBarataim(t *testing.T) {
	baratno1, baratno2 := legjobbBaratnoim()
	fmt.Println(baratno1, baratno2)
}
func TestEnEletkorom(t *testing.T) {

	szoveg := "én"
	kor := enEletkorom()
	szoveg2 := "vagyok"
	fmt.Println(szoveg, kor, szoveg2)
}

// gyakorlás
func TestKedvencLovaimNevei(t *testing.T) {
	fmt.Println(kedvencLovaimNevei())
}
func TestLovakEletkoraiNevei(t *testing.T) {
	fmt.Println(lovakEletkoraiNevei())
}
func TestKedvencFagyiIzeim(t *testing.T) {
	fmt.Println(kedvencFagyiIzeim())
}
func TestVicc(t *testing.T) {
	fmt.Println(vicc())
}
func TestBaratomNevKor(t *testing.T) {
	fmt.Println(baratomNevKor())

}
func TestFruzsiSzuperPtogramozo(t *testing.T) {
	fmt.Println(fruzsiSzuperProgramozo())
}
func TestFruzsiTortEletkora(t *testing.T) {
	fmt.Println(fruzsiTortEletkora())
}

// gyakorlás vége
func TestGetFamily(t *testing.T) {
	fmt.Println(getFamily("Éva"))

}
func TestGreeting(t *testing.T) {
	fmt.Println(greeting("Frida", "Gebri", "de"))
}

// gyakorlás
func TestKorNev(t *testing.T) {
	fmt.Println(korNev("Fruzsi", "Gebri", 11))
}
func TestLovakIstaloi(t *testing.T) {
	fmt.Println(lovakIstaloi("Berci", "Lucifer", "Spirit lovastanyába"))
}
func TestRobot(t *testing.T) {
	fmt.Println(robot(50))
}
func TestLoEletkora(t *testing.T) {
	fmt.Println(loEletkora(5))
}
func TestRobotMerete(t *testing.T) {
	fmt.Println(robotMerete(1))
}
func TestTobbFeltetel(t *testing.T) {
	fmt.Println(tobbFeltetel(20))
}
func TestElsoNyeregProject(t *testing.T) {
	fmt.Println(elsoNyeregProject(""))
}
func TestLoFajtak(t *testing.T) {
	fmt.Println(loFajtak("telivér"))

}
func TestLoElohelyek(t *testing.T) {
	fmt.Println(loElohelyek("magyarország"))
}

func TestNevemEsLonevem(t *testing.T) {
	fmt.Println(nevemEsLovamNeve("Ibiza"))
}

func TestLoKengyelBefuzes(t *testing.T) {
	fmt.Println(loKengyelBefuzes("Ibiza"))
}

func TestTanulas(t *testing.T) {
	fmt.Println(tanulas(30))
}

func TestIsmetles(t *testing.T) {
	fmt.Println(ismetles("izland"))
}

func TestGyakorlasIsmetles(t *testing.T) {
	gyakorlasIsmetles()
}
