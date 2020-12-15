package main

import (
	"log"
	"runtime"
	"sort"
	"strings"
	"time"
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
	modList := map[int64]int64{}

	max := 0
	for i, n := range busses {
		if n == -1 {
			continue
		}
		offset := (n - i)
		for offset < 0 {
			offset += n
		}
		offset %= n
		modList[int64(n)] = int64(offset)
		if n > max {
			max = n
		}
	}

	log.Println(modList, max)
	delete(modList, int64(max))
	a := make([]int, 0, len(modList))
	b := make([]int64, len(modList))
	for aa := range modList {
		a = append(a, int(aa))
	}
	sort.Ints(a)
	for i, n := range a {
		b[i] = modList[int64(n)]
	}
	numcores := int64(runtime.NumCPU())
	incr := int64(max)
	innerIncr := incr * numcores
	start := int64(0)
	for start%incr != modList[incr] {
		start++
	}

	startT := time.Now()
	var maxSeen int64
	for i := int64(0); i < numcores; i++ {
		thisStart := start + (i * incr)
		go func(start int64, i int64) {
			for {
				if i == 0 {
					maxSeen = start
				}
				ok := true
				for i := len(a) - 1; i >= 0; i-- {
					if start%int64(a[i]) != b[i] {
						ok = false
						break
					}
				}
				if ok {
					log.Println("!!!!!!", start, time.Now().Sub(startT))
					log.Fatal(start)
				}
				start += innerIncr
			}
		}(thisStart, i)
	}

	for {
		time.Sleep(time.Second)
		elapsed := time.Now().Sub(startT)
		left := time.Duration((float64(elapsed) / float64(maxSeen)) * float64(825305207525452-maxSeen))
		log.Println(left, time.Now().Sub(startT), float64(maxSeen)/float64(825305207525452))
	}
	return
}

var _ = register(13, day13, `1008141
17,x,x,x,x,x,x,41,x,x,x,x,x,x,x,x,x,523,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,13,19,x,x,x,23,x,x,x,x,x,x,x,787,x,x,x,x,x,37,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,29`)
