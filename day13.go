package main

import (
	"log"
	"strings"
)

func day13(input string) (p1Result, p2Result int) {
	input = `
7,13,x,x,59,x,31,19`
	lines := Lines(input)
	current := Atoi(lines[0])
	log.Println(current)
	split := strings.Split(lines[1], ",")
	busses := []int{}
	for _, bus := range split {
		if bus != "x" {
			busses = append(busses, Atoi(bus))
		}
	}
	// p1
	min := 9999
	for _, bus := range busses {
		goneBy := current % bus
		remaining := bus - goneBy
		if remaining < min {
			min = remaining
			p1Result = remaining * bus
		}
	}

	return
}

var _ = register(13, day13, `1008141
17,x,x,x,x,x,x,41,x,x,x,x,x,x,x,x,x,523,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,13,19,x,x,x,23,x,x,x,x,x,x,x,787,x,x,x,x,x,37,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,29`)
