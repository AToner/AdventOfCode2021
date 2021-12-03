package day02

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
		{name: "The Test", fileName: "../../input/day02/day02.txt", want: 1250395},
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
		want     float64
	}{
		{name: "The Test", fileName: "../../input/day02/day02.txt", want: 1451210346},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part2(tt.fileName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
