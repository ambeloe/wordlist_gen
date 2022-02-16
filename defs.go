package main

const (
	allOp  = '*'
	noneOp = '!'

	allSymOp  = '='
	numSymOp  = '@'
	sideSymOp = ':'

	allLetterOp  = '%'
	upLetterOp   = 'A'
	downLetterOp = 'a'

	numOp = '#'
)

const (
	numSymSet  = "~`!@#$%^&*()_-+="
	sideSymSet = "{[}]|\\:;\"'<,>.?/"

	numSet = "0123456789"

	upLetterSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	downLetterSet = "abcdefghijklmnopqrstuvwxyz"
)

//struct representing the possible values a character can have
type cPatt struct {
	none bool

	numSym  bool
	sideSym bool

	num bool

	upLetter   bool
	downLetter bool

	custom []rune
}
