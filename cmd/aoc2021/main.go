package main

import (
	"andytoner.com/aoc2021/pkg/day16"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(start)
	fmt.Printf("%d\n", day16.Part2("./input/day16/day16.txt"))
	fmt.Println(time.Now().Sub(start))
}
