package main

import (
	"log"
	"regexp"
	"strings"
)

func day19(input string) (p1Result, p2Result int) {
	rulesIn := `21: 71 69
100: 92 69 | 13 54
116: 34 13 | 41 92
101: 92 93 | 13 96
122: 13 22 | 92 19
124: 70 13 | 76 92
29: 9 92 | 133 13
2: 13 75 | 92 73
117: 132 92 | 109 13
0: 8 11
43: 92 13 | 13 13
56: 41 13 | 34 92
54: 92 92
16: 92 128 | 13 106
39: 13 82 | 92 72
121: 99 13 | 70 92
110: 59 13 | 123 92
4: 63 13 | 33 92
120: 130 92 | 54 13
7: 41 71
119: 76 13 | 68 92
63: 13 99 | 92 76
44: 92 13 | 92 92
3: 57 92 | 117 13
30: 13 68
127: 91 13 | 1 92
49: 13 98 | 92 5
5: 41 13 | 112 92
74: 92 17 | 13 25
58: 13 34 | 92 68
128: 43 13
77: 13 37 | 92 128
112: 13 92
130: 92 92 | 13 13
46: 92 83 | 13 122
31: 13 3 | 92 26
85: 69 92 | 60 13
111: 13 41 | 92 70
76: 92 13
24: 76 13 | 70 92
93: 92 84 | 13 50
78: 92 21 | 13 135
57: 66 92 | 40 13
84: 13 45 | 92 41
109: 48 92 | 78 13
17: 92 60 | 13 43
103: 115 13 | 56 92
99: 13 92 | 71 13
65: 13 95 | 92 120
79: 121 92 | 128 13
134: 34 13 | 76 92
68: 13 92 | 92 13
14: 60 13 | 76 92
83: 4 92 | 27 13
70: 13 92 | 92 71
62: 92 79 | 13 90
66: 2 92 | 108 13
52: 54 92 | 130 13
47: 13 38 | 92 52
108: 119 13 | 134 92
33: 13 130 | 92 45
9: 13 43 | 92 69
64: 24 13 | 86 92
118: 71 70
6: 112 13 | 130 92
72: 71 43
69: 13 13 | 92 71
97: 89 92 | 62 13
135: 36 92 | 99 13
82: 13 70 | 92 54
61: 77 92 | 28 13
90: 92 125 | 13 100
129: 13 45 | 92 54
86: 130 92 | 41 13
106: 92 99 | 13 54
38: 92 99 | 13 69
126: 92 74 | 13 105
45: 13 92 | 13 13
10: 61 13 | 94 92
28: 30 13 | 35 92
71: 13 | 92
27: 118 92 | 9 13
1: 88 92 | 126 13
11: 42 31
105: 13 21 | 92 102
41: 92 92 | 13 92
34: 13 13
125: 60 13 | 112 92
55: 13 41 | 92 43
20: 13 45 | 92 60
113: 43 92
96: 13 111 | 92 113
104: 92 99 | 13 130
87: 54 92 | 112 13
132: 103 92 | 110 13
12: 124 13 | 20 92
133: 13 45 | 92 76
50: 54 13 | 76 92
26: 46 13 | 97 92
42: 13 127 | 92 67
131: 119 92 | 121 13
73: 92 36 | 13 76
19: 55 13 | 129 92
8: 42
75: 60 13 | 130 92
102: 13 68 | 92 99
22: 92 6 | 13 85
80: 13 54 | 92 70
107: 14 92 | 80 13
48: 92 58 | 13 21
23: 92 116 | 13 113
98: 70 92 | 69 13
37: 54 13 | 41 92
59: 13 76 | 92 45
35: 71 99
40: 92 15 | 13 23
81: 13 51 | 92 114
89: 92 49 | 13 32
94: 92 64 | 13 107
36: 71 71
18: 92 65 | 13 53
53: 92 7 | 13 100
15: 119 92 | 87 13
60: 13 71 | 92 92
115: 60 92 | 41 13
91: 92 18 | 13 101
123: 76 13 | 44 92
51: 47 92 | 29 13
114: 13 16 | 92 12
25: 60 13 | 99 92
88: 13 131 | 92 39
95: 92 45 | 13 68
67: 81 92 | 10 13
32: 50 13 | 104 92`
	//13: "b"
	//92: "a"`

	rLines := Lines(rulesIn)
	rules := map[int]string{}
	for _, line := range rLines {
		parts := strings.Split(line, ":")
		num := parts[0]
		rule := strings.TrimSpace(parts[1])
		rules[Atoi(num)] = rule
	}
	var regex func(int) string
	regex = func(n int) string {
		if n == 13 {
			return "b"
		}
		if n == 92 {
			return "a"
		}
		rule := rules[n]
		re := ""
		for i, choice := range strings.Split(rule, "|") {
			if i > 0 {
				re += "|"
			}
			for _, sub := range strings.Split(choice, " ") {
				if sub != "" {
					re += "("
					re += regex(Atoi(sub))
					re += ")"
				}
			}

		}
		return re
	}
	lines := Lines(input)
	log.Println(len(lines))
	f := func() (i int) {
		r0 := "^" + regex(0) + "$"
		//log.Println(r0)
		rg := regexp.MustCompile(r0)
		for _, line := range lines {
			if rg.MatchString(line) {
				i++
			}
		}
		return
	}
	p1Result = f()
	// 8: 42 | 42 8
	//11: 42 31 | 42 11 31
	rules[8] = "42 | 42 42 | 42 42 42 | 42 42 42 42 | 42 42 42 42 42 | 42 42 42 42 42 42"
	rules[11] = "42 31 | 42 42 31 31 | 42 42 42 31 31 31 | 42 42 42 42 31 31 31 31"
	p2Result = f()
	return
}

