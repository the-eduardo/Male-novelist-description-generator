package main

import (
	"fmt"
	"math/rand"
	"time"
)

func debug5() int {

	rand.Seed(time.Now().UnixNano())
	rep1 := rand.Intn(26)

	return rep1
}

func repE() string {

	rep1 := debug5()

	switch rep1 {
	case 1:
		return "ravish"
	case 2:
		return "climb"
	case 3:
		return "invade"
	case 4:
		return "grope"
	case 6:
		return "marry"
	case 7:
		return "raw dog it with"
	case 5:
		return "caress"
	case 8:
		return "proposition"
	case 9:
		return "correct"
	case 10:
		return "emotionally manipulate"
	case 11:
		return "spar with"
	case 12:
		return "compliment"
	case 13:
		return "hire"
	case 14:
		return "booty call"
	case 15:
		return "mansulain to"
	case 16:
		return "Insult"
	case 17:
		return "dale"
	case 18:
		return "teabag"
	case 19:
		return "ignore"
	case 20:
		return "fondle"
	case 21:
		return "worship"
	case 22:
		return "examine"
	case 23:
		return "manhandle"
	case 24:
		return "touch"
	case 25:
		return "admire"
	case 26:
		return "demand things from"
	case 0:
		return "demand things from"
	}
	return fmt.Sprintf("teste")
}
