package utils

import (
	"fmt"
	"os"
	"strconv"
)

func ToInt(array []string) []int {
	var result []int
	for _, str := range array {
		current, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		result = append(result, current)
	}
	return result
}
