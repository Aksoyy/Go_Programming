package main

import "testing"

func Test_Parameter(t *testing.T) {
	gelenVeri := Cikti("Ali", "Aksoy")
	beklenenCikti := "AliAkso"

	if gelenVeri != beklenenCikti {
		t.Errorf("Gelen veri: %s \n Beklenen veri: %s", gelenVeri, beklenenCikti)
	}
}
