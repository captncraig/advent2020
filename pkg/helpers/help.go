package helpers

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Lines(input string) []string {
	return strings.Split(input, "\n")
}

func LinesRegexp(input string, regex string) [][]string {
	lines := Lines(input)
	r := regexp.MustCompile(regex)
	results := make([][]string, len(lines))
	for i, line := range lines {
		match := r.FindStringSubmatch(line)
		if len(match) == 0 {
			log.Fatalf("Line '%s' doesn't match regex", line)
		}
		results[i] = match[1:]
	}
	return results
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

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("Cannot convert '%s' to int", s)
	}
	return i
}

func TimeMe() func() {
	start := time.Now()
	return func() {
		log.Println(time.Now().Sub(start))
	}
}
