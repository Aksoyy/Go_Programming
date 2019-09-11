package main

import "testing"

func Test_Parameter(t *testing.T) {

	// t.Run("Deneme", func(t *testing.T) {
	// 	got := Cikti("Mehmet", "Ali")
	// 	want := "MehmetAl"

	// 	if got != want {
	// 		t.Errorf("got %q want %q", got, want)
	// 	}
	// })

	commonOperation := func(t *testing.T, got, want string) {
		// t.Helper() --> Lock()-unLock() operation
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("MehmetAliBeklenenTEst", func(t *testing.T) {
		got := Cikti("Mehmet", "Ali")
		want := "MehmetAl"
		commonOperation(t, got, want)
	})

	t.Run("AksoybeklenenTEst", func(t *testing.T) {
		got := Cikti("Ak", "soy")
		want := "Aksoyy"
		commonOperation(t, got, want)
	})
}
package main

import "testing"

func Test_Parameter(t *testing.T) {

	// t.Run("Deneme", func(t *testing.T) {
	// 	got := Cikti("Mehmet", "Ali")
	// 	want := "MehmetAl"

	// 	if got != want {
	// 		t.Errorf("got %q want %q", got, want)
	// 	}
	// })

	commonOperation := func(t *testing.T, got, want string) {
		// t.Helper() --> Lock()-unLock() operation
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("MehmetAliBeklenenTEst", func(t *testing.T) {
		got := Cikti("Mehmet", "Ali")
		want := "MehmetAl"
		commonOperation(t, got, want)
	})

	t.Run("AksoybeklenenTEst", func(t *testing.T) {
		got := Cikti("Ak", "soy")
		want := "Aksoyy"
		commonOperation(t, got, want)
	})
}

// --- FAIL: Test_Parameter (0.00s)
//     --- FAIL: Test_Parameter/MehmetAliBeklenenTEst (0.00s)
//         ..Go_Programming/Unit_Test/test4/refactor_test.go:19: got "MehmetAli" want "MehmetAl"
//     --- FAIL: Test_Parameter/AksoybeklenenTEst (0.00s)
//         ..Go_Programming/Unit_Test/test4/refactor_test.go:19: got "Aksoy" want "Aksoyy"