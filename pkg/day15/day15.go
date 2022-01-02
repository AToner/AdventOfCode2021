package day15

import (
	"andytoner.com/aoc2021/pkg/utils"
	"math"
	"strconv"
	"strings"
)

/*
--- Day 15: Chiton ---
You've almost reached the exit of the cave, but the walls are getting closer together. Your submarine can barely still
fit, though; the main problem is that the walls of the cave are covered in chitons, and it would be best not to bump
any of them.

The cavern is large, but has a very low ceiling, restricting your motion to two dimensions. The shape of the cavern
resembles a square; a quick scan of chiton density produces a map of risk level throughout the cave (your puzzle input).
For example:

1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581
You start in the top left position, your destination is the bottom right position, and you cannot move diagonally. The
number at each position is its risk level; to determine the total risk of an entire path, add up the risk levels of
each position you enter (that is, don't count the risk level of your starting position unless you enter it; leaving it
adds no risk to your total).

Your goal is to find a path with the lowest total risk. In this example, a path with the lowest total risk is
highlighted here:

1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581

The total risk of this path is 40 (the starting position is never entered, so its risk is not counted).

What is the lowest total risk of any path from the top left to the bottom right?

*/

type Cell struct {
	height      int
	lowestScore int
}

func Part1(fileName string) int {
	input := utils.ReadLines(fileName)
	grid := parseInputLines(input)

	lastRow := len(grid) - 1
	lastColumn := len(grid[0]) - 1

	var walkRoute func(int, int, int)
	walkRoute = func(row int, column int, previousScore int) {
		if row < 0 || row > lastRow || column < 0 || column > lastColumn {
			return
		}

		score := previousScore + grid[row][column].height

		if score >= grid[row][column].lowestScore {
			return
		}

		grid[row][column].lowestScore = score

		walkRoute(row-1, column, score)
		walkRoute(row, column-1, score)
		walkRoute(row+1, column, score)
		walkRoute(row, column+1, score)
	}

	walkRoute(0, 0, -grid[0][0].height)
	return grid[lastRow][lastColumn].lowestScore
}

func parseInputLines(inputLines []string) [][]Cell {
	rows := len(inputLines)
	columns := len(inputLines[0])

	result := emptyGrid(rows, columns)

	for rowIndex, row := range inputLines {
		for columnIndex, valueString := range strings.Split(row, "") {
			value, _ := strconv.Atoi(valueString)
			result[rowIndex][columnIndex] = Cell{
				height:      value,
				lowestScore: math.MaxInt,
			}
		}
	}
	return result
}

func emptyGrid(rows int, columns int) [][]Cell {
	result := make([][]Cell, rows)
	for row := range result {
		result[row] = make([]Cell, columns)
	}
	return result
}
