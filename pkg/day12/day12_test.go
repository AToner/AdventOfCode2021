package day12

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
		{name: "Sample", args: args{fileName: "../../input/day12/day12_sample.txt"}, want: 226},
		{name: "Actual", args: args{fileName: "../../input/day12/day12.txt"}, want: 5252},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part1(tt.args.fileName); got != tt.want {
				t.Errorf("Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
