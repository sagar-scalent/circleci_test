package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddition(t *testing.T) {
	sum := adder(4, 3)
	assert.Equal(t, 7, sum, "Sum not the same")
}

func Test_adder(t *testing.T) {
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
			name: "Positive numbers",
			args: args{
				a: 1,
				b: 2,
			},
			want: 3,
		},
		{
			name: "First Positive ",
			args: args{
				a: 1,
				b: -2,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := adder(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("adder() = %v, want %v", got, tt.want)
			}
		})
	}
}
