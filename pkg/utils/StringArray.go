package utils

import (
	"fmt"
	"os"
	"strconv"
)

type StringArray struct {
	data []string
}

func (array *StringArray) ToInt() []int {
	var result []int
	for _, str := range array.data {
		current, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		result = append(result, current)
	}
	return result
}
