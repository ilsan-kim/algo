package main

import "testing"

func TestCase1(t *testing.T) {
	s1 := StatusOne
	s2 := StatusTwo
	s3 := StatusThree
	s4 := StatusFour
	case1(s1)
	case1(s2)
	case1(s3)
	case1(s4)

	var si1 StatusInterface = StatusOneImpl("")
	var si2 StatusInterface = StatusTwoImpl("")
	var si3 StatusInterface = StatusThreeImpl("")
	var si4 StatusInterface = StatusFourImpl("")
	case2(si1)
	case2(si2)
	case2(si3)
	case2(si4)
}
