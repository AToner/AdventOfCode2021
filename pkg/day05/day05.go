package day05

import (
	"andytoner.com/aoc2021/pkg/utils"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
--- Day 5: Hydrothermal Venture ---
You come across a field of hydrothermal vents on the ocean floor! These vents constantly produce large, opaque clouds,
so it would be best to avoid them if possible.

They tend to form in lines; the submarine helpfully produces a list of nearby lines of vents (your puzzle input) for
you to review. For example:

0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2
Each line of vents is given as a line segment in the format x1,y1 -> x2,y2 where x1,y1 are the coordinates of one end
the line segment and x2,y2 are the coordinates of the other end. These line segments include the points at both ends.
In other words:

An entry like 1,1 -> 1,3 covers points 1,1, 1,2, and 1,3.
An entry like 9,7 -> 7,7 covers points 9,7, 8,7, and 7,7.
For now, only consider horizontal and vertical lines: lines where either x1 = x2 or y1 = y2.

So, the horizontal and vertical lines from the above list would produce the following diagram:

.......1..
..1....1..
..1....1..
.......1..
.112111211
..........
..........
..........
..........
222111....
In this diagram, the top left corner is 0,0 and the bottom right corner is 9,9. Each position is shown as the number of
lines which cover that point or . if no line covers that point. The top-left pair of 1s, for example, comes from
2,2 -> 2,1; the very bottom row is formed by the overlapping lines 0,9 -> 5,9 and 0,9 -> 2,9.

To avoid the most dangerous areas, you need to determine the number of points where at least two lines overlap. In the
above example, this is anywhere in the diagram with a 2 or larger - a total of 5 points.

Consider only horizontal and vertical lines. At how many points do at least two lines overlap?


*/

type point struct {
	x int
	y int
}

type line struct {
	start point
	end   point
}

func Part1(fileName string) int {
	input := utils.ReadLines(fileName)
	lines, maxSize := inputToLines(input)

	matrix := makeMatrix(maxSize)

	for _, line := range lines {
		if !isDiagonal(line) {
			xMin, xMax, yMin, yMax := lineDetail(line)
			for x := xMin; x <= xMax; x++ {
				for y := yMin; y <= yMax; y++ {
					matrix[y][x]++
				}
			}
		}
	}
	return countMatrix(matrix, 2)
}

func inputToLines(input []string) ([]line, int) {
	var result []line

	var maxSize float64

	stringToInt := func(input string) int {
		result, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		maxSize = math.Max(maxSize, float64(result))
		return result
	}

	stringToPoint := func(input string) point {
		coords := strings.Split(input, ",")
		return point{
			x: stringToInt(coords[0]),
			y: stringToInt(coords[1]),
		}
	}

	for _, inputLine := range input {
		pointsString := strings.Split(inputLine, " -> ")
		startString, endString := pointsString[0], pointsString[1]
		result = append(result, line{
			start: stringToPoint(startString),
			end:   stringToPoint(endString),
		})
	}
	return result, int(maxSize)
}

func isDiagonal(input line) bool {
	return !(input.start.x == input.end.x || input.start.y == input.end.y)
}

func lineDetail(input line) (int, int, int, int) {
	xMin := int(math.Min(float64(input.start.x), float64(input.end.x)))
	xMax := int(math.Max(float64(input.start.x), float64(input.end.x)))
	yMin := int(math.Min(float64(input.start.y), float64(input.end.y)))
	yMax := int(math.Max(float64(input.start.y), float64(input.end.y)))
	return xMin, xMax, yMin, yMax
}

func countMatrix(input [][]int, limit int) int {
	result := 0
	for x := 0; x < len(input); x++ {
		for y := 0; y < len(input[x]); y++ {
			if input[y][x] >= limit {
				result++
			}
		}
	}
	return result
}

func makeMatrix(maxSize int) [][]int {
	matrix := make([][]int, maxSize+1)
	for i := range matrix {
		matrix[i] = make([]int, maxSize+1)
	}
	return matrix
}
