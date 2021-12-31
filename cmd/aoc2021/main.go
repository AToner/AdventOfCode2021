package main

import (
	"andytoner.com/aoc2021/pkg/day13"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(start)
	fmt.Printf("%d\n", day13.Part1("./input/day13/day13.txt"))
	fmt.Println(time.Now().Sub(start))
}
