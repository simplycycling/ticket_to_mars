package main

import (
	"fmt"
	"math/rand"
)

func main() {
	distance := 57600000
	company := ""
	speed := rand.Intn(16) + 15
	days := (distance / speed) / 86164
	
	fmt.Printf("%-18v %4v %v %7v\n", "Spaceline", "Days", "Round-trip", "Price")
	fmt.Println("===============================================")
	fmt.Println(distance / speed, days)
	for count :=0; count < 10; count++ {
		
		fmt.Printf("%-18v %4v %v $%4v\n", company, days, "Round trip", 22)
	}
}
