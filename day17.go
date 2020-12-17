package main

func day17(input string) (p1Result, p2Result int) {
	lines := Lines(input)
	type tuple struct {
		y, x, z, w int
	}
	active := map[tuple]bool{}
	for i, line := range lines {
		for j, ch := range line {
			if ch == '#' {
				active[tuple{i, j, 0, 0}] = true
			}
		}
	}
	for gen := 0; gen < 6; gen++ {
		newActive := map[tuple]bool{}
		possible := map[tuple]int{}
		for t := range active {
			neighbors := 0
			for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
					for dz := -1; dz <= 1; dz++ {
						for dw := -1; dw <= 1; dw++ {
							if dx == 0 && dy == 0 && dz == 0 && dw == 0 {
								continue
							}
							nt := tuple{y: t.y + dy, x: t.x + dx, z: t.z + dz, w: t.w + dw}
							if active[nt] {
								neighbors++
							} else {
								possible[nt]++
							}
						}
					}
				}
			}
			if neighbors == 2 || neighbors == 3 {
				newActive[t] = true
			}
		}
		for t, n := range possible {
			if n == 3 {
				newActive[t] = true
			}
		}
		active = newActive
	}
	p1Result = len(active)
	return
}

var _ = register(17, day17, `.##..#.#
##.#...#
##.#.##.
..#..###
####.#..
...##..#
#.#####.
#.#.##.#`)
