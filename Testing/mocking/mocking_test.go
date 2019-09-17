package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	type args struct {
		sleeper Sleeper
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		// {
		// 	name: "test",
		// 	// args: args{
		// 	// 	sleeper: Sleeper(),
		// 	// },
		// 	wantOut: "deneme",
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			Countdown(out, tt.args.sleeper)
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("Countdown() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
