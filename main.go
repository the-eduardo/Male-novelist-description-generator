package main

import (
	"fmt"
)

func main() {

	debug := true
	if debug == true {
		fmt.Printf("Debug: %v, %v, %v, %v, %v\n\n", debug1(), debug2(), debug3(), debug4(), debug5())
	}

	fmt.Println("She had", repA(), "like a", repB(), repC(), "and I", repD(), "to", repE(), "her")

}
