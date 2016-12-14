package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	const distance = 57600000
	spaceline := ""
	trip := ""
	
	fmt.Printf("%-18v %4v %v %7v\n", "Spaceline", "Days", "Round-trip", "Price")
	fmt.Println("==========================================")
	for count :=0; count < 10; count++ {
		
		speed := rand.Intn(16) + 15
		days := (distance / speed) / 86164
		price := days / 2
		
		
		switch rand.Intn(3) {
		case 0:
			spaceline = "Space Adventures"
		case 1:
			spaceline = "SpaceX"
		case 2:
			spaceline = "Virgin Galactic"
		}
		
		switch rand.Intn(2) {
		case 0:
			trip = "One Way"
		case 1:
			trip = "Round Trip"
		}
		
		/*
		The assignment actually called for having the days listed for just a
		one-way trip. I did that, and then decided to change it so that the
		listed days would show the total, round trip time, assuming the return
		distance and rate were the same.
		*/
		if strings.Contains(trip, "Round") {
			price = price * 2
			days = days * 2
		}
		
		fmt.Printf("%-18v %4v %-12v $%4v\n", spaceline, days, trip, price)
	}
}
