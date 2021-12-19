package day09

import (
	"andytoner.com/aoc2021/pkg/utils"
	"math"
	"sort"
	"strconv"
	"strings"
)

/*
--- Day 9: Smoke Basin ---

These caves seem to be lava tubes. Parts are even still volcanically active; small hydrothermal vents release smoke
into the caves that slowly settles like rain.

If you can model how the smoke flows through the caves, you might be able to avoid it and be that much safer. The
submarine generates a heightmap of the floor of the nearby caves for you (your puzzle input).

Smoke flows to the lowest point of the area it's in. For example, consider the following heightmap:

2199943210
3987894921
9856789892
8767896789
9899965678

Each number corresponds to the height of a particular location, where 9 is the highest and 0 is the lowest a location
can be.

Your first goal is to find the low points - the locations that are lower than any of its adjacent locations. Most
locations have four adjacent locations (up, down, left, and right); locations on the edge or corner of the map have
three or two adjacent locations, respectively. (Diagonal locations do not count as adjacent.)

In the above example, there are four low points, all highlighted: two are in the first row (a 1 and a 0), one is in
the third row (a 5), and one is in the bottom row (also a 5). All other locations on the heightmap have some lower
adjacent location, and so are not low points.

The risk level of a low point is 1 plus its height. In the above example, the risk levels of the low points are 2, 1, 6,
and 6. The sum of the risk levels of all low points in the heightmap is therefore 15.

Find all of the low points on your heightmap. What is the sum of the risk levels of all low points on your heightmap?
486
*/

func Part1(fileName string) int {
	input := utils.ReadLines(fileName)
	grid, rowCount, columnCount := parseInputLines(input)

	var lowests []int
	for row := 0; row < rowCount; row++ {
		for column := 0; column < columnCount; column++ {
			cell := grid[row][column]
			neighbors := getNeighbors(grid, row, column)
			min := minInArray(neighbors)
			if cell < min {
				lowests = append(lowests, cell)
			}
		}
	}

	result := 0
	for _, lowest := range lowests {
		result += lowest + 1
	}

	return result
}

func getNeighbors(grid [][]int, row int, column int) []int {
	rowCount := len(grid)
	columnCount := len(grid[0])

	inboundsCell := func(row int, column int) []int {
		if row < 0 || row >= rowCount {
			return []int{}
		}
		if column < 0 || column >= columnCount {
			return []int{}
		}
		return []int{grid[row][column]}
	}

	var result []int
	result = append(result, inboundsCell(row-1, column)...)
	result = append(result, inboundsCell(row+1, column)...)
	result = append(result, inboundsCell(row, column-1)...)
	result = append(result, inboundsCell(row, column+1)...)
	return result
}

func minInArray(input []int) int {
	result := math.MaxInt
	for _, value := range input {
		if value < result {
			result = value
		}
	}
	return result
}

/*
--- Part Two ---
Next, you need to find the largest basins so you know what areas are most important to avoid.

A basin is all locations that eventually flow downward to a single low point. Therefore, every low point has a basin,
although some basins are very small. Locations of height 9 do not count as being in any basin, and all other locations
will always be part of exactly one basin.

The size of a basin is the number of locations within the basin, including the low point. The example above has four
basins.

The top-left basin, size 3:
2199943210
3987894921
9856789892
8767896789
9899965678

The top-right basin, size 9:
2199943210
3987894921
9856789892
8767896789
9899965678

The middle basin, size 14:
2199943210
3987894921
9856789892
8767896789
9899965678

The bottom-right basin, size 9:
2199943210
3987894921
9856789892
8767896789
9899965678

Find the three largest basins and multiply their sizes together. In the above example, this is 9 * 14 * 9 = 1134.

What do you get if you multiply together the sizes of the three largest basins?
1059300
*/

func Part2(fileName string) int {
	input := utils.ReadLines(fileName)
	grid, rowCount, columnCount := parseInputLines(input)

	var flood func(int, int) int
	flood = func(row int, column int) int {
		if row < 0 || row >= rowCount || column < 0 || column >= columnCount {
			return 0
		}

		if grid[row][column] == 9 || grid[row][column] == -1 {
			return 0
		}

		grid[row][column] = -1
		result := 1
		result += flood(row-1, column)
		result += flood(row+1, column)
		result += flood(row, column-1)
		result += flood(row, column+1)

		return result
	}

	var basinSizes []int
	for row := 0; row < rowCount; row++ {
		for column := 0; column < columnCount; column++ {
			if basinCount := flood(row, column); basinCount > 0 {
				basinSizes = append(basinSizes, basinCount)
			}
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))

	return basinSizes[0] * basinSizes[1] * basinSizes[2]
}

func parseInputLines(lines []string) ([][]int, int, int) {
	rowCount := len(lines)
	columnCount := len(lines[0])

	result := make([][]int, rowCount)
	for row := range result {
		result[row] = make([]int, columnCount)
	}

	for rowIndex, line := range lines {
		row := strings.Split(line, "")
		for columnIndex, cell := range row {
			result[rowIndex][columnIndex], _ = strconv.Atoi(cell)
		}
	}
	return result, rowCount, columnCount
}