var _ = register(19, day19, `bbabbaaaabbaabbbaabbabaabbaabbaababaabab
abaababbbaaaaaabbabbabaabbabbaaaabaababbababbbab
baaabbbaabbabaabbbbbbbaa
bbaaaabababbbaaabbbabaab
ababbababbabbababaaababaabbbbaaaabbaabba
bbaabbabbbbbabbabaaabbbaabbabbab
bbabbaaaabbbabbaabbbabbabbabbbaa
abbbaaababbaaaaababaabaaaabababbabbaabab
abaabbabbbabaaaabaabaaba
aabaaabbaaaaabaabbabbbba
aabbbaaaaabbbbabbaabbbaabaabbbbbbbbababbbbbbaabbbbbaabaaabbbabbabbabaaaababaaaaabbbabbbabbaabbaa
abbaabaaabbaaabbaaaaaaabbbabababbbaabababbbaababaababbabbaabbaaaaaaaabbb
bbbbababbabbaabbbabaaabbaaaabbab
baaabababaabbbaabaabbaab
abbabbbbbbbbbbbaabaaabaaabbbbbababbabaaaabbbaaabbaaabaab
bbabababbaaabbbaaabaabbababbaabaaabbababbbbbbaaa
aaabbbaaaaaaababaabbaabb
ababbababaaabbabbbaaababbaaaaaaa
aaaaababbbabbaabbabaaabaabbababaaabbaaaa
aabaabbababababaabbabbab
babbbbbaaabbaaababbabbab
abaaaabbbabaaaabababaabb
babbbaaaabaaababaabababb
abaaaabbbaaaabaaaabaaaab
aabbaaabbaabbbaaabbbaaba
babbbaabababbbbbaabbbabbaaababbbbaabababaabaabaaabaaaaab
ababbabbbaaabaaaaababbabababbbaababaaaabababababbbbabababaaabaaaaabbbbaa
babaaabbbbaaabbabbbababa
bbbaababbbababaabbabbabb
babbabbbaaaabaaabaabaaaa
abaaaaaaaabbbabbbabbaabaaaaabbbb
bbaabaaaabaabaaaaabaaaabaabbbbbbbaaaabaaabbbaabababbbbbabaabaabbbbbababaabaababa
bbabaabbbababababbabbaabbbbababbaabbbbba
aaaabaabbaababbabbabbbaa
bbbbaabbbabaaaaaabbaaaba
aaaabbaaaabaabbbabbabbaabbbabaaa
abbbabaabbaababaaababaabbabbabab
bbbabbaabbbabbbabbababbabaaaababaababaaabbaabbaa
baabbbababbabbaaabaaababaabbbbba
bbbaabbabaabbbaabbaabbab
aaabbabbbbaabbabaabababa
aaaabaaaaabbbaaaabbaabbbabbabaaabbbabbbb
abbbabbbbbababbaababbaab
aababbbbbbabbbbbbaabbabaaabbbababbbbbaba
aabaabbbbaaababaabababaa
baaaaabaabbabbaaabbbbaaa
aaabbbabbbabbaaaabbbabab
bbababaabbaaabbbaaaaabaababaabaabaaabaab
abaabaababaababbbabbaababbaabbbbaabaabaa
abaaaabbbbabababaabbaaaa
babaabaaaaaabbaabbbbabbbabbabbbaaaabbabaaaabbbbbbabababaababbbbbbaaaaaab
abaabbababababbbbbaaabbaaabaaaaa
abbbaaabbbabaabbbababbaaaaaaaaba
abbabaabbbabababbbbbaabaaabbaabbaabaaabbaaabbbbaaaaabbabaaaaabaa
babbaaabbbbbbbbaababbaab
ababbabaababbaaabbabaaba
abababababababbbabaabbaa
abaabababbbbbbbaabbbbaaa
bbbbababbbaabbbbbababbab
bbaabbabababababbbaaaaba
aabaabbabbabbbaaabaaababbaaaaabaaabbbabbbaaabbabbaabbaba
baaabbaaaaababbabbbbabababaababaabbaaabbbbbbaaaabbababbb
baababbbbbbaaaabbbbbbbaaaaaababaabbaaabaaaabaabababaaaaa
bbabababaaabababaabbaabaabaaaabbababbbbaabaaaaab
abbabbbababaababaaaaabbbaabaaabb
babababbabbbbaabbbaaababaabbbaab
baabbbaaaaaabaaabbbaabbaaaabbabbaaaabbbbbaabbaaabbaaaabbbabaababbabbbbabaababbab
bbaaabaaaaabbbaababbaaabbabaaabbabbabbba
abaaababbabbabaaabababbbababaaaaabbaaaaa
abaababaaabbaaababaabbabaabbbbbb
abaababbbbbababbbbbaabbaaabbbaaaaabbbabb
aaabbbaabbaabababbabbbbb
abaaababbabbbaaababbbbbb
baaaabaaabbbabaabbaaabbb
babbaabbabaaaaaabbbbaabbaaaaaaabbbbbabbabbbbabaa
babaabbbaaaaaabbbaaaabbaaaabaaaababbaabbabbaababbbbaaabb
babaaabbabbbbababaaaaabababbabbabaaaaaababaabbbb
abaaaabababbaabbaabbabaaaaaabaabbbabbbbbabbbaaabbaaaabba
baabbbbaabbbbababbbbabaa
bbaaaaabbbbbbbbababaabbabbbbaaaaabbaaaab
abbbabbbbabababbaaabababaabababbbaaaaabaaaabbaaaabbabaaabbaabbba
baaaaabaaaababbbaaaaababbaaabaababaaabba
aaaabaabbbaabbababbababbabaaabbbbbaabbabaaabaaab
abbabaababaaaababbbbbbaabbaaaaaaaaaaaabb
aaaabbaababbabbabbaaabbb
baabababbabaabbaaabbabbbbbbbababbaababbaabbaabbbabaaabba
baaaabaabaabbabaabbbbbbbaaaababbaaabbaab
aabbbbbbabbbbaaaaabbbbbbbbbbabbb
baaabbaabbabababaaaabaabbbbaabbabaaaaaaaaabbbbababbbbbabaabbaaaaabaaabba
abaaababbabbbaaabaaaabaabbaaaabbaabaabaa
babbbabababaaaaaababbbab
baababbbbbaaabaabaabbbaaabbaabbaabbbabab
abaabbababbbbbbabbababaabbbbbbbaaaabaabbbaaaabbb
aaaaaaababaabbaaaabaaaaabbbabbbb
abbbabaabaabbbaabaabbaab
abbabbaaabbbbbbaaaabbaabbabbbaaababaabba
baaabababababaabaabbbabbabaabbabbabaaabbbabbabbbbbababbb
aaaaaaabaaababbababaaaba
bbaababababaaaabbbbbaaaa
bbbbaabbbbbbaabbbbbabaab
abaaaabaabbabaaaababababbbabbbbbbabbabab
babbaabbabbabbaaabbabbbbaabbbabbaabaaaaaaababbbb
babaaaabaaaabaaababbabaabbbbbbaababbaaaa
abbaabaaaaabbaaabbaaabbb
bbbabbabaaababbaabaabbbaaaaaabaa
bbaaabbababaabbbaaabababbabbaababbbbaabaabaabbaaaabaaabbbabaabbabaaabbababbaababbaabbaaa
bbbaaabababbabababaaaabaaabbabbbbaaabbabaaaabbbaabbababbbaaaabababaabaaabbbababb
babbbaaabbbbabababbbbabb
baaaaabbaabbaabaabbaabba
babbbaabbababaabaaaabaaaaaaabababbbaabbbababbabbbbabaabababbbabbaababaaa
baabbbbaaabaababababbaaa
abbbabbbaaabbaaababbabab
bbabbaabbbaabaabbbbaaabaabaaababbbabbaabbbabaaaaabbabaaabbababbbaababbaa
babaaabbbabaaabbbbabbaabbbbabaab
abbbbbaabaaaaabababaabbabaabbbaaabbbaaabbbbabababbabbbbbbaababbbaababbab
aaababbbbbaabbbaabbbabbaaaaaaaabbbaaaaaabbaaabbbbbbbbabb
ababbabaababbabbbaaabbbababbabbabbbbaaabbaabbaab
ababbabaaabaababbbabbabaaabbaaaabaaaabab
aabbaaaabaabbbbbaaabaaba
babbbbbaaaababbbbbaaaabbbbbbbababaabbbab
abbbbbbabababaabbabbbaabbbaabbaa
bbaaabaabbbaababbaaabbbb
aabbabaaaaabbbaaabababaa
aabbaaababbbbaabbaabaaab
bbbbbbabaaaabbbabbbbababaabbbaab
babbaabbbbbbbbbabbaaaaababbabbabaaaaabaa
bbbaaabaaaaabaabbaaaaaaabaabaaba
baabbabbbbaabbababbaaabbaaaababa
ababaaaaababbaababbbbbaabbbaaaaaaabaaaababbaaaab
abbaaabbbbaabbababbbabbaaabaabaaaaaababb
ababbaaabbaaaabbabbaaaaabbaabaabbbaaaaba
ababbbbbabaaaabaabaaaaab
babababbababbbaabbabbaab
baaaabaababbbbbaabbbabbaaaabababaaabbaaaaaaaabbbbabbababbbbaabbbabbaaaab
bbbbababaaaaaaabbaaabbabaabbabbaaaabbbbaabbbaaba
baaaaabaabaaababaaaaabbaaaabbaaaabbbaabb
bbbbabbabbabbbaabaabbbaaabaabaabbabbbabb
bbaaabaaaabaabbaaaaababb
abbababaaababaabbabbbbaa
aabaabbaaaaaababbabbaaabaaabaaba
abbbbaabbbbabbabaaaaaaaaaabbabbabbaabaaaaabbbbbaaaaaaabb
bbaaabbababbaaabbbaaabaaabbaaaaa
aabaaabbabbbaababababbaa
bbbbaababbabaaaaababbbba
abbbbaabbaaaaabbabbabaabbbbbaabbbbaabbabaabaaaaa
aaaabbbabaababbaabbbbbab
babababbaabbaaaaaabbbbaa
babbaaabaaabbabbabaabaababbaaaaabaabbbbb
bbaabbbaabaababbbbbaaabb
baaababaaabbaabaaabbaabb
baaaaabaabbbabbabbbaabbb
abaababbbbaabbbbabbbabaabbbaabbaababbbba
babababbabaaaaaaaababbbaabbaabbbaabababbaaaabbab
aababaabbabbaaabaaaababababaabbb
aaaaabbaaabbaaabbbaaaababbbbabbbabbabbbbaabababababaaababbaaabab
aaabbbabbbbbaaaaaabababb
bbbbbbaaaaabbabbabbbbbabbaaabaababbbaaabbbbbbabaabbbaaaa
ababbbbbbaaaaabaaaaabbaabaabbbbaababababbbbabaab
bbaaabaaababbabababaabbaaabbaaabaaabaaab
baaaabaaababababaababbaaaaaaabababaabbbabbaabaabbbbabbabbaaabbab
babbabbabaabababaaaaaabb
aaaabababbabbaaabbbbaabababaaaabbbaabababbabbbbb
baabbbabbbaaabaaabbbababaabbababbabaaabbbabaabaabaaababaabbabaaaaabbbaba
baaaaabbbabbbabaaaaaabaa
bbaabababbabbaaabbbaaababaaaabab
babbaaabbaaaaaabaabababa
bbabbabbbbbabababaaaabba
bbabbbaaababababaabaabbbaaababbbbabbbbaa
baaabbaaabbaaabbbaaaabba
abbbbbaababbbbbaabaaabba
aaabbbabbaabbbbabbaaaaaa
aaaababaabbbabaaabbabababaabaaabaaabbaba
bbbaabbaabaabababaaabbaaababbabb
bbbabbabbbaaaaabbaaaaaabbbababba
baabbaababbaaabbabbaaaabaabbaabbbaabbaaaabbbaaaa
abbabbbbaabaaaabaaaabaaabaaaababababbaaababbbaaaabbabbbbaabbbaaaaaaabaaaabbababbbaaabababbabbbba
bbaaabbabbaaabaababaabab
bbababaabababaabaaaaabaa
bbbaabaaabbaabaaaaaaaabb
abbabbaabaabbbaabbbaabbb
aabbbabbbabbbaabaabbaaabaaaababbaabbbbbb
abbbaaaaabaababbabbbbabb
bbbaaabaabbabaabaaabababbbbabbbaaaabbabbaababaabaaabaabb
bbbaabbabaaabbabbbaaaaaa
baabbbaaababaaabbbabbabaabaabababbbbaaaa
bbbbbbababbbabaaabababbbaaaababbabbbaaabaababbbbaabbbababbaaabbb
bbbabbabbaababbbbbabbbaaabbababbbabaaaba
aaababbabbbbababababbbba
bbbbbbabababbbaabbbbaabababaaaababbbaaaabaabbababbbbbbab
abaaababbaaababababbbabb
abbabbbaabbbbabbbaabbbababaaabbaababbbababaabbaabaaaaaab
abbabbbbaaaabaaaaabbbaba
bbbaaababbabaabbabaababbaabbaabb
aaababababbabbaabaabbaaa
babbabaaabbbbababababbaabbabbbbb
bbbabbaabbbabbbabaaabbaaaabbabbabbbabaaa
babbabaabaaaaaabbbbbbbbababaabaa
bbaabbbabaabbbbaababbbbbbbaababaabbaaaabaaabaaab
aababbbbbabbababbbababaabaaabbaa
aabbabbbbaabbbababaaabbb
aaababbbaabaababbbabbabaababaaba
abaaaaaaaaaababababbaaabbaaaaaababaaabababbababbbbabbaababbbbabbbbabbbab
abaaaabbbbbbbaabbbabbbaaabbbaabb
bbabaabbaaaabbbaabbbabbbabaaaaaabbaabbbbaabbaabbbbbaabbbbbabbbab
aabbabaabaababbababbbbbaabbaaabbbbbababbbbbababaabbbbabb
bbabaaaabbaaabbaabbbabab
babaaaabbaababbbbbaabbbabbbbbbbaaaaabababbabbaab
abaabbbaabaabababaababbababbbbaa
ababbbbbbbaababaaabbabbbababbbbaabaabbbb
aababbbaaabbaaabaabaaaaa
babbaaaaabbbbababbbaabaabaabbaaa
aabbbaaaabbaabbbbaababbababababaabababaabbbbaaaa
abaababaababbbaaaababaabbabababbababababaababaaabababbbb
aaaabbbaabaabaaaababbaabbbbbbbbb
baaabbabaabbabbbbaababbabaaaaaababbabbbb
babbaaaababaaaabaabaabbabbbaaabbaabbaabb
abbbaabbbaabbbbbbaaabaabaabbababbbbababa
aabbaaababbaabaabbabbaaababbaaabbabaaaaaaabbbbbbbabbbbbbabbabbba
abbbaababbababababbabbbabbababaaabbaaaaaaaabbbbb
aababbbaaaaabaababbaabba
abbbbababbbbbbbaabaabbabaaaabababaaaabbbaabaaabb
bbbaabbabbaaaaabbbabbbba
abbbaaaaaabbabbbaaabaaba
abababbbaabbaabababaabbabaaaabbb
abbaaaabbaababbbbaaaaababbabbababbbbababbbabaaababaaaabbabbbaabbabbbbababbaaabaaaabbaabbbbaaabaa
abaaaabaabababababbaaaaa
aabbaaaaaabaaabaaabaaaaa
babbbaaaaaaabbbaabaaaabbbababababaaababbbbbbbaba
abaaaabbaaaaabaabbabbbabbaabbbbabbabbabaaabaabbababbabbabbaaabbaabbbbbaabbbbaaab
bbaabababaabababababbabb
abbbbaabbababababbbbbbbb
aabbaabaaababaababababaabbbaabbabbbababaabaaabbbabaabaaaaabbbbaabaaabbaaabababbb
ababbbbbbbabbababbbbaababbbabbaaaabbabbaaababbaabbabaaba
abbaaaabababaabbbabbbbbbbbbaababbbbaabbbabbbaabababaaaabbaababaaabbabaab
bbbaaaaabbaaabababababaa
babaaabbabbbbbaabbbabbbababaabab
bbbbbbabbbbbbbbababbbababaaaaababbbbbaabaaaaabbbbaabaaaaaaaaaabb
baabbbbaaabbabbbbbabbbaaaabaaabb
aabaababbabbbabaaaabbabbbaabaaaa
aabbaababbbaaabaaaabbaab
babaaaaaaaabbbabbabbaaaaaaaaabbabbbaababbaaabbab
bbabbbbababbbbaaabbbbbabbbabaaaabbbaaababababababaabbbbbbbbaabbaabbbaaaa
bbababbabbbaabababbbbaabbbababbb
ababbbbbabbbabbbabaabbbb
bababaabbbababaabbaaabbaaabbabbabbaaaabb
abbaaaabbbabaaaaabbbaabbaabbaabbabaaaabaaababbbabaaaaababaabaabbaabbabbbabbbabba
baababbabbbbbaabbbbaaabb
bbbabbaaabbabaabaababaaa
baaaaabababbabaaaaabbbabbaababbabbbaabaaabbbbaabbbbabaab
aabbbbbabaaabaabaaabaaaa
abbbabbbaaaaaaababbbaaba
abbbabbaabbbbbbabbabbabb
bbbaaababbaabababababaababaaaabbabbabbbabbabbabbababbaaa
aaaaaaabaabbabaaaabaabaa
babbaabababbaaaaabbaabbbbabbabaaabbbbaabbbbbbaba
bbabbaaabaabbbababbaaabbbaaaaabaababbababbabbaaababaabbbabbbbbabaababbbb
babbbaabaabbabaaaabaabbaabaabbaa
bbababbabababaabaabaabbbaabbabaaaababbaaabbabbba
aaaaabbababaaaaabaaaaabaabaabaababaaaabbbaaabaaa
babaabaabbbbabaabbabaaabaabaaaab
bbbbbbabbaaabbbabbbababa
bbabbaaababbabaaaaabbbba
bbaaabaaaaaabaabaaabaaab
bbaabbabbaababaaabaabaabbbbbbaabbabbbabb
ababaaabbaabbbbabbbbababababbbbbbababaabbababbaaabbaaaba
bbaabbbaaabaaabbbbaaaabb
babbabbbabbbbbbababbbabbabaaabbbbbbaababbbbaabbbaabbaabaaaaaaaaababaaaaa
babaaabbaabbaabaababaaababbabbaaabbababaabbabbabbbbabaabbbbbbaaa
bbbbaaabaababbbaabbaababbbbbabaa
baaabbbaaaaabbbabbbbbaba
aababaababbbbabaaabaabbbbbbaababbabaaaabbbbbbaabaaababaa
bbbbabbaabbbaaaabbabababbbabaabbbbbbababbbbaaabbbbbaaabb
bbbaaaaabbabbabbbbabbabaaabbbaab
baabbbbabbbbbbbaabbabababbbababbababbabb
aabbabbbbbbabbabbaaaabbaaabbbbbbbabbababbbbbbabb
abbbabaaabaaaabaaaaabbab
aabaabbaabaababbaaaaaaababaabaababbaaaba
bbabbabababbabbbbabaabbb
bbaaaaabaabbbaaababbabbababbbbbbbbbabaaa
aabbaaababbbbabaaabbbbbb
aabbaabaaaaabbbbbaaabaaaabbbbbab
aaaabbaaaabbaababaaaaabbbaababbbbbbbbbbb
abaaababbababababaababbababaabbb
aaababbabbbaaabaababbaaa
abbbabbaabbbabaababbbabababbabbababbaabbaabaaababbbaaaabbabaabaabbaabaababbabbba
babbbbababbaaabaabbbbabbbbabaaba
aaaaabaaababaaaabababbbaababaaaaababbbbbbabbbabbbabbbbbb
abbaabbbbbbbaabbaaaabbbaabababaa
babbbbbaabbabababbaabbbbaabbaabb
bbbabbbabbbbaabbabbbaaab
babbbaaabbbabaabbbbaaabb
bbbbbbbaaaaaaaababbabbba
bbbbbbabaaabababbabbbbbabbaabbaabbbabaaa
bbaabbababbaaabbaabaababbaaaababaaabaabb
aaabbbaababbaabbabaaababbbbbbaabbababaaabbbaabbb
bbbabbbabbbaaaaabbaabbbbbabbbaabbabaaaba
abababbbaababbbabaababababbbaabb
bbbbbaabbabbaababbbababa
aaaabbbabbbaaaaaaabaaaba
aaaaaaaabbaabbabaaaabbbaababbaab
aaababbabbbbaabaabaabbbb
abaaabbbbabbbaabababbabababaaaababaaaaababbbababbbbbaabababaababaababbba
baaaabbbbbaabaaababbabab
aabaabbbbbaabbbabaaabbbbbababbabbaabbabaabaaabababbabaaabbabbaaa
baaabaabbabbbbaabbabaaabaabbbbbb
ababbbbbbaabbabbabaababaabbbbbab
baababbbbbbaaaaabbabbabb
babababbaaabbbabaaabababbbbababa
baabbaaaabbabaabbabaabaaaaabbabaababaabbabbababaababaaaabbbabbbb
aabaababaaabaabbbbbabbbbbabaabbb
aaaaabbabbababaaababbabaaabaaabbbabbbbbb
baabbababbbabbbaaabaaabaaaaaabbabaabaabaabaaaaaaabaabaaabbbbabbababaaaaa
babbabbabbbaabaaaaaaaaabaabbaaaa
baababaababbaaababaaabababaababbbbbababa
abbaabaaaabbbaaabbaaaaabbbaaaaabaabbaababbbababb
baabbbabaaaabbbaabbbbbaa
aaabbbabbaabbbbaababbaaa
abbabbbaabbbbabaaabaabbbbaabaabbaabababbbbaabbba
baabbaaabbbabbabbaaabbbbaaaaaabaaaaabbbabaaaaabbaabbabbaaabbbbbaaaabaaaa
baabbbaabbbaaaababbabbbbaaababbabbbbabbbaaaaaabbaaababbaaaaababa
baaabbabaabbabaabbbbbabb
ababbbbbabbbabaabbaaaabb
aaabbaababaababbabbbababbaababbababbbaaabaabbaaabbabbabbbabababa
aaaaaaabaabbabaaababbabaababaabb
baababaabbabbaaababbbaabaabbbabbaaaaabababaababaababaaaa
abaabaabaabbbabbaaaaaabb
bbaabbbbbabaaaabbbbbbbabaabbbaba
baabbbaabbaaaaabbbbabbbbbbaaabbbaabbbbaaaababaaabaababaa
bababababbaaabbabbbbbbbaabbbaaaaababbaaa
abbbbbaabbababbaabababab
bbaaababbaababaabbbabbabbbbabbababbaaabbabaaabba
abbbabbabbaaabbababaaaba
abaabbabababbbbbabbabbaabbbabbbb
aaaabaaaabaabbabbbabbabb
bbbbbbababbaabaaaabbaaaa
babbaaababbbaaabaabbbbbbaabbabaaabbbababbbabbababbbaabaaaaabbaabbbbbbbaa
babaaaabbabaabbaaaaabaaaababbabb
abbbbaabaaabbbabbbaabbbbbbaababaaaaababbaabbaabb
bbbabaaabbbbbbabaabaabbabbbbbabababababaabbabbbaabaaabaaaaababbbabbababaabbbbbaa
baaaabaababbbaababbbbbaaaaaababaaababbbb
bababbaabaaaaaaabbbbbbbbbbbbbbbbabbbbaaa
bbbbaaabaabbabbbbabbaabbaaabbbabbbbaaaaaabbaaaab
abbabababaababababaaaaaabbbbbbaaabbbbbbaaaaabbbbabbaaaabaaaabbabababaaba
baabbbaaabaabbaaabbabbababbaaaaaabababba
abaaabaabbabbbaabbbababbaaabaaba
baabbbbabaababababbaaaab
abbabababbaaabbabbbbaabbaabababa
abbabbbbabbbabaaabbbbbbb
aaabbbbaabbaabaaababbbbabbbbaabbaababbbaababbaaaababbbbbbbbbaaab
abaabababbbaababbabaaaabababababbaaababbaabababb
babbabbaabbbbababababbaaaabbbbaabbabbbbaaabbbbbababbbaaabbbbbaaabbabbabb
abbaaabbababbabaabaaaaaabbaabbbaabaabbabaababbaabaaaabbabbbababaaababbbb
aababaabaaabababbbabaaab
aabbabbabbbaababbbaabaaa
abbbabaaaababbbabaabababbabbbaaabababaaa
abbbbaabbbbaabbaaaaaabbb
aabaababbbabababbaabbbbaabbbaaaabaaabababababbabaaabaaaaabbbaababaaabbbb
aaaabbaabbababaababababaaababbbb
bbaaabaaabbaaabbbbbbbaaa
bbaabbbbbbbbaaabbabbbbbaabbbaaba
abbabaabaaabbaaaaaabbbaa
baababaabbbbbaabbaabbaba
aabaabaaabaaaabaabbbabbabbbabbbaaabaaaabaaababbb
babbbaabaaabbaaaaabaabbababababbaaaaabbbabababaabbabaaab
abababaabbaaabbbaaabbaabaabababa
abaababbbbbabbbabbabbbba
bbbbbbabbbbabbabaaabbbaaababaaaa
aabbbabbaaabbbababbabbba
baababaabaaaaababbaaabaaabbabbbbbabaabab
aaababbabaababbbbbbbababbbbbaaabbaaabbaababbabab
aaabbaaaabaababbaaaabaaababbaababaaabaaa
baabbbaababbabbabbabbbab
abbbabbaabaabababaaabbbb
aaabbaaaabaaaabaaabbbaaaaaabbbaabbbaababaaaababbbaabaabbaabbbbbb
babbabaaabbabbbbbbbbbabbaabbbbbabbaaabbabbaaaabaabbbbbabbaaaabbbbbaabbbbaabbabbaabbabaaa
aaabbaaababaabbaaaaaaabb
abaaaabbabbabaabaaaabbab
bbbbabababbbbabaabaaaaaaaaaabaaaababbabbabbaabbabbbabbbbabbababb
abbabaaaaababbbababbaaaaaaabbbaaaaaaabbb
bbbaabaababaaabbbbbbaaabababaabb
abbbabbabbbabbaaaaababbabaababbaaabaaaab
bbbabbaababaabbaaababbaa
bbbbbbbababbbbbababbaabaaabaaabbaababbaaabbbbbabbabbbabbaababbaa
baabbbaaaababaabaabbbaab
bbaaabaaaaaabbaaabaaabaaaabbbbbabbbabbbb
abbaabaabababaabbabbabaaabbababa
babbbaababbaaabbbaaababb
bbbbaababaaaabaaaaabbbbb
aabbbaaabbababaaaaaaaaabbaaaaaaa
abaababaaaababbabbbabbaabaabbabbababaaabaababaabababaaaa
ababbbaaaaabbbaaabbbabab
bbbabbaabaaabbbabaabaaaa
bbbbababbaaaabbbbababbbbbaabbbaabaaaaabb
abaaababbaababbababababbababababbbbbabbb
babaaaaaaabaaaaaabaabaaa
babbaabbbbbbababbabbbbbb
bbbaabaababababaaaaaaabb`)
