package main

import "testing"

func TestHello(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Testin ismi1",
			args: args{
				name: "Hakan",
			},
			want: "Hello, worldHakan",
		},
		{
			name: "Testin ismi2",
			args: args{
				name: "Aksoy",
			},
			want: "Hello, worldHakan",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(tt.args.name); got != tt.want {
				t.Errorf("Hello() = %v, want %v", got, tt.want)
			}
		})
	}
}
