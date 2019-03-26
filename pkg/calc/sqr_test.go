package calc

import (
	"fmt"
	"testing"
)

type args struct {
	a int
}

type testSet struct {
	name string
	args args
	want int
}

var (
	tests = []testSet{
		{
			name: "0 is 0",
			args: args{
				a: 0,
			},
			want: 0,
		},
		{
			name: "2 is 4",
			args: args{
				a: 2,
			},
			want: 4,
		},
		{
			name: "3 is 9",
			args: args{
				a: 3,
			},
			want: 9,
		},
	}
)

func TestSqr(t *testing.T) {
	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sqr(tt.args.a); got != tt.want {
				t.Errorf("Sqr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSqrNative(t *testing.T) {
	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SqrNative(tt.args.a); got != tt.want {
				t.Errorf("SqrNative() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSqrMath(t *testing.T) {
	t.Parallel()
	for _, tt := range tests {
		if testing.Verbose() {
			t.Logf("Start test %s", tt.name)
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := SqrMath(tt.args.a); got != tt.want {
				t.Errorf("SqrMath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSqr(b *testing.B) {
	sqrs := []struct {
		name string
		fun  func(a int) int
	}{
		{"slow", Sqr},
		{"native", SqrNative},
		{"math", SqrMath},
	}
	for _, sqr := range sqrs {
		for k := 100; k <= 500; k = k + 100 {
			b.Run(fmt.Sprintf("%s/%d", sqr.name, k), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					sqr.fun(k)
				}
			})
		}
	}
}
