package main

import (
	"fmt"
	"math/rand"
	"time"
)

func debug4() int {

	rand.Seed(time.Now().UnixNano())
	rep1 := rand.Intn(26)

	return rep1
}

func repD() string {

	rep1 := debug4()

	switch rep1 {
	case 1:
		return "longed"
	case 3:
		return "lusted"
	case 4:
		return "wished"
	case 5:
		return "wanted"
	case 6:
		return "resolved"
	case 7:
		return "dared"
	case 8:
		return "detested"
	case 9:
		return "trembled"
	case 10:
		return "thirsted"
	case 11:
		return "expected"
	case 12:
		return "planned"
	case 13:
		return "deigned"
	case 14:
		return "proposed"
	case 15:
		return "shuddered"
	case 16:
		return "refused"
	case 17:
		return "hated"
	case 18:
		return "scorned"
	case 19:
		return "dreaded"
	case 2:
		return "did not care"
	case 20:
		return "ached"
	case 21:
		return "intended"
	case 22:
		return "hungered"
	case 23:
		return "W  craved"
	case 24:
		return "hankered"
	case 25:
		return "needed"
	case 26:
		return "pined"
	case 0:
		return "pined"
	}
	return fmt.Sprintf("teste")
}
