package helpers

import (
	"log"
	"strconv"
	"strings"
)

func Lines(input string) []string {
	return strings.Split(input, "\n")
}

func Ints(input string) []int {
	lines := Lines(input)
	nums := make([]int, len(lines))
	var err error
	for i, n := range lines {
		nums[i], err = strconv.Atoi(n)
		if err != nil {
			log.Fatalf("%s is not a number!", n)
		}
	}
	return nums
}
