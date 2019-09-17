// Stringify fonksiyonu çalıştırılınca oluşabilecek
// herhangi 2 davranışı test edin.
package main

import (
	"testing"
)

func TestStringify(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "string control",
			args: args{
				value: "test",
			},
			want:    "test",
			wantErr: true,
		},
		{
			name: "int control",
			args: args{
				value: 9,
			},
			want:    "9",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Stringify(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Stringify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Stringify() = %v, want %v", got, tt.want)
			}
		})
	}
}
