package main

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	var err error

	//var outfile = flag.String("o", "", "output file")
	var patt = flag.String("p", "", "pattern to use when generating the wordlist")

	flag.Parse()

	v := *patt
	if v == "" {
		_, _ = fmt.Fprintln(os.Stderr, "Empty pattern")
		os.Exit(1)
	}

	//better code starts here
	ts := make([]token, 0)
	esc := false //escape sequence flag
	open := 0    //0: outside token, 1: inside token, 2: inside custom character area
	s := -1      //start index of number; -1 if no number
	tempTok := token{}
	for i := 0; i < len(v); i++ {
		switch open {
		case 1: //main statement parsing
			switch {
			case v[i] == allOp: //everything that isn't none
				tempTok.p.numSym = true
				tempTok.p.sideSym = true
				tempTok.p.num = true
				tempTok.p.upLetter = true
				tempTok.p.downLetter = true
			case v[i] == noneOp:
				tempTok.p.none = true
			case v[i] == allSymOp:
				tempTok.p.numSym = true
				tempTok.p.sideSym = true
			case v[i] == numSymOp:
				tempTok.p.numSym = true
			case v[i] == sideSymOp:
				tempTok.p.sideSym = true
			case v[i] == allLetterOp:
				tempTok.p.upLetter = true
				tempTok.p.downLetter = true
			case v[i] == upLetterOp:
				tempTok.p.upLetter = true
			case v[i] == downLetterOp:
				tempTok.p.downLetter = true
			case v[i] == numOp:
				tempTok.p.num = true
			case v[i] == '{':
				open = 2
			case v[i] == ']':
				ts = append(ts, tempTok)
				tempTok = token{}
				open = 0
				s = -1
			default:
				_, _ = fmt.Fprintf(os.Stderr, "char %d: invalid character in body of token", i)
				os.Exit(1)
			}

		case 2: //custom character set parsing
			if esc {
				tempTok.p.custom = append(tempTok.p.custom, rune(v[i]))
				esc = false
			} else {
				switch v[i] {
				case '\\':
					esc = true
				case '}':
					open = 1
				default:
					tempTok.p.custom = append(tempTok.p.custom, rune(v[i]))
				}
			}

		case 0: //number/opening parsing
			switch {
			case v[i] >= '0' && v[i] <= '9':
				if s == -1 {
					s = i
				}
			case v[i] == '[':
				if s >= 0 {
					var n uint64
					n, err = strconv.ParseUint(v[s:i], 10, 32)
					if err != nil {
						_, _ = fmt.Fprintf(os.Stderr, "char %d: error parsing number", i)
						os.Exit(1)
					}
					tempTok.reps = int(n)
				} else {
					tempTok.reps = 1
				}
				open = 1
			default:
				_, _ = fmt.Fprintf(os.Stderr, "char %d: invalid character out of token", i)
				os.Exit(1)
			}
		}
	}
	if !reflect.DeepEqual(tempTok, token{}) {
		_, _ = fmt.Fprintf(os.Stderr, "Unterminated bracket somewhere")
		os.Exit(1)
	}
	//for _, t := range ts {
	//	fmt.Println(t)
	//}

	//shitty code below this line
	var cpos []int
	var charsets = make([][]rune, 0)

	for _, t := range ts {
		for i := 0; i < t.reps; i++ {
			cs := make([]rune, 0)
			if t.p.none {
				cs = append(cs, -1)
			}
			if t.p.numSym {
				cs = append(cs, []rune(numSymSet)...)
			}
			if t.p.sideSym {
				cs = append(cs, []rune(sideSymSet)...)
			}
			if t.p.num {
				cs = append(cs, []rune(numSet)...)
			}
			if t.p.upLetter {
				cs = append(cs, []rune(upLetterSet)...)
			}
			if t.p.downLetter {
				cs = append(cs, []rune(downLetterSet)...)
			}
			cs = append(cs, t.p.custom...)

			charsets = append(charsets, cs)
		}
	}
	cpos = make([]int, len(charsets))

	//for _, runes := range charsets {
	//	for _, r := range runes {
	//		fmt.Print(string(r))
	//	}
	//	fmt.Println()
	//}

	b := strings.Builder{}
	for {
		for i := len(charsets) - 1; i >= 0; i-- {
			if cpos[i] == len(charsets[i]) {
				cpos[i] = 0
				if i == 0 {
					goto gtfo
				}
				cpos[i-1]++
			}
		}
		for i, charset := range charsets {
			if charset[cpos[i]] != -1 {
				b.WriteRune(charset[cpos[i]])
			}
		}
		fmt.Println(b.String())
		b.Reset()

		cpos[len(cpos)-1]++
	}
gtfo:
}

type token struct {
	reps int
	p    cPatt
}
