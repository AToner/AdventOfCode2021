package main

import (
	"andytoner.com/aoc2021/pkg/day15"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(start)
	fmt.Printf("%d\n", day15.Part2("./input/day15/day15_sample.txt"))
	fmt.Println(time.Now().Sub(start))
}
