package main

import (
	"andytoner.com/aoc2021/pkg/day17"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(start)
	fmt.Printf("%d\n", day17.Part2("./input/day17/day17.txt"))
	fmt.Println(time.Now().Sub(start))
}
