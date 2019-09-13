package main

import "testing"

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