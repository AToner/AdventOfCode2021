package main

import (
	"andytoner.com/aoc2021/pkg/day12"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(start)
	fmt.Printf("%d\n", day12.Part1("./input/day12/day12_sample.txt"))
	fmt.Println(time.Now().Sub(start))
}
