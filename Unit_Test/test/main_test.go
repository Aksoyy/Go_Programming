package main

import "testing"

func TestMain(t *testing.T) {
	cikti := Cikti()
	beklenen := "Test cikti"

	if cikti != beklenen {
		t.Errorf("Beklenen deger: %s \n Cikti: %s", beklenen, cikti)
	}
}
