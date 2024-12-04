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
		{"test", args{"data/test.txt"}, 18},
		{"solve", args{"data/input.txt"}, 0},
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
		want int
	}{
		{"test", args{"data/test.txt"}, 9},
		{"solve", args{"data/input.txt"}, 1990},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part2(tt.args.fileName); got != tt.want {
				t.Errorf("Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_At(t *testing.T) {
	type fields struct {
		Rows int
		Cols int
		data []string
	}
	type args struct {
		row int
		col int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{"simple-get-d", fields{2, 2, []string{"a", "b", "c", "d"}}, args{1, 1}, "d"},
		{"simple-get-a", fields{2, 2, []string{"a", "b", "c", "d"}}, args{0, 0}, "a"},
		{"simple-get-b", fields{2, 2, []string{"a", "b", "c", "d"}}, args{0, 1}, "b"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Matrix{
				Rows: tt.fields.Rows,
				Cols: tt.fields.Cols,
				data: tt.fields.data,
			}
			if got := m.At(tt.args.row, tt.args.col); got != tt.want {
				t.Errorf("At() = %v, want %v", got, tt.want)
			}
		})
	}
}
