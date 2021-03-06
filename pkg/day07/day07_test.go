package day07

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
		{name: "Sample", args: args{fileName: "../../input/day07/day07_sample.txt"}, want: 37},
		{name: "Actual", args: args{fileName: "../../input/day07/day07.txt"}, want: 356922},
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
		{name: "Sample", args: args{fileName: "../../input/day07/day07_sample.txt"}, want: 168},
		{name: "Actual", args: args{fileName: "../../input/day07/day07.txt"}, want: 100347031},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part2(tt.args.fileName); got != tt.want {
				t.Errorf("Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
