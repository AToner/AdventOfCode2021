package main

import (
	"andytoner.com/aoc2021/pkg/day15"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(start)
	fmt.Printf("%d\n", day15.Part1("./input/day15/day15.txt"))
	fmt.Println(time.Now().Sub(start))
}
