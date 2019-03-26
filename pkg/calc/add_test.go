package calc

import "testing"

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1+2=3",
			args: args{
				a: 1,
				b: 2,
			},
			want: 3,
		},
		{
			name: "2+3=5",
			args: args{
				a: 2,
				b: 3,
			},
			want: 5,
		},
		{
			name: "4+7=11",
			args: args{
				a: 4,
				b: 7,
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
