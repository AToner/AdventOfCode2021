package day17

import (
	"andytoner.com/aoc2021/pkg/utils"
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
)

/*
--- Day 17: Trick Shot ---
You finally decode the Elves' message. HI, the message says. You continue searching for the sleigh keys.

Ahead of you is what appears to be a large ocean trench. Could the keys have fallen into it? You'd better send a probe
to investigate.

The probe launcher on your submarine can fire the probe with any integer velocity in the x (forward) and y (upward, or
downward if negative) directions. For example, an initial x,y velocity like 0,10 would fire the probe straight up, while
an initial velocity like 10,-1 would fire the probe forward at a slight downward angle.

The probe's x,y position starts at 0,0. Then, it will follow some trajectory by moving in steps. On each step, these
changes occur in the following order:

The probe's x position increases by its x velocity.
The probe's y position increases by its y velocity.
Due to drag, the probe's x velocity changes by 1 toward the value 0; that is, it decreases by 1 if it is greater than 0,
increases by 1 if it is less than 0, or does not change if it is already 0.
Due to gravity, the probe's y velocity decreases by 1.
For the probe to successfully make it into the trench, the probe must be on some trajectory that causes it to be within
a target area after any step. The submarine computer has already calculated this target area (your puzzle input). For
example:

target area: x=20..30, y=-10..-5
This target area means that you need to find initial x,y velocity values such that after any step, the probe's x
position is at least 20 and at most 30, and the probe's y position is at least -10 and at most -5.

Given this target area, one initial velocity that causes the probe to be within the target area after any step is 7,2:

.............#....#............
.......#..............#........
...............................
S........................#.....
...............................
...............................
...........................#...
...............................
....................TTTTTTTTTTT
....................TTTTTTTTTTT
....................TTTTTTTT#TT
....................TTTTTTTTTTT
....................TTTTTTTTTTT
....................TTTTTTTTTTT
In this diagram, S is the probe's initial position, 0,0. The x coordinate increases to the right, and the y coordinate
increases upward. In the bottom right, positions that are within the target area are shown as T. After each step (until
the target area is reached), the position of the probe is marked with #. (The bottom-right # is both a position the
probe reaches and a position in the target area.)

Another initial velocity that causes the probe to be within the target area after any step is 6,3:

...............#..#............
...........#........#..........
...............................
......#..............#.........
...............................
...............................
S....................#.........
...............................
...............................
...............................
.....................#.........
....................TTTTTTTTTTT
....................TTTTTTTTTTT
....................TTTTTTTTTTT
....................TTTTTTTTTTT
....................T#TTTTTTTTT
....................TTTTTTTTTTT
Another one is 9,0:

S........#.....................
.................#.............
...............................
........................#......
...............................
....................TTTTTTTTTTT
....................TTTTTTTTTT#
....................TTTTTTTTTTT
....................TTTTTTTTTTT
....................TTTTTTTTTTT
....................TTTTTTTTTTT
One initial velocity that doesn't cause the probe to be within the target area after any step is 17,-4:

S..............................................................
...............................................................
...............................................................
...............................................................
.................#.............................................
....................TTTTTTTTTTT................................
....................TTTTTTTTTTT................................
....................TTTTTTTTTTT................................
....................TTTTTTTTTTT................................
....................TTTTTTTTTTT..#.............................
....................TTTTTTTTTTT................................
...............................................................
...............................................................
...............................................................
...............................................................
................................................#..............
...............................................................
...............................................................
...............................................................
...............................................................
...............................................................
...............................................................
..............................................................#
The probe appears to pass through the target area, but is never within it after any step. Instead, it continues down and
to the right - only the first few steps are shown.

If you're going to fire a highly scientific probe out of a super cool probe launcher, you might as well do it with
style. How high can you make the probe go while still reaching the target area?

In the above example, using an initial velocity of 6,9 is the best you can do, causing the probe to reach a maximum y
position of 45. (Any higher initial y velocity causes the probe to overshoot the target area entirely.)

Find the initial velocity that causes the probe to reach the highest y position and still eventually be within the
target area after any step. What is the highest y position it reaches on this trajectory?
4005
*/
type point struct {
	x int
	y int
}

type targetBounds struct {
	topLeft     point
	bottomRight point
}

func (tb *targetBounds) onTarget(location point) bool {
	return location.x >= tb.topLeft.x && location.x <= tb.bottomRight.x &&
		location.y <= tb.topLeft.y && location.y >= tb.bottomRight.y
}

func (tb *targetBounds) pastTarget(location point) bool {
	return location.x > tb.bottomRight.x || location.y < tb.bottomRight.y
}

func Part1(fileName string) int {
	input := utils.ReadLines(fileName)

	tb := getTargetArea(input[0])
	highestY := math.MinInt

	for xRate := 1; xRate < 200; xRate++ {
		for yRate := 1; yRate < 200; yRate++ {
			location := point{x: 0, y: 0}
			velocity := point{x: xRate, y: yRate}
			highestYForVelocity := math.MinInt
			for tick := 0; tick < 200; tick++ {
				location, velocity = calcPosition(location, velocity)
				if location.y > highestYForVelocity {
					highestYForVelocity = location.y
				}
				if tb.onTarget(location) {
					if highestYForVelocity > highestY {
						highestY = highestYForVelocity
					}
					break
				}
				if tb.pastTarget(location) {
					break
				}
			}
		}
	}
	return highestY
}

func minInt(a int, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func calcPosition(start point, velocity point) (point, point) {
	result := point{
		x: start.x + velocity.x,
		y: start.y + velocity.y,
	}

	nextVelocity := point{}
	if velocity.x > 0 {
		nextVelocity.x = velocity.x - 1
	} else if velocity.x < 0 {
		nextVelocity.x = velocity.x + 1
	}
	nextVelocity.y = velocity.y - 1
	return result, nextVelocity
}

func getTargetArea(input string) targetBounds {
	targetRegEx, _ := regexp.Compile("([x|y])=(\\-?[0-9]+)\\.\\.(\\-?[0-9]+)")

	var result targetBounds
	if targetRegEx.MatchString(input) {
		matches := targetRegEx.FindAllStringSubmatch(input, 2)
		var xMatch, yMatch int
		if matches[0][1] == "x" {
			xMatch, yMatch = 0, 1
		} else {
			xMatch, yMatch = 1, 0
		}

		toInt := func(stringIn string) int {
			result, err := strconv.ParseInt(stringIn, 10, 32)
			if err != nil {
				fmt.Printf("FAILED: Int %s. %v\n", stringIn, err)
				return -1
			}
			return int(result)
		}

		xs := []int{
			toInt(matches[xMatch][2]),
			toInt(matches[xMatch][3]),
		}
		sort.Ints(xs)

		ys := []int{
			toInt(matches[yMatch][2]),
			toInt(matches[yMatch][3]),
		}
		sort.Ints(ys)

		result = targetBounds{
			topLeft: point{
				x: xs[0],
				y: ys[1],
			},
			bottomRight: point{
				x: xs[1],
				y: ys[0],
			},
		}
	} else {
		fmt.Println("Nope")
	}
	return result
}
