package main

import "testing"

func TestLooping(t *testing.T) {
	type args struct {
		character string
		amount    int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test three times",
			args: args{
				amount:    3,
				character: "H",
			},
			want: "HHH",
		},
		{
			name: "test one time",
			args: args{
				amount:    1,
				character: "H",
			},
			want: "H",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Looping(tt.args.character, tt.args.amount); got != tt.want {
				t.Errorf("Looping() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumOf(t *testing.T) {
	t.Run("test sumOfArray returns correct result", func(t *testing.T) {
		slice := make([]int, 2)
		slice[0] = 1
		slice[1] = 2
		got := sumOfArray(slice)
		if got != 3 {
			t.Errorf("got: %v but wanted 3", got)
		}
	})
}
