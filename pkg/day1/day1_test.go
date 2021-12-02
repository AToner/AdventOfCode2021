package day1

import (
	"reflect"
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		name     string
		fileName string
		want     int
	}{
		{name: "Sample", fileName: "../../input/day1/day1_sample.txt", want: 7},
		{name: "Actual", fileName: "../../input/day1/day1.txt", want: 1400},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part1(tt.fileName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name     string
		fileName string
		want     int
	}{
		{name: "Sample", fileName: "../../input/day1/day1_sample.txt", want: 5},
		{name: "Actual", fileName: "../../input/day1/day1.txt", want: 1429},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part2(tt.fileName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
