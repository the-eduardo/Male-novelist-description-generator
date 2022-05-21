package main

import (
	"fmt"
	"math/rand"
	"time"
)

func debug2() int {

	rand.Seed(time.Now().UnixNano())
	rep1 := rand.Intn(26)

	return rep1
}

func repB() string {

	rep1 := debug2()

	switch rep1 {
	case 0:
		return "creamy"
	case 1:
		return "silken"
	case 2:
		return "wiry"
	case 3:
		return "gleaming"
	case 4:
		return "withholding"
	case 5:
		return "bulbous"
	case 6:
		return "brittle"
	case 7:
		return "soft"
	case 8:
		return "dewy"
	case 9:
		return "fat"
	case 10:
		return "shrill"
	case 11:
		return "haughty"
	case 12:
		return "expensive"
	case 13:
		return "plump"
	case 14:
		return "withered"
	case 15:
		return "tempestuous"
	case 16:
		return "voluptuous"
	case 17:
		return "silken"
	case 18:
		return "frigid"
	case 19:
		return "luscious"
	case 20:
		return "wrinkled"
	case 21:
		return "rippling"
	case 22:
		return "muddle-aged"
	case 23:
		return "juicy"
	case 24:
		return "ripe"
	case 25:
		return "shiny"
	case 26:
		return "creamy"
	}
	return fmt.Sprintf("teste")
}
