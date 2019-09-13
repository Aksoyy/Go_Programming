package main

import (
	"testing"
)

func TestDictionary_Search(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name    string
		d       Dictionary
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "mytest",
			d:    Dictionary{"a": "ali", "b": "berat"},
			args: args{
				word: "a",
			},
			want:    "ali",
			wantErr: false,
		},
		{
			name: "mytest2",
			d:    Dictionary{"a": "ali", "b": "berat"},
			args: args{
				word: "b",
			},
			want:    "berat",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.Search(tt.args.word)
			if (err != nil) != tt.wantErr {
				t.Errorf("Dictionary.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Dictionary.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ok
// Success: Tests passed.
// PASS

// ok
// Running tool: ..go test -timeout 30s -run ^(TestDictionary_Search)$
// Success: Tests passed.

func TestDictionary_Add(t *testing.T) {
	type args struct {
		word       string
		definition string
	}
	tests := []struct {
		name    string
		d       Dictionary
		args    args
		wantErr bool
	}{
		{
			name: "mytest",
			d: Dictionary{
				"ali": "veli",
			},
			args: args{
				word:       "genc",
				definition: "hakan",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.Add(tt.args.word, tt.args.definition); (err != nil) != tt.wantErr {
				t.Errorf("Dictionary.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// PASS
// Success: Tests passed.
// PASS
// Success: Tests passed.

func TestDictionary_Update(t *testing.T) {
	type args struct {
		word       string
		definition string
	}
	tests := []struct {
		name    string
		d       Dictionary
		args    args
		wantErr bool
	}{
		{
			name: "mytest",
			d: Dictionary{
				"a": "ali",
				"b": "berk",
			},
			args: args{
				word:       "b",
				definition: "veli",
			},
			wantErr: false,
		},
		{
			name: "mytest2",
			d: Dictionary{
				"a": "ali",
				"b": "berk",
			},
			args: args{
				word:       "c",
				definition: "veli",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.Update(tt.args.word, tt.args.definition); (err != nil) != tt.wantErr {
				t.Errorf("Dictionary.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// PASS
// Success: Tests passed.

func TestDictionary_Delete(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		d    Dictionary
		args args
	}{
		{
			name: "test27",
			d: Dictionary{
				"a": "b",
				"b": "c",
				"c": "d",
			},
			args: args{
				word: "c",
			},
		},
		{
			name: "test3",
			d: Dictionary{
				"a": "b",
				"b": "c",
				"c": "d",
			},
			args: args{
				word: "d",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.Delete(tt.args.word); err != true {
				t.Errorf("Dictionary.Update() error = %v", err)
			}
		})
	}
}

// PASS
// Success: Tests passed.
