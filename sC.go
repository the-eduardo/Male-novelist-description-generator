package main

import (
	"fmt"
	"math/rand"
	"time"
)

func debug3() int {

	rand.Seed(time.Now().UnixNano())
	rep1 := rand.Intn(26)

	return rep1
}

func repC() string {

	rep1 := debug3()

	switch rep1 {
	case 1:
		return "kitten"
	case 2:
		return "mountain"
	case 3:
		return "pillow"
	case 4:
		return "ice cream cone"
	case 5:
		return "tulip"
	case 6:
		return "lake"
	case 7:
		return "fortress"
	case 8:
		return "lemon"
	case 9:
		return "popsicle"
	case 10:
		return "librarian"
	case 11:
		return "python"
	case 12:
		return "pony"
	case 13:
		return "melon"
	case 14:
		return "bedsheet"
	case 15:
		return "muffin"
	case 16:
		return "bunny rabbit"
	case 17:
		return "fish"
	case 18:
		return "princess"
	case 19:
		return "ghost"
	case 20:
		return "waterfall"
	case 21:
		return "mango"
	case 22:
		return "harpy"
	case 23:
		return "ship"
	case 24:
		return "nun"
	case 25:
		return "berry"
	case 26:
		return "car"
	case 0:
		return "car"
	}
	return fmt.Sprintf("teste")
}
