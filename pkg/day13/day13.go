package day13

import (
	"andytoner.com/aoc2021/pkg/utils"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
--- Day 13: Transparent Origami ---
You reach another volcanically active part of the cave. It would be nice if you could do some kind of thermal
imaging so you could tell ahead of time which caves are too hot to safely enter.

Fortunately, the submarine seems to be equipped with a thermal camera! When you activate it, you are greeted with:

Congratulations on your purchase! To activate this infrared thermal imaging camera system, please enter the code found
on page 1 of the manual.
Apparently, the Elves have never used this feature. To your surprise, you manage to find the manual; as you go to
open it, page 1 falls out. It's a large sheet of transparent paper! The transparent paper is marked with random dots
and includes instructions on how to fold it up (your puzzle input). For example:

6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5
The first section is a list of dots on the transparent paper. 0,0 represents the top-left coordinate. The first
value, x, increases to the right. The second value, y, increases downward. So, the coordinate 3,0 is to the right
of 0,0, and the coordinate 0,7 is below 0,0. The coordinates in this example form the following pattern, where # is a
dot on the paper and . is an empty, unmarked position:

...#..#..#.
....#......
...........
#..........
...#....#.#
...........
...........
...........
...........
...........
.#....#.##.
....#......
......#...#
#..........
#.#........
Then, there is a list of fold instructions. Each instruction indicates a line on the transparent paper and wants you
to fold the paper up (for horizontal y=... lines) or left (for vertical x=... lines). In this example, the first
fold instruction is fold along y=7, which designates the line formed by all of the positions where y is 7
(marked here with -):

...#..#..#.
....#......
...........
#..........
...#....#.#
...........
...........
-----------
...........
...........
.#....#.##.
....#......
......#...#
#..........
#.#........
Because this is a horizontal line, fold the bottom half up. Some of the dots might end up overlapping after the fold
is complete, but dots will never appear exactly on a fold line. The result of doing this fold looks like this:

#.##..#..#.
#...#......
......#...#
#...#......
.#.#..#.###
...........
...........
Now, only 17 dots are visible.

Notice, for example, the two dots in the bottom left corner before the transparent paper is folded; after the fold is
complete, those dots appear in the top left corner (at 0,0 and 0,1). Because the paper is transparent, the dot just
below them in the result (at 0,3) remains visible, as it can be seen through the transparent paper.

Also notice that some dots can end up overlapping; in this case, the dots merge together and become a single dot.

The second fold instruction is fold along x=5, which indicates this line:

#.##.|#..#.
#...#|.....
.....|#...#
#...#|.....
.#.#.|#.###
.....|.....
.....|.....
Because this is a vertical line, fold left:

#####
#...#
#...#
#...#
#####
.....
.....
The instructions made a square!

The transparent paper is pretty big, so for now, focus on just completing the first fold. After the first fold in
the example above, 17 dots are visible - dots that end up overlapping after the fold is completed count as a single dot.

How many dots are visible after completing just the first fold instruction on your transparent paper?
753
*/
func Part1(fileName string) int {
	input := utils.ReadLines(fileName)
	grid, foldSteps := parseInputLines(input)

	grid = doFold(grid, foldSteps[0])

	displayGrid(grid)
	return countDots(grid)
}

/*
--- Part Two ---
Finish folding the transparent paper according to the instructions. The manual says the code is always eight capital
letters.

What code do you use to activate the infrared thermal imaging camera system?
HZLEHJRK
*/
func Part2(fileName string) int {
	input := utils.ReadLines(fileName)
	grid, foldSteps := parseInputLines(input)

	for _, foldStep := range foldSteps {
		grid = doFold(grid, foldStep)
	}
	displayGrid(grid)
	return countDots(grid)
}

func doFold(grid [][]bool, step string) [][]bool {
	stepInt, err := strconv.Atoi(step[2:])
	if err != nil {
		fmt.Printf("BARF! %s", step)
		os.Exit(1)
	}
	if step[0] == 'y' {
		grid = foldByY(grid, stepInt)
	} else {
		grid = foldByX(grid, stepInt)
	}
	return grid
}

func foldByY(grid [][]bool, yfold int) [][]bool {
	// Fold by Y (row merge)
	lastRow := yfold * 2
	for row := 0; row < yfold; row++ {
		for column, _ := range grid[row] {
			grid[row][column] = grid[row][column] || grid[lastRow][column]
		}
		lastRow--
	}
	grid = grid[:yfold]
	return grid
}

func foldByX(grid [][]bool, xfold int) [][]bool {
	columns := len(grid[0]) - 1

	// Fold by X (column range)
	for row := 0; row < len(grid); row++ {
		lastColumn := columns
		for column := 0; column < xfold; column++ {
			grid[row][column] = grid[row][column] || grid[row][lastColumn]
			lastColumn--
		}
		grid[row] = grid[row][:xfold]
	}
	return grid
}

func parseInputLines(inputLines []string) ([][]bool, []string) {
	var foldSteps []string

	columnCount, rowCount := maxDimensions(inputLines)

	result := make([][]bool, rowCount+1)
	for row := range result {
		result[row] = make([]bool, columnCount+1)
	}

	foldRegEx, _ := regexp.Compile(".=(?P<value>[0-9]+)")

	for _, line := range inputLines {
		if foldRegEx.MatchString(line) {
			match := foldRegEx.FindStringSubmatch(line)
			foldSteps = append(foldSteps, match[0])
		} else if line != "" {
			values := strings.Split(line, ",")
			column, _ := strconv.Atoi(values[0])
			row, _ := strconv.Atoi(values[1])
			result[row][column] = true
		}
	}
	return result, foldSteps
}

func maxDimensions(inputLines []string) (int, int) {
	maxColumn, maxRow := 0.0, 0.0

	for _, line := range inputLines {
		values := strings.Split(line, ",")
		if len(values) == 2 {
			column, _ := strconv.Atoi(values[0])
			row, _ := strconv.Atoi(values[1])
			maxColumn = math.Max(maxColumn, float64(column))
			maxRow = math.Max(maxRow, float64(row))
		}
	}
	return int(maxColumn), int(maxRow)
}

func displayGrid(grid [][]bool) {
	for rowIndex, _ := range grid {
		for columnIndex, _ := range grid[rowIndex] {
			if grid[rowIndex][columnIndex] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func countDots(grid [][]bool) int {
	result := 0
	for rowIndex, _ := range grid {
		for columnIndex, _ := range grid[rowIndex] {
			if grid[rowIndex][columnIndex] {
				result++
			}
		}
	}
	return result
}
