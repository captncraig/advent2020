package main

import (
	"fmt"
	"log"
)

func day22(input string) (p1Result, p2Result int) {
	var p1 = []int{26, 16, 33, 8, 5, 46, 12, 47, 39, 27, 50, 10, 34, 20, 23, 11, 43, 14, 18, 1, 48, 28, 31, 38, 41}
	var p2 = []int{45, 7, 9, 4, 15, 19, 49, 3, 36, 25, 24, 2, 21, 37, 35, 44, 29, 13, 32, 22, 17, 30, 42, 40, 6}

	for len(p1) > 0 && len(p2) > 0 {
		a := p1[0]
		b := p2[0]
		p1 = p1[1:]
		p2 = p2[1:]
		if a > b {
			p1 = append(p1, a, b)
		} else {
			p2 = append(p2, b, a)
		}
	}

	score := func(winner []int) (s int) {
		mult := 1
		for i := len(winner) - 1; i >= 0; i-- {
			s += winner[i] * mult
			mult++
		}
		return s
	}
	if len(p1) > 0 {
		p1Result = score(p1)
	} else {
		p1Result = score(p2)
	}

	p1 = []int{26, 16, 33, 8, 5, 46, 12, 47, 39, 27, 50, 10, 34, 20, 23, 11, 43, 14, 18, 1, 48, 28, 31, 38, 41}
	p2 = []int{45, 7, 9, 4, 15, 19, 49, 3, 36, 25, 24, 2, 21, 37, 35, 44, 29, 13, 32, 22, 17, 30, 42, 40, 6}
	var recurse func(p1, p2 []int, d int) bool
	thewinner := []int{}
	tries := 0
	recurse = func(p1, p2 []int, depth int) bool {
		seen := map[string]bool{}
		for len(p1) > 0 && len(p2) > 0 {
			tries++
			hash := fmt.Sprint(p1, p2)
			if seen[hash] {
				thewinner = p1
				return true
			}
			seen[hash] = true
			//log.Println(len(p1), len(p2), depth)
			//log.Println(hash)
			a := p1[0]
			b := p2[0]
			p1 = p1[1:]
			p2 = p2[1:]
			aWins := false
			if len(p1) >= a && len(p2) >= b {
				p11 := make([]int, a)
				p22 := make([]int, b)
				for i := 0; i < a; i++ {
					p11[i] = p1[i]
				}
				for i := 0; i < b; i++ {
					p22[i] = p2[i]
				}
				aWins = recurse(p11, p22, depth+1)
			} else {
				aWins = a > b
			}
			if aWins {
				p1 = append(p1, a, b)
			} else {
				p2 = append(p2, b, a)
			}
		}
		if len(p1) > 0 {
			thewinner = p1
		} else {
			thewinner = p2
		}
		return len(p1) > 0
	}
	recurse(p1, p2, 0)
	p2Result = score(thewinner)
	log.Printf("%d tries", tries)
	return
}

var _ = register(22, day22, ``)
