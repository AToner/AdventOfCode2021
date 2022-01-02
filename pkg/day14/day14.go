package day14

import (
	"andytoner.com/aoc2021/pkg/utils"
	"math"
	"strings"
	"sync"
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
After step 2: NBCCNBBBCBHC
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

func Part1(fileName string) uint64 {
	input := utils.ReadLines(fileName)
	template, insertions := parseInputLines(input)

	return runSteps(template, insertions, 10)
}

func runSteps(template string, insertions map[string]string, steps int) uint64 {
	wg := &sync.WaitGroup{}
	finalCount := make(map[string]uint64)

	for _, char := range strings.Split(template, "") {
		addToCount(finalCount, char)
	}

	incomingCountChannel := make(chan map[string]uint64)
	for i := 0; i < len(template)-1; i++ {
		wg.Add(1)
		go firstPair(wg, template[i:i+2], insertions, steps-1, incomingCountChannel)
	}

	go func() {
		wg.Wait()
		close(incomingCountChannel)
	}()

	for count := range incomingCountChannel {
		finalCount = mergeCounts(finalCount, count)
	}

	high, low := getCounts(finalCount)
	return high - low
}

func firstPair(wg *sync.WaitGroup, template string, insertions map[string]string, step int, counts chan map[string]uint64) {
	defer wg.Done()
	count := make(map[string]uint64)
	addPair(template, insertions, step, count)
	counts <- count
}

func addPair(template string, insertions map[string]string, step int, count map[string]uint64) {
	additionalChar := insertions[template]
	addToCount(count, additionalChar)

	if step == 0 {
		return
	}

	pair1 := template[:1] + additionalChar
	pair2 := additionalChar + template[1:]
	addPair(pair1, insertions, step-1, count)
	addPair(pair2, insertions, step-1, count)
}

func getCounts(count map[string]uint64) (uint64, uint64) {
	highCount := uint64(0)
	lowCount := uint64(math.MaxUint64)

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

func addToCount(count map[string]uint64, char string) {
	if _, ok := count[char]; ok {
		count[char] += 1
	} else {
		count[char] = 1
	}
}

func mergeCounts(count1 map[string]uint64, count2 map[string]uint64) map[string]uint64 {
	result := make(map[string]uint64)
	for key, value := range count1 {
		result[key] = value
	}
	for key, value := range count2 {
		result[key] = value + count1[key]
	}
	return result
}

/*
--- Part Two ---
The resulting polymer isn't nearly strong enough to reinforce the submarine. You'll need to run more steps of the pair
insertion process; a total of 40 steps should do it.

In the above example, the most common element is B (occurring 2192039569602 times) and the least common element is H
(occurring 3849876073 times); subtracting these produces 2188189693529.

Apply 40 steps of pair insertion to the polymer template and find the most and least common elements in the result.
What do you get if you take the quantity of the most common element and subtract the quantity of the least common
element?

*/
func Part2(fileName string) uint64 {
	input := utils.ReadLines(fileName)
	template, insertions := parseInputLines(input)
	return runSteps(template, insertions, 40)
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
