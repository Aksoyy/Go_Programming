package dictionary

import "testing"

func TestDictionary_Search(t *testing.T) {
	//dictionary := Dictionary{"test": "this is just a test"}
	type args struct {
		word string
	}
	tests := []struct {
		name string
		d    Dictionary
		args args
		want string
	}{
		{
			name: "mytest",
			d:    Dictionary{"mykey": "this is just a test"},
			args: args{
				word: "mykey",
			},
			want: "this is just a test",
		},
		{
			name: "mytest2",
			d:    Dictionary{"mykey2": "this is just a test"},
			args: args{
				word: "mykey2",
			},
			want: "this is just a test2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Search(tt.args.word); got != tt.want {
				t.Errorf("Dictionary.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
