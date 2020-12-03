package main

import (
	"log"
	"os"
	"time"

	"github.com/captncraig/advent2020/pkg/helpers"
)

var days = make([]func() (int, int), 26)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Need to provide day number")
	}
	num := helpers.Atoi(os.Args[1])
	f := days[num]
	if f == nil {
		log.Fatalf("day %d not defined", num)
	}
	start := time.Now()
	p1, p2 := f()
	log.Printf("Part 1: %d", p1)
	log.Printf("Part 2: %d", p2)
	log.Println(time.Now().Sub(start))
}
