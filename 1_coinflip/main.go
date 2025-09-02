package main

import (
	"fmt"
	"log/slog"
	"math/rand"
	"strings"
	"time"
)

func main() {

	tipp := ""
	fmt.Print("tipp:")
	_, err := fmt.Scanln(&tipp)
	if err != nil {
		slog.Error("nem tudtuk beolvasni a user üzenetét", "err", err)
		return
	}
	tipp = strings.TrimSpace(tipp)
	tarolo := eremPrograma()
	if tipp == tarolo {
		fmt.Println("\nszép volt")
	} else {
		fmt.Println("\nnem talált, ez van")
	}

}

func eremPrograma() string {

	eremLehetoseg := []string{
		"fej", "írás",
	}
	eremAlapforgas := 4
	ero := rand.Intn(10) + 1
	forgasSzama := ero * eremAlapforgas
	eremIndulasiValtozat := rand.Intn(2)
	for i := 0; i < forgasSzama; i++ {
		fmt.Print("\r", eremLehetoseg[eremIndulasiValtozat])
		if i < forgasSzama-1 {
			if eremIndulasiValtozat == 0 {
				eremIndulasiValtozat = 1
			} else if eremIndulasiValtozat == 1 {
				eremIndulasiValtozat = 0
			}
		}

		time.Sleep(500 * time.Millisecond)
	}
	return eremLehetoseg[eremIndulasiValtozat]
}
