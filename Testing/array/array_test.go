package array

import "testing"

func TestSum(t *testing.T) {
	type args struct {
		numbers [5]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				numbers: [5]int{1, 2, 3, 4, 5},
			},
			want: 15,
		},
		{
			name: "fail_test",
			args: args{
				numbers: [5]int{1, 2, 3, 4, 5},
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.numbers); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
