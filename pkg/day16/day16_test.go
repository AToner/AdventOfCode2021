package day16

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
		{name: "Sample", args: args{fileName: "../../input/day16/day16_sample.txt"}, want: 82},
		{name: "Actual", args: args{fileName: "../../input/day16/day16.txt"}, want: 893},
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
		{name: "Sample", args: args{fileName: "../../input/day16/day16_sample2.txt"}, want: 1},
		{name: "Actual", args: args{fileName: "../../input/day16/day16.txt"}, want: 4358595186090},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part2(tt.args.fileName); got != tt.want {
				t.Errorf("Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
