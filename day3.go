package main

func day3(input string) (p1Result, p2Result int) {
	lines := Lines(input)
	count := func(over, down int) int {
		x, y := 0, 0
		trees := 0
		for y < len(lines) {
			if lines[y][x] == '#' {
				trees++
			}
			x += over
			y += down
			if x >= len(lines[0]) {
				x -= len(lines[0])
			}
		}
		return trees
	}
	p1Result = count(3, 1)
	p2Result = count(1, 1) * p1Result * count(5, 1) * count(7, 1) * count(1, 2)
	return
}

var _ = register(3, day3, `....#...............##...#...#.
#...#..#.....##.##.#.##....#...
...#.....#...#.................
#..#..#.......#...#.#..........
...##..#.#..........##...#.....
........###.#.##..#............
...###......##.#..#.#...#.#....
......##..#.#....#...#.........
.................#......#......
..............##....#..........
#.....................#...#.#.#
.##..#............##...##.##..#
.....#.####...#..##......#.#..#
#.......#.#..#......##.#.#....#
.....##...###.#..........##....
#...........#.##....##.....#..#
..###..##.##.....#....#........
...#.#.#............#.#..#....#
#......#....#...##.#.#.#.#..#..
.......#.#...#..#..#..##......#
.....#..#.............#..#...#.
##..#.##.....#........#........
....##....#..#...........#...#.
.......#........##.......##....
..##...#.......#........##.#...
..........#..#.....##........#.
..#..##..#............#........
.#.#...#...#.......#......#....
....#....#.....#.#.........###.
.............#...#....#..#.#...
##.#...#..#......#.#.##.....#..
#...##.#..........#..#.#...#...
#####.......#.#.....#..#.......
#...#...#........#....#...#....
......##.#..#..#............#..
....#....#.......#...###.......
.#......##...#.##....#...#.....
..#....#...##.....#.#...##.#...
#.......#........#.####........
#.##..#..#.........#.#........#
.#...#.#.#.#......#....#.#..#..
#...####...##.##.#....#......##
..#...#......##........#.....#.
...#.#....##...................
...##................#.........
...##.....##........#....#..#..
.........#..#.....#............
.#..#.......................#..
.#.........#..##........#.#.#.#
......#.....##..#.##...#..#.##.
..#..............##.......#....
...............##....#...##..#.
###...#..###.........#...#.....
...#..#...#....#.....##........
....#..##...#........#.........
..#......#.......#.....#..#....
.#...........##.....###....#...
.#..#.....##.........##.....#..
....##.#.....#................#
..#..#......#.#..#....#..##....
#.....#...##............#......
.#.............#....#.......#..
#.........#..#...##.#...#.#.##.
...#......#..####....#.#.....#.
......#........#..........##.##
......##.#..##.##.....#........
##.....#..##.##..#.......##....
.##.........#...........#...#..
.....#...###..#...#...........#
..........#.#......#.###.....#.
...#.............#.##......##..
#.##.........#..###...........#
....#..##....#..#..#........###
...#........##.......##..#..#..
...#......#..#.#...............
#......###....#.#..#.#..#....#.
#.#.####.#.........#..#.#.#....
.....#....#...............#...#
.#........#......#.#...#.......
................#...#.....##...
.............#...####..........
.................##....##.###..
#................#......#......
.###.#........#..##.....#..###.
..#.#..#...#..#.#...#.#.....#.#
.....#............#..##..#..#.#
#........##.#...#.....#........
#.#.#..###......###............
...#..#...........##...#.....#.
......#........#...#.#....#....
....#..........#.#..#.#....##..
...#.....##..#......#.#.##...#.
.........#..#................#.
..#....#.##.....#.......#......
...#.....#.......##.##.....#...
#...#..............#..###..#...
#.#......#.#....#........##..#.
...#...##...##..#.........#....
..#...#......##.#.#.#....##....
#.......#.......#..#..#........
.........#..#.....#....#.....##
.#......#.......#.#..#..#...#.#
..#....#.#..#..................
#.....####..........#.#.....#.#
.#..#.#.#....#.#.....#.#.......
....##......#..#.....#.#.#...#.
...##...#......##.#....##.#....
..#..##....#...#...........#...
.......#........#...##.#.......
#.#..#....##.#....##...........
.......#............#..##..##..
#.#.#.....#....##.#.#.#.....#..
##...#...#.......#..#...#.....#
##..##.##..........#........##.
..............#.....#..#..##...
.......#...#.........#....#.#..
...#..#..#....#.#....##........
..#.......#....#....##.........
#...#.....#..#.#...##....#.....
.....##..#..##..#..............
.....##............#....#.#....
..#.....#....##.#.....#..#.....
#...#..#..#......#.#.#..##.....
.............................##
#...#.#................#....#.#
.#.#.#....##......###..##......
#.....#..#.##.#.#.##...###.....
.........#............##..#....
.#..#...#....#.....#.#........#
...............#......#..#.....
...................###........#
.###..##..##.......#.#.........
#.........#......#....#.#...#..
.#.#....#.......#.#..##...##...
.#.....#....##.......#.#.....#.
.........#...#....#.#..........
....###..#..##.#...##....#..#..
...#.#..##.#.........###.#..#..
#...#...........#....#.........
....##...........#.#.#......###
#....#...........##..#.........
###....#.....#.......#....###..
.#.......#....#.#.#.#......#.#.
........#...............#.#.#..
....#.........#.....#...##.##.#
...#............#.............#
..........#..#.................
........#.....##............#.#
..#...##........#...#.....#.#..
....#........#.#.#..........#..
#.#...#...........#............
....#.#...##...........#.....#.
...........#.#..#.....#........
.....#..#..#..#.....#.#.....#.#
#.....#.......#.......#...#....
#.........#....#.#........#..#.
...#..#.........#.....#..#.....
...#..#.............#..........
.#.......#..........#.....#...#
.....#.#......#.......#....#...
...#.....#..#..##....##....#...
.#.#.#..#...#.....#....#.......
..##.#..........#.....#.#......
..#..#.............#...##..##..
.#.............#..#....##...#..
..#...#.....#.................#
..##.......#.....#...#....#....
.#..#.##.........#...#.#...#...
...##.......##..#.....##.##...#
........####.#.........#.......
..#.#...##.#..#..#.......##.#..
.#..#............###..#..#.....
#.....#.#...#.#.......#........
..........#......#.#...#...#...
..#......#..#..#.#...#.........
..###........#.#....#.#...##...
.#.....#..#.#......#........#..
.#...#..#...#....#.......#..#..
..#....#..#.....#.#........#...
#..#.#.........#..........#..#.
.#.....##....#.........#.#.#.#.
#.#...#.....#.#.#....#.#..#....
.........#...................#.
..#.....#..##...#..........#.#.
..............#....#.........#.
.#....#.....#..............##..
#...#...#.#........##.........#
....###....#.#....#.#.........#
.....#........#.....##.........
.#...##..##..#.........##......
............#.....#........#...
..#....#.......#......#..#.#.#.
#.......#.#...........#..##.#..
......#.##......#....#.......#.
.....#........#...###.....#....
###..........#........#.#.#....
.....#...#.#...#...#...##.....#
.##...#.#........#.#....#......
......#.........#.....#.#......
.....#.##.....###.#...#...##..#
.#.#.......##....#..#..#.##....
.####...###.#.#.#.#............
......#..##...#..........#.##.#
......#............#...........
.....#.#..#.......##...##......
......#........#..#....#.#.#.#.
#..#..#.....#..#.....#.......#.
.#...#.....#..............#....
.#....#..#.##.#............####
..........#....#.##...#.#......
...#.#.#.#.#.......#.........#.
##........#..##..#.........#...
..#......#...#..#.#.....#......
..#.#......#...#...#.#.........
........................##.....
...#.##.#........#...#.......#.
..#.#......#....##........#.#..
#......#.##........#..#......#.
.....#..#..#.............#.....
......#......#........#....#...
...#....###.....#..#.#....#....
#.......................#....#.
..#...#...................#....
....#..#.....##.#..#...#.....#.
...#.........#...#.......#.....
..#....#.....#...#...#.#####...
.....####......#...........#...
......#.#..........#...#.#.#..#
###..#.#....#..#...............
...#...###..#..#.#.#...........
.....#...#.##.#.#.###..##......
.........#...........#....##.#.
....#..#......#................
...........#..#..#...#.#.......
..#.....#......##.###..........
.........#...................#.
..........#...#.#....##........
..#...##....#....#.......#...##
#......#.....#...#...#...#.....
....##...#.#.......#.#...##....
...#.....#....#.....#....#.....
#....##.....##..##..........##.
.....#.....#.#.#...............
.#.##....#.....#.#..#....#..##.
.....#.#.....##....#...........
.........#..#.......##..##.....
..#....##.....###...#....#.#...
............#......#.#...#..#..
#..##......#.#.##....#.#.......
.#.#.....#...#.#.#....#.....#..
#....#..#.#....#...#...........
......#.#.....#...#.#.#......#.
###..#....#.###.............#..
..............#####........###.
..#..#.#.#.#......#......#.....
###.........#.#..........#..#.#
.#.........#...#......####.....
..#.......####..#....#...#..#..
#.#..#.#...............#.#.#.#.
###....#.....##.#....#......##.
..#..#........#....###.#.#.....
...#.#..........#.....#...#....
....#......##.#............#..#
...##...#.....#..##....#..#.#.#
.......#.....#..#....#....##.#.
.#..#....#..#......##....##...#
..#......#...#.#..###..#.##....
#...#.....#......##...#.......#
.....#.#.....#...##............
.#..##.##..#..##.#........#....
....#.#......##...#.#.#.#..##..
.#..............##........#....
.##....#..#..#....#...#......#.
............###....##.......##.
..............####.....#.......
........##..##.#...#.......#...
....#..#.....##.......#####...#
.##..##..#.....#...#..#..#....#
##..#.#.#...........#..........
#..#......#...#....#...........
...#..##.#..........#..#.......
........#.#.....#......##......
.....#....#............#.......
.#.#..#....##......#.......###.
.#..#.#........#......#...##..#
...#....#......#..#........#.##
.........#..#...#..#.#.##......
....###.#...........#...#......
.##............#.......#..##...
##...#.#...............#.#...##
..#..#.....#.#..#..#...........
..#..#.##..#......#.##..#.#....
..#...#......#.#...#....##.#...
...###....####......#....#...#.
.......##........#.....##....#.
.........##..........#...#.....
.....#............#.##.#....#.#
..........#...#....##..........
....................#......#...
#......#..#...#.............##.
...........#...................
..#...#.........#.##.#..##.#...
#.#....#.#.....#............#..
.#..#.....#.....####......#.#..
#....#.......##..#...........#.
............#...#.....#..#.#...
#...........#...#####....#...#.
..........#...###..##.........#
#.....###............#..#..#.#.
...##.....#....#......#.....#..
#....#.......#..#......###...#.
...##.##......##..##..........#
.......#.#..#.#..#.#.#.#..#..#.
..#..###...#....#.....#......#.
...#.........#..#.##.#.....###.
..#.........#.##.#..#..#..###..
..####..#.........#.........#.#
..#.#...#.......#....##........
.#......#.#....................
..........#.......#.#..#..#....
..#........#....#.#..#.........
..#.....#.............#....#...
##...#.........#.....#...#.....`)
