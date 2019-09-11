package main

import "testing"

func Test_Parameter(t *testing.T) {
	// gelenVeri := Cikti("Ali", "Aksoy")
	// beklenenCikti := "AliAkso"

	// if gelenVeri != beklenenCikti {
	// 	t.Errorf("Gelen veri: %s \n Beklenen veri: %s", gelenVeri, beklenenCikti)
	// }

	t.Run("Deneme", func(t *testing.T) {
		got := Cikti("Mehmet", "Ali")
		want := "MehmetAl"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Deneme2", func(t *testing.T) {
		got := Cikti("Mehmet", "Ali")
		want := "MehmetA"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

// --- FAIL: Test_Parameter (0.00s)
//     --- FAIL: Test_Parameter/Deneme (0.00s)
//         ..Go_Programming/Unit_Test/test3/parameters_test.go:18: got "MehmetAli" want "MehmetAl"
//     --- FAIL: Test_Parameter/Deneme2 (0.00s)
//         ..Go_Programming/Unit_Test/test3/parameters_test.go:27: got "MehmetAli" want "MehmetA"
