package main

import "testing"

func TestPart1(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test", args{"data/test.txt"}, 7},
		{"solve", args{"data/input.txt"}, 1163},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part1(tt.args.fileName); got != tt.want {
				t.Errorf("Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test", args{"data/test.txt"}, "co,de,ka,ta"},
		{"solve", args{"data/input.txt"}, "bm,bo,ee,fo,gt,hv,jv,kd,md,mu,nm,wx,xh"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part2(tt.args.fileName); got != tt.want {
				t.Errorf("Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
