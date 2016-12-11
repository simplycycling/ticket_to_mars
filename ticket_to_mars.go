package main

import (
	"fmt"
	"math/rand"
)

func main() {
	distance := 56000000
	speed := rand.Intn(16) + 15
	days := (distance / speed) / 86164
	//days := rand.Intn(16) + 15
	//price := days / 2
	fmt.Printf("%-18v %4v %v %7v\n", "Spaceline", "Days", "Round-trip", "Price")
	fmt.Println("===============================================")
	fmt.Println(distance / speed, days)
	/*spaceline := {
		switch rand.Intn(2) {
		case 0:
			"Virgin Galactic"
		case 1:
			"SpaceX"
		case 2:
			"Space Adventures"
		}
	}*/
	/*for count :=0; count < 10; count++ {
		fmt.Printf("%-18v %4v %v $%4v\n", spaceline, days, "Round trip", 22)
	}*/
}
