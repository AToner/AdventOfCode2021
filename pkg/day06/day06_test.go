package day06

import (
	"testing"
)

func TestPart1(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Sample", args: args{fileName: "../../input/day06/day06_sample.txt"}, want: 5934},
		{name: "Actual", args: args{fileName: "../../input/day06/day06.txt"}, want: 359344},
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
		{name: "Sample", args: args{fileName: "../../input/day06/day06_sample.txt"}, want: 26984457539},
		{name: "Actual", args: args{fileName: "../../input/day06/day06.txt"}, want: 1629570219571},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part2(tt.args.fileName); got != tt.want {
				t.Errorf("Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
