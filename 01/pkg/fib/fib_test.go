package fib

import "testing"

func TestGet(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "f0",
			n:    0,
			want: 0,
		},
		{
			name: "f1",
			n:    1,
			want: 1,
		},
		{
			name: "f2",
			n:    2,
			want: 1,
		},
		{
			name: "f3",
			n:    3,
			want: 2,
		},
		{
			name: "f5",
			n:    5,
			want: 5,
		},
		{
			name: "f7",
			n:    7,
			want: 13,
		},
		{
			name: "f10",
			n:    10,
			want: 55,
		},
		{
			name: "f20",
			n:    20,
			want: 6765,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Num(tt.n); got != tt.want {
				t.Errorf("call Num(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
