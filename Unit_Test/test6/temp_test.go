package main

import "testing"

func TestCikti(t *testing.T) {
	type args struct {
		value []string
	}
	tests := []struct {
		name     string
		args     args
		wantVeri string
	}{
		{
			name: "First test",
			args: args{
				value: []string{"aksoy", "deneme2"},
			},
			wantVeri: "denemedeneme2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotVeri := Cikti(tt.args.value...); gotVeri != tt.wantVeri {
				t.Errorf("Cikti() = %v, want %v", gotVeri, tt.wantVeri)
			}
		})
	}
}

// --- FAIL: TestCikti (0.00s)
//     --- FAIL: TestCikti/First_test (0.00s)
//         ../Go_Programming/Unit_Test/test6/temp_test.go:25:
//	 						Cikti() = aksoydeneme2, want denemedeneme2
