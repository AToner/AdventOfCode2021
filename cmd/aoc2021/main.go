package main

import (
	"andytoner.com/aoc2021/pkg/day14"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(start)
	fmt.Printf("%d\n", day14.Part1("./input/day14/day14.txt"))
	fmt.Println(time.Now().Sub(start))
}
