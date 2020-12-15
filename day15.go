package main

import (
	"log"
	"strings"
)

func day15(input string) (p1Result, p2Result int) {
	//input = "0,3,6"
	lastSeen := map[int]int{}
	seed := strings.Split(input, ",")
	seed = append([]string{""}, seed...)
	log.Println(seed)

	var (
		value  int
		exists bool
		prior  int
		diff   int
	)
	for i := 1; ; i++ {

		if i < len(seed) {
			value = Atoi(seed[i])
		} else {
			if !exists {
				value = 0
			} else {
				value = diff
			}
		}

		prior = lastSeen[value]
		exists = prior != 0
		diff = i - prior
		lastSeen[value] = i
		if i == 2020 {
			p1Result = value
		}
		if i == 30000000 {
			p2Result = value
			return
		}
	}
}

var _ = register(15, day15, `11,0,1,10,5,19`)
