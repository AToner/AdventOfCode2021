package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(path string) StringArray {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
		os.Exit(1)
	}
	return StringArray{data: lines}
}