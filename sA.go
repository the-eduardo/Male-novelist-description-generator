package main

import (
	//"fmt"
	"math/rand"
	"time"
)

func debug1() int {

	rand.Seed(time.Now().UnixNano())
	rep1 := rand.Intn(26)

	return rep1
}

func repA() string {

	rep1 := debug1()

	switch rep1 {
	case 0:
		return "tresses"
	case 1:
		return "boobs"
	case 2:
		return "a bust"
	case 3:
		return "a butt"
	case 4:
		return "lips"
	case 5:
		return "an ass"
	case 6:
		return "breasts"
	case 7:
		return "hair"
	case 8:
		return "eyes"
	case 9:
		return "a voice"
	case 10:
		return "curves"
	case 11:
		return "a rump"
	case 12:
		return "legs"
	case 13:
		return "a rear end"
	case 14:
		return "mammaries"
	case 15:
		return "jugs"
	case 16:
		return "calves"
	case 17:
		return "a badonkadonk"
	case 18:
		return "gams"
	case 19:
		return "knockers"
	case 20:
		return "a complexion"
	case 21:
		return "a pout"
	case 22:
		return "hooters"
	case 23:
		return "a booty"
	case 24:
		return "cheeks"
	case 25:
		return "thighs"
	case 26:
		return "tresses"
	}
	return "null"
}
