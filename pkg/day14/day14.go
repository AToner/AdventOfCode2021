package day14

import (
	"andytoner.com/aoc2021/pkg/utils"
	"fmt"
	"math"
	"strings"
)

/*
--- Day 14: Extended Polymerization ---
The incredible pressures at this depth are starting to put a strain on your submarine. The submarine has polymerization
equipment that would produce suitable materials to reinforce the submarine, and the nearby volcanically-active caves
should even have the necessary input elements in sufficient quantities.

The submarine manual contains instructions for finding the optimal polymer formula; specifically, it offers a polymer
template and a list of pair insertion rules (your puzzle input). You just need to work out what polymer would result
after repeating the pair insertion process a few times.

For example:

NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C
The first line is the polymer template - this is the starting point of the process.

The following section defines the pair insertion rules. A rule like AB -> C means that when elements A and B are
immediately adjacent, element C should be inserted between them. These insertions all happen simultaneously.

So, starting with the polymer template NNCB, the first step simultaneously considers all three pairs:

The first pair (NN) matches the rule NN -> C, so element C is inserted between the first N and the second N.
The second pair (NC) matches the rule NC -> B, so element B is inserted between the N and the C.
The third pair (CB) matches the rule CB -> H, so element H is inserted between the C and the B.
Note that these pairs overlap: the second element of one pair is the first element of the next pair. Also, because
all pairs are considered simultaneously, inserted elements are not considered to be part of a pair until the next step.

After the first step of this process, the polymer becomes NCNBCHB.

Here are the results of a few steps using the above rules:

Template:     NNCB
After step 1: NCNBCHB
After step 2: NBCCNBBBCBHCB
After step 3: NBBBCNCCNBBNBNBBCHBHHBCHB
After step 4: NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB

This polymer grows quickly. After step 5, it has length 97; After step 10, it has length 3073. After step 10,
B occurs 1749 times, C occurs 298 times, H occurs 161 times, and N occurs 865 times; taking the quantity of the most
common element (B, 1749) and subtracting the quantity of the least common element (H, 161) produces 1749 - 161 = 1588.

Apply 10 steps of pair insertion to the polymer template and find the most and least common elements in the result.
What do you get if you take the quantity of the most common element and subtract the quantity of the least common
element?
2010
*/

func Part1(fileName string) int {
	input := utils.ReadLines(fileName)
	template, insertions := parseInputLines(input)

	count := make(map[string]int)

	addToCount := func(char string) {
		if _, ok := count[char]; ok {
			count[char] += 1
		} else {
			count[char] = 1
		}
	}

	getCounts := func() (int, int) {
		highCount := 0
		lowCount := math.MaxInt

		for _, value := range count {
			if value < lowCount {
				lowCount = value
			}

			if value > highCount {
				highCount = value
			}
		}
		return highCount, lowCount
	}

	for _, char := range strings.Split(template, "") {
		addToCount(char)
	}

	for step := 1; step <= 10; step++ {
		result := ""

		for i := 0; i < len(template)-1; i++ {
			pair := template[i : i+2]
			addToCount(insertions[pair])
			result = fmt.Sprintf("%s%c%s", result, pair[0], insertions[pair])
		}
		result = fmt.Sprintf("%s%c", result, template[len(template)-1])

		template = result
	}
	high, low := getCounts()
	return high - low
}

func parseInputLines(inputLines []string) (string, map[string]string) {
	insertions := make(map[string]string)
	template := inputLines[0]

	for _, line := range inputLines[2:] {
		insertion := strings.Split(line, " -> ")
		insertions[insertion[0]] = insertion[1]
	}

	return template, insertions
}
