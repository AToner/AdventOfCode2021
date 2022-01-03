package day14

import "testing"

func TestPart1(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{name: "Sample", args: args{fileName: "../../input/day14/day14_sample.txt"}, want: 1588},
		{name: "Actual", args: args{fileName: "../../input/day14/day14.txt"}, want: 2010},
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
		want uint64
	}{
		{name: "Sample", args: args{fileName: "../../input/day14/day14_sample.txt"}, want: 2188189693529},
		{name: "Actual", args: args{fileName: "../../input/day14/day14.txt"}, want: 2010},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part2(tt.args.fileName); got != tt.want {
				t.Errorf("Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
