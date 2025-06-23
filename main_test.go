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
