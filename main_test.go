package main

import "testing"

func TestCalculateSpeed(t *testing.T) {

	bekuldottSzint := 5
	elvartSebesseg := 15

	kapottSebesseg := calculateSpeed(bekuldottSzint)

	if kapottSebesseg != elvartSebesseg {
		t.Error("a kapott sebesség nem jó:", kapottSebesseg)
	}

}

func TestCalculateEnergy(t *testing.T) {

	bekuldottSzint := 5
	elvartEnergia := 500

	kapottEnergia := calculateEnergy(bekuldottSzint)

	if kapottEnergia != elvartEnergia {
		t.Error("a kapott energia nem jó", kapottEnergia)
	}
}
