package main

import (
	"strings"
)

func day13(input string) (p1Result, p2Result int) {
	//input = `939
	//7,13,x,x,59,x,31,19`
	lines := Lines(input)
	current := Atoi(lines[0])
	split := strings.Split(lines[1], ",")
	busses := []int{}
	for _, bus := range split {
		if bus != "x" {
			busses = append(busses, Atoi(bus))
		} else {
			busses = append(busses, -1)
		}
	}
	// p1
	min := 9999
	for _, bus := range busses {
		if bus == -1 {
			continue
		}
		goneBy := current % bus
		remaining := bus - goneBy
		if remaining < min {
			min = remaining
			p1Result = remaining * bus
		}
	}

	// p2
	// modList := map[int64]int64{}

	// max := 0
	// for i, n := range busses {
	// 	if n == -1 {
	// 		continue
	// 	}
	// 	offset := (n - i)
	// 	for offset < 0 {
	// 		offset += n
	// 	}
	// 	offset %= n
	// 	modList[int64(n)] = int64(offset)
	// 	if n > max {
	// 		max = n
	// 	}
	// }

	// log.Println(modList, max)
	// incr := int64(max)
	// start := int64(0)
	// for start%incr != modList[incr] {
	// 	start++
	// }
	// delete(modList, max)
	// for {
	// 	log.Println(start)
	// 	ok := true
	// 	for n, mod := range modList {
	// 		if start%n != mod {
	// 			ok = false
	// 			break
	// 		}
	// 	}
	// 	if ok {
	// 		log.Println(start)
	// 		p2Result = int(start)
	// 		return
	// 	}
	// 	start += incr
	// }
	return
}

var _ = register(13, day13, `1008141
17,x,x,x,x,x,x,41,x,x,x,x,x,x,x,x,x,523,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,13,19,x,x,x,23,x,x,x,x,x,x,x,787,x,x,x,x,x,37,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,29`)
