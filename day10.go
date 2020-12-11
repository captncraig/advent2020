package main

import (
	"fmt"
	"log"
	"sort"
)

func day10(input string) (p1Result, p2Result int) {
	nums := Ints(input)
	nums = append(nums, 0)
	sort.Ints(nums)
	nums = append(nums, nums[len(nums)-1]+3)
	log.Println(nums)
	ones := 0
	threes := 0
	chain := 0
	p2Result = 1
	for i := 1; i < len(nums); i++ {
		b := nums[i]
		a := nums[i-1]
		if b == a+1 {
			ones++
			chain++
		} else if b == a+3 {
			threes++
			switch chain {
			case 0:
			case 1:
			case 2:
				p2Result *= 2
			case 3:
				p2Result *= 4
			case 4:
				p2Result *= 7
			default:
				panic("!!!")
			}
			log.Println("!!!", chain, a, p2Result)
			chain = 0
		} else {
			panic(fmt.Sprintf("%d-%d", a, b))
		}
	}
	p1Result = ones * threes
	return
}

var _ = register(10, day10, `118
14
98
154
71
127
38
50
36
132
66
121
65
26
119
46
2
140
95
133
15
40
32
137
45
155
156
97
145
44
153
96
104
58
149
75
72
57
76
56
143
11
138
37
9
82
62
17
88
33
5
10
134
114
23
111
81
21
103
126
18
8
43
108
120
16
146
110
144
124
67
79
59
89
87
131
80
139
31
115
107
53
68
130
101
22
125
83
92
30
39
102
47
109
152
1
29
86`)
