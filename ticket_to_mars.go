package main

import (
	"fmt"
	"math/rand"
)

func main() {
	distance := 57600000
	company := ""
	trip := ""
	
	fmt.Printf("%-18v %4v %v %7v\n", "Spaceline", "Days", "Round-trip", "Price")
	fmt.Println("===============================================")
	//fmt.Println(distance / speed, days)
	for count :=0; count < 10; count++ {
		switch rand.Intn(3) {
		case 0:
			company = "Space Adventures"
		case 1:
			company = "SpaceX"
		case 2:
			company = "Virgin Galactic"
		}
		speed := rand.Intn(16) + 15
		days := (distance / speed) / 86164
		price := days / 2
		
		switch rand.Intn(2) {
		case 0:
			trip = "One Way"
		case 1:
			trip = "Round Trip"
		}
		fmt.Printf("%-18v %4v %v $%4v\n", company, days, trip, price)
	}
}
