package main

import (
	"fmt"
	"math/rand"
	"time"
)

func repA() string {

	rand.Seed(time.Now().UnixNano())
	rep1 := rand.Intn(26) + 1
	fmt.Println(rep1)
	if rep1 == 1 {
		return "boobs"
	} else if rep1 == 2 {
		return "a bust"
	} else if rep1 == 3 {
		return "a butt"
	} else if rep1 == 4 {
		return "lips"
	} else if rep1 == 5 {
		return "an ass"
	} else if rep1 == 6 {
		return "breasts"
	} else if rep1 == 7 {
		return "hair"
	} else if rep1 == 8 {
		return "eyes"
	} else if rep1 == 9 {
		return "a voice"
	} else if rep1 == 10 {
		return "curves"
	} else if rep1 == 11 {
		return "a rump"
	} else if rep1 == 12 {
		return "legs"
	} else if rep1 == 13 {
		return "a rear end"
	} else if rep1 == 14 {
		return "mammaries"
	} else if rep1 == 15 {
		return "jugs"
	} else if rep1 == 16 {
		return "calves"
	} else if rep1 == 17 {
		return "a badonkadonk"
	} else if rep1 == 18 {
		return "gams"
	} else if rep1 == 19 {
		return "knockers"
	} else if rep1 == 20 {
		return "a complexion"
	} else if rep1 == 21 {
		return "a pout"
	} else if rep1 == 22 {
		return "hooters"
	} else if rep1 == 23 {
		return "a booty"
	} else if rep1 == 24 {
		return "cheeks"
	} else if rep1 == 25 {
		return "thighs"
	} else if rep1 == 26 {
		return "tresses"
	} else if rep1 == 0 {
		return "BUG"
	}
	return fmt.Sprintf("teste")
}
